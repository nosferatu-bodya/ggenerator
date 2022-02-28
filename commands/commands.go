package commands

import (
	"fmt"
	"ggenerator/gamegenerator"
)

var AvailableCommands = []string{"help", "addwords", "listwords", "clearwords", "gengames", "exit"}

func StrToCommand(str string) {
	switch str {
		case "help":
			help()
		case "addwords":
			addWords()
		case "listwords":
			listWords()
		case "clearwords":
			clearWords()
		case "gengames":
			genGames()
		default:
			fmt.Println("Command not found. Type $help to see the list of available commands.")
	}
}

func help(){
	for _, c := range AvailableCommands{
		fmt.Println("- " + c)
	}
}

func addWords() {
	var input string

	fmt.Println("Recommended amout of words is at least 15.")

	for {
		fmt.Print("Enter a word (type $f to finish) : ")
		fmt.Scan(&input)

		if input == "$f" {
			break
		}
		gamegenerator.EnteredWords = append(gamegenerator.EnteredWords, input)
	}
}

func listWords() {
	if len(gamegenerator.EnteredWords) == 0 {
		fmt.Println("You have to enter words first.")
		return
	}
	gamegenerator.PrintAllWords(gamegenerator.EnteredWords)
}

func clearWords() {
	gamegenerator.EnteredWords = []string{}
}

func genGames() {
	var noOfGames int
	var noOfFeatures int
	var input string
	var usedWords []string
	var useOwnWords bool

	out:
	for {
		fmt.Print("Do you want to generate games from your words or use the build-in library (m/b) ? : ")
		fmt.Scan(&input)

		switch input {
			case "m":
				useOwnWords = true
				usedWords = gamegenerator.EnteredWords
				break out
			case "b":
				useOwnWords = false
				usedWords = gamegenerator.Words
				break out
			default: 
				fmt.Println("Wrong input.")
				return
		}
	}

	if len(gamegenerator.EnteredWords) == 0 && useOwnWords{
		fmt.Println("You have to enter words first.")
		return
	}

	fmt.Print("How many features the game must have? : ")
	if _, err := fmt.Scan(&noOfFeatures); err != nil {
		fmt.Println("Wrong input.")
		return
	}

	if noOfFeatures > len(gamegenerator.EnteredWords) && useOwnWords{
		fmt.Println("Please, enter a smaller number like 3 or 2.")
		return
	}

	fmt.Print("How many games you want to generate? : ")
	if _, err := fmt.Scan(&noOfGames); err != nil {
		fmt.Println("Wrong input.")
		return
	}

	gamegenerator.GeneratedGames = gamegenerator.GetGeneratedGames(usedWords, noOfFeatures, noOfGames)
	gamegenerator.PrintAllGames(gamegenerator.GeneratedGames)
}