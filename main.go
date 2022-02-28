package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	features []string
}

var allWords []string
var generatedGames []Game
// var allWords = []string{"forest", "dick", "vagina", "gun", "closet", "dog", "Bohdan", "visim", "devyat", "tea"}

func main(){
	var input string
	
	for {
		fmt.Print("Enter a word ($finish to finish) : ")
		fmt.Scan(&input)
		if input == "$finish"{
			break
		}
		allWords = append(allWords, input)
		printAllWords(allWords)
		fmt.Printf("You enetered %v words.\n", len(allWords))
	}

	for {
		fmt.Println("Generated games: ")
		generatedGames = getGeneratedGames(allWords, 3)
		printAllGames(generatedGames)
		
		fmt.Print("Shuffle again (y/n) ?:")
		fmt.Scan(&input)
		if input == "y"{
			continue
		}else if input == "n"{
			break
		}else {
			fmt.Println("Wrong input!!!")
		}
	}
}

func printAllWords(words []string) {
	var str string
	for i, word := range words{
		if i == (len(words)-1){
			str += word
			continue
		}
		str += word + ", "
	}
	fmt.Println(str)
}

func printGame(game Game){
	var str string = "[ "
	for i, w := range game.features {
		if i == (len(game.features) - 1){
			str += w
			continue
		}
		str += w + " - "
	}
	str += " ]"
	fmt.Println(str)
}

func printAllGames(games []Game){
	for _, g := range games {
		printGame(g)
	}
}

func getGeneratedGames(words []string, noOfFeatures int) []Game {
	var returnList []Game 
	words = shuffleArray(words, 121)
	group := 0
	index := 0
	curFeature := 0

	for index < (len(words)/3 + 1) {
		var curGame Game

		for curFeature < noOfFeatures {
			if group < len(words){
				curGame.features = append(curGame.features, words[group])
			}

			curFeature += 1
			group += 1
		}

		returnList = append(returnList, curGame)

		curFeature = 0
		index += 1
	}

	return returnList
}

func shuffleArray(array []string, iterations int) []string {
	returnArray := array
	curIteration := 0
	for curIteration <= iterations {
		rand.Seed(time.Now().UnixNano())

		fromIndex := rand.Intn(len(array) - 1)
		toIndex := rand.Intn(len(array) - 1)

		var aux string = returnArray[fromIndex]
		returnArray[fromIndex] = returnArray[toIndex]
		returnArray[toIndex] = aux

		curIteration += 1
	}
	return returnArray
}