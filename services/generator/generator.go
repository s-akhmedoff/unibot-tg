package generator

import (
	"bytes"
	"os/exec"
	"unibot-tg/utils"
)

/*
TODO: Fix path to executing
*/

//Generate - ...
func Generate(arg string) string {
	cmd := exec.Command("bash ./services/generator/generator.sh", arg) // *
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	utils.FailOnError(err, "Failed to execute shell script")
	return out.String()
}
