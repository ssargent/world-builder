package service

import (
	"context"
	"testing"

	"github.com/ssargent/world-builder/wb-api-go/internal/repository"
	mock_repository "github.com/ssargent/world-builder/wb-api-go/internal/repository/mocks"
	mock_service "github.com/ssargent/world-builder/wb-api-go/internal/service/mocks"
	"github.com/ssargent/world-builder/wb-api-go/pkg/entities"
	"go.uber.org/mock/gomock"
)

func TestTypeService_CreateType(t *testing.T) {
	type fields struct {
		cache   Cache
		reader  repository.ReaderDB
		writer  repository.WriterDB
		manager repository.Manager
		queries EntityDataProvider
	}
	type args struct {
		ctx context.Context
		in  *entities.EntityType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mock   func(*fields)
		want   func(*entities.EntityType, error)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				cache:   mock_service.NewMockCache(ctrl),
				reader:  mock_repository.NewMockReaderDB(ctrl),
				writer:  mock_repository.NewMockWriterDB(ctrl),
				manager: mock_repository.NewMockManager(ctrl),
				queries: mock_service.NewMockEntityDataProvider(ctrl),
			}
			if tt.mock != nil {
				tt.mock(&f)
			}

			tr := &TypeService{
				cache:   tt.fields.cache,
				reader:  tt.fields.reader,
				writer:  tt.fields.writer,
				queries: tt.fields.queries,
			}
			got, err := tr.CreateType(tt.args.ctx, tt.args.in)
			tt.want(got, err)
		})
	}
}
