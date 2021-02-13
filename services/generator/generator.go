package generator

import (
	"unibot-tg/utils/generate"
)

//Generate - ...
func Generate(arg string) string {
	return generate.Password(arg)
}
