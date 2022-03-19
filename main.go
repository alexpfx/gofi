package main

import (
	"fmt"
	"log"

	"github.com/alexpfx/gofi/dmenu"
	"github.com/alexpfx/gofi/internal/rofi"
)
func main(){
	m := dmenu.New(dmenu.Config{
		Sep: "|",
		Prompt: "Selecione o país",
		Active: "3",		
	})

	 

	out, err := rofi.CallRofi("Japão|Brasil|Argentina|Portugal|Espanha", m.Build())
	if err != nil{
		log.Fatal("erro: ",err)
	}	
	fmt.Println(out)
}
