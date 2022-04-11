package gofi

import (
	"fmt"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	config := Config{
		Error: "mensagem de erro",
	}

	r := New(config)
	fmt.Println(r.Build())
}

func TestCallRofi(t *testing.T) {
	config := Config{
		Error: "Mensagem de erro",
	}

	r := New(config)

	cr, err := CallRofi("", r.Build())
	if err != nil {
		log.Print(err)
	}
	log.Println(cr)
}
