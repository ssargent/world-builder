package service

import (
	"context"
	"errors"
	"testing"

	mock_repository "github.com/ssargent/world-builder/wb-api-go/internal/repository/mocks"
	mock_service "github.com/ssargent/world-builder/wb-api-go/internal/service/mocks"
	"github.com/ssargent/world-builder/wb-api-go/internal/tools"
	"github.com/ssargent/world-builder/wb-api-go/pkg/entities"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestTypeService_CreateType(t *testing.T) {
	type fields struct {
		cache       *mock_service.MockCache
		reader      *mock_repository.MockReaderDB
		writer      *mock_repository.MockWriterDB
		manager     *mock_repository.MockManager
		queries     *mock_service.MockEntityDataProvider
		transaction *mock_repository.MockTransaction
	}
	type args struct {
		in *entities.EntityType
	}
	tests := map[string]struct {
		fields fields
		args   args
		mock   func(*fields)
		want   func(*entities.EntityType, error)
	}{
		// TODO: Add test cases.
		"success": {
			args: args{
				in: &entities.EntityType{
					Parent: &entities.TypeReference{
						TypeID:   tools.UUID(1),
						TypeName: "test",
					},
					Wbtn:            "wbtn:mytype",
					TypeName:        "MyType",
					TypeDescription: "MyType is a type",
					Attributes: []*entities.Attribute{
						{
							Wbatn: "wbatn:coordinates",
						},
						{
							Wbatn: "wbatn:faction",
						},
					},
				},
			},
			mock: func(f *fields) {
				f.queries.EXPECT().
					CreateType(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return(tools.EntityType(2, 1), nil)
				f.cache.EXPECT().
					Get("wbatn:coordinates").
					Times(1).
					Return(tools.Attribute(2, "wbatn:coordinates"), true)
				f.queries.EXPECT().
					CreateTypeAttribute(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return(tools.TypeAttribute(1, 2, 1, false), nil)
				f.cache.EXPECT().
					Get("wbatn:faction").
					Times(1).
					Return(tools.Attribute(3, "wbatn:faction"), true)
				f.queries.EXPECT().
					CreateTypeAttribute(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return(tools.TypeAttribute(1, 3, 2, false), nil)
				f.queries.EXPECT().
					GetTypeByID(gomock.Any(), gomock.Any(), tools.UUID(2)).
					Times(1).Return(tools.EntityType(2, 9), nil)
				f.queries.EXPECT().
					GetTypeByID(gomock.Any(), gomock.Any(), tools.UUID(9)).
					Times(1).
					Return(tools.EntityType(9, 11), nil)
				f.queries.EXPECT().
					GetAttributesForType(gomock.Any(), gomock.Any(), tools.UUID(2)).
					Times(1).
					Return(tools.AttributesForType(2), nil)
				f.manager.EXPECT().
					Transaction(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return(f.transaction, nil)
				f.transaction.EXPECT().
					Commit().
					Times(1).
					Return(nil)
				f.transaction.EXPECT().
					Rollback().
					Times(1).
					Return(nil) // no op call to deferred rollback.
			},
			want: func(got *entities.EntityType, err error) {
				require.NoError(t, err)
				require.NotNil(t, got)
			},
		},
		"error-no-attribute": {
			args: args{
				in: &entities.EntityType{
					Parent: &entities.TypeReference{
						TypeID:   tools.UUID(1),
						TypeName: "test",
					},
					Wbtn:            "wbtn:mytype",
					TypeName:        "MyType",
					TypeDescription: "MyType is a type",
					Attributes: []*entities.Attribute{
						{
							Wbatn: "wbatn:coordinates",
						},
						{
							Wbatn: "wbatn:faction",
						},
					},
				},
			},
			mock: func(f *fields) {
				f.queries.EXPECT().
					CreateType(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return(tools.EntityType(2, 1), nil)
				f.cache.EXPECT().
					Get("wbatn:missing-attribute").
					Times(1).
					Return(nil, false)
				f.queries.EXPECT().
					GetAttributeByWBATN(gomock.Any(), gomock.Any(), "wbatn:missing-attribute").
					Times(1).
					Return(nil, errors.New("bad error"))
				f.manager.EXPECT().
					Transaction(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return(f.transaction, nil)
				f.transaction.EXPECT().
					Rollback().
					Times(1).
					Return(nil)
			},
			want: func(got *entities.EntityType, err error) {
				require.Error(t, err)
				require.Nil(t, got)
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()
			f := fields{
				cache:       mock_service.NewMockCache(ctrl),
				reader:      mock_repository.NewMockReaderDB(ctrl),
				writer:      mock_repository.NewMockWriterDB(ctrl),
				manager:     mock_repository.NewMockManager(ctrl),
				queries:     mock_service.NewMockEntityDataProvider(ctrl),
				transaction: mock_repository.NewMockTransaction(ctrl),
			}

			tt.fields = f

			if tt.mock != nil {
				tt.mock(&f)
			}

			tr := &TypeService{
				cache:   tt.fields.cache,
				reader:  tt.fields.reader,
				writer:  tt.fields.writer,
				manager: tt.fields.manager,
				queries: tt.fields.queries,
			}
			got, err := tr.CreateType(ctx, tt.args.in)
			tt.want(got, err)
		})
	}
}
