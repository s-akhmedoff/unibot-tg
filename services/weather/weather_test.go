package weather

import (
	"testing"
	"unibot-tg/config"
)

func TestGetWeather(t *testing.T) {
	config := config.Load(true)

	type args struct {
		Country string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test", args: args{Country: "Tashkent"}, want: "Now"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWeather(tt.args.Country, *config)[20:23]; got != tt.want {
				t.Errorf("GetWeather() = %v, want %v", got, tt.want)
			}
		})
	}
}