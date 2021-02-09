package cli

import (
	"log"
	"os/exec"
	"strings"
)

type preparedCommand struct {
	command string
	args    []string
}

func (pc *preparedCommand) format(raw string) {
	split := strings.Split(raw, " ")
	pc.command = split[0]
	pc.args = split[1:]
}

//Cli - execute command
func Cli(arg string) string {
	if arg == "" {
		commandNotInputted := "Command is empty"
		log.Println(commandNotInputted)
		return commandNotInputted
	}

	var command preparedCommand
	command.format(arg)

	cmd := exec.Command(command.command, command.args...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Println(err)
		return "Error detected!"
	}

	if len(out) < 1 {
		out = []byte("Done")
	}

	if len(out) > 4070 {
		return string(out)[:4070] + "\n"
	}

	return string(out)
}
