package gofi

import (
	"fmt"
	"os/exec"
)

const rofi = "rofi"

func CallRofi(input string, args []string) (string, error) {
	fmt.Println(input)
	fmt.Println(args)
	cmd := exec.Command(rofi, args...)
	in, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	go func() {
		fmt.Fprintln(in, input)
		defer in.Close()
	}()

	res, err := cmd.Output()
	return string(res), err
}
