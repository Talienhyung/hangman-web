package hangmanweb

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// UploadUserData uploads user data to a CSV file
// It takes a 2D slice 'allData' and replaces or appends the user's data to it
func (data *Data) UploadUserData(allData [][]string) {
	userData := []string{data.Username,
		data.Email,
		data.Password,
		strconv.Itoa(data.Win),
		strconv.Itoa(data.Loose),
		strconv.Itoa(data.Score),
		strconv.Itoa(data.BestScore),
		strconv.Itoa(data.WinHard),
		strconv.Itoa(data.WinMedium),
		strconv.Itoa(data.WinEasy)}

	// Replace or append user data in the 2D slice
	allData = RemplaceData(userData, allData)

	// Create or open the CSV file
	file, err := os.Create("./BDD/data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Initialize csv writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write all rows to the CSV file
	writer.WriteAll(allData)
}

// Log checks if the given email and password match any user in the data
// It returns true if there is a match, otherwise false
func Log(email, password string, data [][]string) bool {
	for _, i := range data {
		if i[1] == email && checkPasswordHash(password, i[2]) {
			return true
		}
	}
	return false
}

// ReadAllData reads all data from the CSV file and returns it as a 2D slice
func ReadAllData() [][]string {
	// Open the CSV file
	f, err := os.Open("./BDD/data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Read CSV values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// RemplaceData replaces the user data in the 2D slice with the provided user data
func RemplaceData(userData []string, data [][]string) [][]string {
	for user, i := range data {
		if i[1] == userData[1] {
			data[user] = userData
		}
	}
	return data
}

// SetNewUserData sets new user data for the provided email, password, and username
// It updates the Data struct and appends the new user data to the existing 2D slice
func (data *Data) SetNewUserData(email, password, username string, allData [][]string) [][]string {
	// Update Data struct with new user information
	data.Email = email
	data.Password = hashPassword(password)
	data.Username = username

	// Convert user data to a string slice
	userData := []string{data.Username,
		data.Email, data.Password,
		strconv.Itoa(data.Win),
		strconv.Itoa(data.Loose),
		strconv.Itoa(data.Score),
		strconv.Itoa(data.BestScore),
		strconv.Itoa(data.WinHard),
		strconv.Itoa(data.WinMedium),
		strconv.Itoa(data.WinEasy)}

	// Append the new user data to the existing 2D slice
	allData = append(allData, userData)
	return allData
}

// EmailAlreadyUsed checks if the given email is already used in the data
// It returns true if the email is already used, otherwise false
func EmailAlreadyUsed(email string, data [][]string) bool {
	for _, i := range data {
		if i[1] == email {
			return true
		}
	}
	return false
}

// SetUserData sets the user data in the Data struct based on the provided email
// It retrieves the user data from the 2D slice and updates the Data struct accordingly
func (userData *Data) SetUserData(email string, data [][]string) {
	var tab []string
	for _, i := range data {
		if i[1] == email {
			tab = i
		}
	}

	// Update Data struct with user information
	userData.Username = tab[0]
	userData.Email = tab[1]
	userData.Password = tab[2]
	userData.Loose, _ = strconv.Atoi(tab[3])
	userData.Win, _ = strconv.Atoi(tab[4])
	userData.Score, _ = strconv.Atoi(tab[5])
	userData.BestScore, _ = strconv.Atoi(tab[6])
	userData.WinHard, _ = strconv.Atoi(tab[7])
	userData.WinMedium, _ = strconv.Atoi(tab[8])
	userData.WinEasy, _ = strconv.Atoi(tab[9])
}

// hashPassword generates a bcrypt hash for the provided password
func hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

// checkPasswordHash compares a password with its corresponding bcrypt hash
// It returns true if the password matches the hash, and false otherwise
func checkPasswordHash(password, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}
