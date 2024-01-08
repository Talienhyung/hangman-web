package hangmanweb

import (
	"log"
	"sort"
	"strconv"
)

// MakeBoard generates a score board based on the specified tab and user's data
func (user Data) MakeBoard(tab int, allData [][]string) []string {
	// Sort allData based on the specified tab in descending order
	sort.Slice(allData, func(i, j int) bool {
		valI, _ := strconv.Atoi(allData[i][tab])
		valJ, _ := strconv.Atoi(allData[j][tab])
		return valI > valJ
	})

	var board []string

	// Populate the board based on the length of allData
	switch len(allData) {
	case 0:
		board = []string{"nobody", "nobody", "nobody"}
	case 1:
		board = []string{allData[0][0] + " -" + allData[0][tab], "nobody", "nobody"}
	case 2:
		board = []string{allData[0][0] + " -" + allData[0][tab], allData[1][0] + " -" + allData[1][tab], "nobody"}
	default:
		board = []string{allData[0][0] + " -" + allData[0][tab], allData[1][0] + " -" + allData[1][tab], allData[2][0] + " -" + allData[2][tab]}
	}

	var sentence string

	// Check user's position in the sorted data
	for i := 0; i < len(allData) && i < 3; i++ {
		if allData[i][1] == user.Email {
			sentence = "You are on the podium!"
			break
		}
	}

	// If user is not in the top 3, provide their ranking position
	if sentence == "" {
		sentence = "You are at the " + user.searchPosition(allData) + "th position"
		if sentence == "You are at the th position" {
			sentence = "Log in to see your ranking position"
		}
	}

	// Add the final sentence to the board
	board = append(board, sentence)
	return board
}

// searchPosition searches for the user's position in the sorted data
func (user Data) searchPosition(allData [][]string) string {
	for index, elem := range allData {
		if elem[1] == user.Email {
			return strconv.Itoa(index + 1)
		}
	}
	return ""
}

// SetScoreBoard sets the score board based on the provided board data
func (Sboard *ScoreBoard) SetScoreBoard(board []string) {
	if len(board) != 4 {
		log.Fatal("Error when generated board")
	}
	Sboard.First = board[0]
	Sboard.Second = board[1]
	Sboard.Third = board[2]
	Sboard.Sentence = board[3]
}

// ChangeBoardId updates the score board identifier based on the provided action
func (Sboard *ScoreBoard) ChangeBoardId(action string) {
	if action == ">" && Sboard.Id == 0 {
		Sboard.Id = 6
	} else if action == "<" && Sboard.Id == 6 {
		Sboard.Id = 0
	} else if action == "<" {
		Sboard.Id++
	} else if action == ">" {
		Sboard.Id--
	}
}

// MakeRatioBoard generates a score board based on the win/loss ratio of each player
func (user Data) MakeRatioBoard(allData [][]string) []string {
	var newData [][]string

	// Calculate win/loss ratio for each player and create a new data set
	for index, idem := range allData {
		newData = append(newData, []string{allData[index][0], allData[index][1], calculRatio(idem[3], idem[4])})
	}

	// Generate a score board using the new data set and the specified tab (2 for ratios)
	return user.MakeBoard(2, newData)
}

// calculRatio calculates the win/loss ratio based on the provided win and loss strings.
func calculRatio(winStr, loseStr string) string {
	win, err1 := strconv.Atoi(winStr)
	lose, err2 := strconv.Atoi(loseStr)

	if err1 != nil || err2 != nil {
		log.Fatal(err1, err2)
	}

	if lose == 0 {
		return winStr
	}

	ratio := float64(win) / float64(lose)
	return strconv.FormatFloat(ratio, 'f', 2, 64)
}
