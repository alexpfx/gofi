package dmenu

import (
	"reflect"
	"testing"
)

var allConfigs = Config{
	Sep:             "|",
	Prompt:          "select",
	Lines:           2,
	InsensitiveCase: true,
	Active:          "2:4",
	Urgent:          "1:2",
	OnlyMatch:       true,
	NoCustom:        true,
}

func Test_buildParams(t *testing.T) {
	type args struct {
		cfg Config
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "all params",
			args: args{
				cfg: allConfigs,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildParams(tt.args.cfg); got != tt.want {
				t.Errorf("buildParams() = %v, want %v", got, tt.want)

			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "all config params",
			args: args{
				config: allConfigs,
			},
			want: "-dmenu -sep '|' -p 'select' -l 2 -i -a '2:4' -u '1:2' -only-match -no-custom ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.config); !reflect.DeepEqual(got.Build(), tt.want) {
				t.Errorf("New() = %v, want %v", got.Build(), tt.want)
			}
		})
	}
}
