package currency

import (
	"testing"
	"unibot-tg/config"
)

func TestGetCurrency(t *testing.T) {

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
			"correct",
			args{""},
			"10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCurrency(tt.args.arg, *config); string([]byte(got)[:2]) != tt.want {
				t.Errorf("GetCurncy() = %v, want %v", got, tt.want)
			}
		})
	}
}
