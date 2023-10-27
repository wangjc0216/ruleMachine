package rulemachine

import (
	"context"
	"testing"
)

func TestExecute(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{"somecontext",
			args{
				context.TODO(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Execute(tt.args.ctx)
		})
	}
}
