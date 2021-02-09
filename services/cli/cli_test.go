package cli

import "testing"

func TestCli(t *testing.T) {
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
			args{"uname -s"},
			"Linux\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cli(tt.args.arg); got != tt.want {
				t.Errorf("Cli() = %v, want %v", got, tt.want)
			}
		})
	}
}
