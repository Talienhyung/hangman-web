package hangmanweb

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type Data struct {
	Username string
	Email    string
	Password string
	Win      int
	Loose    int
}

func (data *Data) UploadUserData(allData [][]string) {
	userData := []string{data.Username, data.Email, data.Password, strconv.Itoa(data.Win), strconv.Itoa(data.Loose)}
	allData = RemplaceData(userData, allData)
	file, err := os.Create("./BDD/data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// initialize csv writer
	writer := csv.NewWriter(file)

	defer writer.Flush()

	// write all rows at once
	writer.WriteAll(allData)
}

func Log(email, password string, data [][]string) bool {
	for _, i := range data {
		if i[1] == email && i[2] == password {
			return true
		}
	}
	return false
}

func ReadAllData() [][]string {
	f, err := os.Open("./BDD/data.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func RemplaceData(userData []string, data [][]string) [][]string {
	for _, i := range data {
		if i[1] == userData[1] {
			i = userData
		}
	}
	return data
}

func (data *Data) SetNewUserData(email, password, username string, allData [][]string) [][]string {
	data.Email = email
	data.Password = password
	data.Username = username
	data.Loose = 0
	data.Win = 0
	userData := []string{data.Username, data.Email, data.Password, strconv.Itoa(data.Win), strconv.Itoa(data.Loose)}
	allData = append(allData, userData)
	return allData
}

func EmailAlreadyUsed(email string, data [][]string) bool {
	for _, i := range data {
		if i[1] == email {
			return true
		}
	}
	return false
}

func (userData *Data) SetUserData(email string, data [][]string) {
	var tab []string
	for _, i := range data {
		if i[1] == email {
			tab = i
		}
	}
	var err error
	userData.Username = tab[0]
	userData.Email = tab[1]
	userData.Password = tab[2]
	userData.Loose, err = strconv.Atoi(tab[3])
	if err != nil {
		log.Fatal(err)
	}
	userData.Win, err = strconv.Atoi(tab[4])
	if err != nil {
		log.Fatal(err)
	}
}
