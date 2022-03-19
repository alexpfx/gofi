package rofi

import (
	"fmt"
	"log"
	"os/exec"
)

const rofi = "rofi"

func checkRofi() bool {
	cmd := exec.Command("command", "-v", rofi)
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func CallRofi(input string, args []string) (string, error) {
	fmt.Println(input)
	fmt.Println(args)
	cmd := exec.Command(rofi, args...)
	in, err := cmd.StdinPipe()
	if err != nil{
		log.Fatal(err)
	}
	
	go func ()  {
		fmt.Fprintln(in, input)
		defer in.Close()
	}()


	res, err := cmd.Output()
	return string(res), err
}
