package pwd

import (
	"testing"
)

func TestIsCommon(t *testing.T) {
	type args struct {
		pwd string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{{
		name: "azerty",
		args: args{pwd: "azerty"},
		want: true,
	},
		{
			name: "fdskmlfksdmlfkdmslfmlds",
			args: args{pwd: "fdskmlfksdmlfkdmslfmlds"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCommon(tt.args.pwd); got != tt.want {
				t.Errorf("IsCommon() = %v, want %v", got, tt.want)
			}
		})
	}
}
