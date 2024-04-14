package db

import (
	"context"
	"testing"
)

func TestQueries_GetLanguageIdByCode(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx  context.Context
		code LanguageCode
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "Get Chinese Code",
			args:    args{ctx: context.Background(), code: "chn"},
			want:    1,
			wantErr: false,
		},
		{
			name:    "Get Japanese code",
			args:    args{ctx: context.Background(), code: "jp"},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testQueries.GetLanguageIdByCode(tt.args.ctx, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.GetLanguageIdByCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Queries.GetLanguageIdByCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
