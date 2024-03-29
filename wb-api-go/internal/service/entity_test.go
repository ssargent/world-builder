package service

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/ssargent/world-builder/wb-api-go/internal/repository"
	mock_repository "github.com/ssargent/world-builder/wb-api-go/internal/repository/mocks"
	mock_service "github.com/ssargent/world-builder/wb-api-go/internal/service/mocks"
	"github.com/ssargent/world-builder/wb-api-go/internal/tools"
	"github.com/ssargent/world-builder/wb-api-go/pkg/entities"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestEntityService_get(t *testing.T) {
	type fields struct {
		c       *mock_service.MockCache
		db      *mock_repository.MockWriterDB
		queries *mock_service.MockEntityDataProvider
		manager *mock_repository.MockManager
	}
	type args struct {
		ctx          context.Context //nolint:containedctx // ok here.
		db           repository.DBTX
		id           uuid.UUID
		associations bool
		attributes   bool
		children     bool
	}
	tests := map[string]struct {
		args args
		mock func(f *fields)
		want func(got *entities.Entity, err error)
	}{
		"success": {
			args: args{
				ctx:          context.Background(),
				id:           tools.UUID(1),
				associations: true,
				attributes:   true,
				children:     true,
			},
			mock: func(f *fields) {
				f.queries.EXPECT().GetEntity(gomock.Any(), gomock.Any(), tools.UUID(1)).Times(1).Return(tools.Entity(1, 2, 3), nil)
				f.queries.EXPECT().
					GetEntityReference(gomock.Any(), gomock.Any(), tools.UUID(2)).
					Times(1).
					Return(tools.EntityReference(2, 3), nil)
				f.queries.EXPECT().
					GetTypeByID(gomock.Any(), gomock.Any(), tools.UUID(3)).
					Times(1).
					Return(tools.EntityType(3, 4), nil)
				f.queries.EXPECT().
					GetTypeByID(gomock.Any(), gomock.Any(), tools.UUID(4)).
					Times(1).
					Return(tools.EntityType(3, 4), nil)
				f.queries.EXPECT().
					GetEntityAttributes(gomock.Any(), gomock.Any(), tools.UUID(1)).
					Times(1).
					Return(tools.EntityAttributes(1), nil)
				f.queries.EXPECT().
					GetAttributesForType(gomock.Any(), gomock.Any(), tools.UUID(3)).
					Times(1).
					Return(tools.AttributesForType(3), nil)
				f.queries.EXPECT().
					GetEntityChildReferences(gomock.Any(), gomock.Any(), tools.UUID(1)).
					Times(1).
					Return(nil, nil)
			},
			want: func(got *entities.Entity, err error) {
				require.NoError(t, err)
				require.NotNil(t, got)
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				c:       mock_service.NewMockCache(ctrl),
				queries: mock_service.NewMockEntityDataProvider(ctrl),
				db:      mock_repository.NewMockWriterDB(ctrl),
				manager: mock_repository.NewMockManager(ctrl),
			}
			if tt.mock != nil {
				tt.mock(&f)
			}

			e := &EntityService{
				cache:   f.c,
				reader:  f.db,
				writer:  f.db,
				manager: f.manager,
				queries: f.queries,
			}
			got, err := e.get(tt.args.ctx, tt.args.db, tt.args.id, tt.args.associations, tt.args.attributes, tt.args.children)
			tt.want(got, err)
		})
	}
}

func TestEntityService_CreateEntity(t *testing.T) {
	type fields struct {
		cache       Cache
		reader      repository.ReaderDB
		writer      repository.WriterDB
		manager     repository.Manager
		queries     EntityDataProvider
		typeService TypeService
	}
	type args struct {
		ctx    context.Context
		entity *entities.Entity
	}
	tests := map[string]struct {
		fields fields
		args   args
		mock   func(f *fields)
		want   func(got *entities.Entity, err error)
	}{
		// TODO: Add test cases.
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				cache:       mock_service.NewMockCache(ctrl),
				reader:      mock_repository.NewMockReaderDB(ctrl),
				writer:      mock_repository.NewMockWriterDB(ctrl),
				manager:     mock_repository.NewMockManager(ctrl),
				queries:     mock_service.NewMockEntityDataProvider(ctrl),
				typeService: mock_service.NewMockTypeService(ctrl),
			}
			if tt.mock != nil {
				tt.mock(&f)
			}
			e := &EntityService{
				cache:       tt.fields.cache,
				reader:      tt.fields.reader,
				writer:      tt.fields.writer,
				manager:     tt.fields.manager,
				queries:     tt.fields.queries,
				typeService: tt.fields.typeService,
			}
			got, err := e.CreateEntity(tt.args.ctx, tt.args.entity)
			tt.want(got, err)
		})
	}
}
