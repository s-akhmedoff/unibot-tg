package generate

import (
	"bytes"
	"os/exec"

	"github.com/s-akhmedoff/unibot-tg/utils"
)

// Password - ...
func Password(arg string) string {
	cmd := exec.Command("./../../utils/generate/generator.sh", arg)

	var out bytes.Buffer

	cmd.Stdout = &out

	err := cmd.Run()
	utils.FailOnError(err, "Failed to execute shell script")

	return out.String()
}
