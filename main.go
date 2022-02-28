package main

import (
	"fmt"
	"ggenerator/commands"
)

func main(){
	var input string

	for {
		fmt.Print("Enter command : ")
		fmt.Scan(&input)
		
		if input == "exit"{
			break
		}
		commands.StrToCommand(input)
	}
}