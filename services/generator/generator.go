package generator

import (
	"os/exec"
	"unibot-tg/utils"
)

/*
TODO: Fix path to executing
 */

//Generate - ...
func Generate(arg string) string {
	cmd := exec.Command("./services/generator/generator.sh", arg) // *
	out, err := cmd.CombinedOutput()
	utils.FailOnError(err, "Failed to execute shell script")
	return string(out)
}
