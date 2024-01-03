package hangmanweb

import (
	"log"
	"sort"
	"strconv"
)

func (user Data) MakeBoard(tab int, allData [][]string) []string {
	sort.Slice(allData, func(i, j int) bool {
		valI, _ := strconv.Atoi(allData[i][tab])
		valJ, _ := strconv.Atoi(allData[j][tab])
		return valI > valJ
	})

	board := []string{allData[0][0] + " -" + allData[0][tab], allData[1][0] + " -" + allData[1][tab], allData[2][0] + " -" + allData[2][tab]}
	var sentence string
	for i := 0; i < 3; i++ {
		if allData[i][1] == user.Email {
			sentence = "You are on the podium!"
			break
		}
	}
	if sentence == "" {
		sentence = "You are at the " + user.searchPosition(allData) + "th position"
	}
	board = append(board, sentence)
	return board
}

func (user Data) searchPosition(allData [][]string) string {
	for index, elem := range allData {
		if elem[1] == user.Email {
			return strconv.Itoa(index + 1)
		}
	}
	return ""
}

func (Sboard *ScoreBoard) SetScoreBoard(board []string) {
	if len(board) != 4 {
		log.Fatal("Error when generated board")
	}
	Sboard.First = board[0]
	Sboard.Second = board[1]
	Sboard.Third = board[2]
	Sboard.Sentence = board[3]
}

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

func (user Data) MakeRatioBoard(allData [][]string) []string {
	var newData [][]string
	for index, idem := range allData {
		newData = append(newData, []string{allData[index][0], allData[index][1], calculRatio(idem[3], idem[4])})
	}
	return user.MakeBoard(2, newData)
}

func calculRatio(winStr, looseStr string) string {
	win, err1 := strconv.Atoi(winStr)
	loose, err2 := strconv.Atoi(looseStr)

	if err1 != nil || err2 != nil {
		log.Fatal(err1, err2)
	}

	if loose == 0 {
		return winStr
	}

	ratio := float64(win) / float64(loose)
	return strconv.FormatFloat(ratio, 'f', 2, 64)
}
