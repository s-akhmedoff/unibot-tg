package definition

import (
	"testing"
	"unibot-tg/config"
)

func TestGetDefinition(t *testing.T) {
	config := config.Load(true)

	type args struct {
		arg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"incorrect",
			args{"qaq"},
			"No definition was found for word: qaq",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDefinition(tt.args.arg, *config); got != tt.want {
				t.Errorf("GetDefinition() = %v, want %v", got, tt.want)
			}
		})
	}
}
