package gofi

import (
	"fmt"
	"os/exec"
)

const rofi = "rofi"

func CallRofi(input string, args []string) (string, error) {
	cmd := exec.Command(rofi, args...)
	in, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	go func() {
		fmt.Fprint(in, input)
		defer in.Close()
	}()

	res, err := cmd.Output()
	return string(res), err
}
