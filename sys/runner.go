package sys

import (
	"fmt"
	"os/exec"
)

func Run(commands []string) {
	cmd := exec.Command(commands[0], commands[1:]...)
	fmt.Println(cmd.Run())
}
