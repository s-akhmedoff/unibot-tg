package generator

import "github.com/s-akhmedoff/unibot-tg/utils/generate"

// Generate - ...
func Generate(arg string) string {
	return generate.Password(arg)
}
