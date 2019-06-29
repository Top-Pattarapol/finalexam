package utility

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestParamToInt(t *testing.T) {

	type args struct {
		c   *gin.Context
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				c: &gin.Context{
					Params: gin.Params{
						gin.Param{"id", "1"},
					},
				},
				key: "id",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParamToInt(tt.args.c, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParamToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParamToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
