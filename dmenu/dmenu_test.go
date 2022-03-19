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

func Test_dmenu_Build(t *testing.T) {
	type fields struct {
		config Config
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "set all",
			fields: fields{
				config: allConfigs,
			},
			want: []string{
				"-dmenu",
				"-sep",
				"|",
				"-p",
				"select",
				"-l",
				"2",
				"-i",
				"-a",
				"2:4",
				"-u",
				"1:2",
				"-only-match",
				"-no-custom",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dmenu{
				config: tt.fields.config,
			}
			if got := d.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dmenu.Build() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}
