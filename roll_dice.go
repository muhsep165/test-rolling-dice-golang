package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type PlayerResult struct {
	Player		string `json:"player"`
	Dice 		[]int `json:dice`
}

type PlayerScore struct {
	Player 		string `json:"player"`
	Score 		int `json:"score"`
}

type PlayerTurnDice struct {
	Player		string `json:"player"`
	Turn 		int `json:turn`
}

func main() {
	var numOfPlayers, numOfDice int
	
	fmt.Printf("Number of Players : ")
	fmt.Scanf("%d", &numOfPlayers)
	fmt.Printf("Number of Dice : ")
	fmt.Scanf("%d", &numOfDice)

	var playerTurn []PlayerTurnDice
	var playerScore []PlayerScore
	for i := 1; i <= numOfPlayers; i++ {
		playerTurn = append(playerTurn, PlayerTurnDice{
			Player: 		"Player " + strconv.Itoa(i),
			Turn: 			numOfDice,
		})

		playerScore = append(playerScore, PlayerScore{
			Player: 		"Player " + strconv.Itoa(i),
			Score: 			0,
		})
	}

	getRound(playerTurn, 1, playerScore)
}

func getRound(playerTurn []PlayerTurnDice, countRound int, playerScore []PlayerScore) {
	var playerDice []PlayerResult
	min := 1
	max := 7 // max dice is 6, plus 1 to make 6 show in rand
	
	time := time.Now().UTC().UnixNano()
	rand.Seed(time)
	
	for _, player := range playerTurn {
		var diceGrouping []int
		for j := 1; j <= player.Turn; j++ {
			dice := min + rand.Intn(max - min)
			diceGrouping = append(diceGrouping, dice)
		}
		playerDice = append(playerDice, PlayerResult{
			Player: 	player.Player,
			Dice:		diceGrouping,
		})
	}
	reviewResult(playerDice, countRound, playerScore)
}

func reviewResult(playerDice []PlayerResult, countRound int, playerScore []PlayerScore) {
	title := "====Result Before Reviewed====\n ========= ROUND "+strconv.Itoa(countRound)+" ========="
	printResult(playerDice, title)
	
	var newPlayerDice []PlayerResult
	var playerTurn []PlayerTurnDice
	var transferDice []int
	for j, playerReviewed := range playerDice {
		var newDice []int
		if len(transferDice) > 0 {
			newDice = append(newDice, transferDice...)
			transferDice = nil
		}
		for _, dice := range playerReviewed.Dice {
			if dice == 1 {
				transferDice = append(transferDice, dice)
			} else if dice == 6 {
				playerScore[j].Score = playerScore[j].Score + 1
			} else {
				newDice = append(newDice, dice)
			}

			if j == (len(playerDice) - 1) && dice == 1 {
				if len(newPlayerDice[0].Player) > 0 {
					newPlayerDice[0].Dice = append(newPlayerDice[0].Dice, dice)
					if len(playerTurn) == 0 {
						playerTurn = append(playerTurn, PlayerTurnDice{
							Player:  	newPlayerDice[0].Player,
							Turn: 		1,
						})
					} else {
						playerTurn[0].Turn = playerTurn[0].Turn + 1
					}
				}
			}
		}

		newPlayerDice = append(newPlayerDice, PlayerResult{
			Player:		 playerReviewed.Player,
			Dice: 		 newDice,
		})

		if len(newDice) > 0 {
			playerTurn = append(playerTurn, PlayerTurnDice{
				Player:  	playerReviewed.Player,
				Turn: 		len(newDice),
			})	
		}
	}

	title = "====Result After Reviewed====\n ========= ROUND "+strconv.Itoa(countRound)+" ========="
	printResult(newPlayerDice, title)
	printScore(playerScore)
	if len(playerTurn) > 1 {
		getRound(playerTurn, countRound+1, playerScore)
	} else {
		max, winners := findTheWinners(playerScore)
		fmt.Printf("********THE WINNER IS*********\n")
		fmt.Printf(strings.Join(winners, ", ") + "\n")
		fmt.Printf("Winning Score : " + strconv.Itoa(max) + "\n")
		fmt.Printf("********SWEET VICTORY*********\n")
	}
}

func printResult(playerDice []PlayerResult, title string){
	fmt.Printf(title + "\n")
	for _, content := range playerDice {
		fmt.Printf(content.Player + " : ")
		fmt.Printf(arrayToString(content.Dice, ", ") + "\n")
	}
}

func printScore(playerScore []PlayerScore) {
	fmt.Printf("=======Player's Score=======\n")
	for _, score := range playerScore {
		fmt.Printf(score.Player + " score : " + strconv.Itoa(score.Score) + "\n")
	}

}

func arrayToString(a []int, delimiter string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delimiter, -1), "[]")
}


func findTheWinners(playerScore []PlayerScore) (max int, players []string){
	max = playerScore[0].Score
	for _, value := range playerScore {
		
		if value.Score > max {
			max = value.Score
			players = nil
			players = append(players, value.Player)
		} else if value.Score == max {
			players = append(players, value.Player)
		}
	}
	return max, players
}