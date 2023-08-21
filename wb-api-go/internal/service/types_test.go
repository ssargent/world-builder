package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/ssargent/world-builder/wb-api-go/pkg/entities"
)

func TestTypeService_CreateType(t *testing.T) {
	type fields struct {
		cache   Cache
		reader  *sqlx.DB
		writer  *sqlx.DB
		queries EntityDataProvider
	}
	type args struct {
		ctx context.Context
		in  *entities.EntityType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.EntityType
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TypeService{
				cache:   tt.fields.cache,
				reader:  tt.fields.reader,
				writer:  tt.fields.writer,
				queries: tt.fields.queries,
			}
			got, err := tr.CreateType(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("TypeService.CreateType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TypeService.CreateType() = %v, want %v", got, tt.want)
			}
		})
	}
}
