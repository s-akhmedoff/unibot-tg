package generator

import "testing"

func TestGenerator(t *testing.T) {
	type args struct {
		arg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"wrong",
			args{},
			"Error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Generate(tt.args.arg); got != tt.want {
				t.Errorf("Generator() = %v, want %v", got, tt.want)
			}
		})
	}
}