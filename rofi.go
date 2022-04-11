package gofi

import (
	"fmt"
	"os/exec"
)

const rofiCmd = "rofi"

func New(config Config) Show {
	return &show{
		config: config,
	}
}

type Show interface {
	Build() []string
}

func CallRofi(input string, args []string) (string, error) {
	cmd := exec.Command(rofiCmd, args...)
	in, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	go func() {
		fmt.Fprint(in, input)
		defer in.Close()
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), fmt.Errorf(string(out), err.Error())
	}
	return string(out), nil
}

func (d *show) Build() []string {
	return append ([]string{"-show"}, buildParams(d.config)...)
}

type show struct {
	config Config
}

func buildParams(cfg Config) []string {
	argArray := make([]string, 0)
	argArray = Add(argArray, "-e", cfg.Error)
	return argArray
}

type Config struct {
	Error string
}
