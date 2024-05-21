package check

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

const redColor = "\033[0;31m"
const noneColor = "\033[0m"

func PasswordChecker(password string) {
	var found bool
	hash := sha1.New()
	hash.Write([]byte(password))
	hexHash := hex.EncodeToString(hash.Sum(nil))
	hexShort := hexHash[6:]

	dirName := "./leaked/combo_not/"
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
out:
	for _, v := range files {
		file, err := os.Open(dirName + v.Name())
		if err != nil {
			log.Fatal("Failed to open file:", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, hexHash) {
				log.Println(redColor, "The password you typed in WAS FOUND in the leaked LinkedIn database, but it HASN'T yet been cracked in this version of the list. You should probably change your LinkedIn password.", noneColor)
				found = true
				break out
			}
			if strings.Contains(line, hexShort) && v.Name() == files[len(files)-1].Name() {
				log.Println(redColor, "The password you typed in WAS FOUND in the leaked LinkedIn database, and it WAS already cracked. You should probably change your LinkedIn password.", noneColor)
				found = true
				break out
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal("Error reading file:", err)
			return
		}

		if !found && v.Name() == files[len(files)-1].Name() {
			log.Println("The password you typed was NOT FOUND in the leaked LinkedIn database.")
		}
	}
	rockYou(password)
}

func rockYou(password string) {
	var found bool
	dirName := "./leaked/rockyou/"
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
out:
	for _, v := range files {
		file, err := os.Open(dirName + v.Name())
		if err != nil {
			log.Fatal("Failed to open file:", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if password == scanner.Text() {
				log.Println(redColor, "This password was FOUND in rockyou password leak, it is easy to crack, better change it!", noneColor)
				found = true
				break out
			}
		}

		if !found && v.Name() == files[len(files)-1].Name() {
			log.Println("The password was NOT FOUND in rockyou password leak! ")
		}
	}
	realHuman(password)
}

func realHuman(password string) {
	var found bool
	dirName := "./leaked/realhuman/"
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
out:
	for _, v := range files {
		file, err := os.Open(dirName + v.Name())
		if err != nil {
			log.Fatal("Failed to open file:", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if password == scanner.Text() {
				log.Println(redColor, "This password was FOUND in realhuman_phill password leak, it is easy to crack, better change it!", noneColor)
				found = true
				break out
			}
		}

		if !found && v.Name() == files[len(files)-1].Name() {
			log.Println("The password was NOT FOUND in realhuman_phill password leak!")
		}
	}
	checkString(password)
}

func checkString(password string) {
	var hasLower, hasUpper, hasNumber, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	if len(password) < 8 {
		log.Println("The password is too short!")
	}
	if !hasLower {
		log.Println("Your password lacks lowercase letters")
	}
	if !hasUpper {
		log.Println("Your password lacks uppercase letters")
	}
	if !hasNumber {
		log.Println("Your password lacks numbers")
	}
	if !hasSpecial {
		log.Println("Your password lacks special charachters")
	}
	if strings.Contains(password, "@gmail.com") ||
		strings.Contains(password, "@yahoo.com") ||
		strings.Contains(password, "@yahoo.fr") ||
		strings.Contains(password, "@mail.ru") ||
		strings.Contains(password, "@yandex.ry") ||
		strings.Contains(password, "@outlook.com") ||
		strings.Contains(password, ".com") || strings.Contains(password, ".kz") || strings.Contains(password, ".ru") {
		log.Println("Possible use of LOGIN/EMAIL in password detected!")
	}
}

func CheckInCreatedDict(password string) {
	var found bool
	file, err := os.Open("Passwords.txt")
	if err != nil {
		log.Println("Failed to open file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if password == scanner.Text() {
			log.Println(redColor, "This password was FOUND in created dictionary, it is possible to crack, better change it!", noneColor)
			found = true
			break
		}
	}

	if !found {
		log.Println("The password was NOT FOUND in created dictionary! ")
	}
}

func PasswordCheckerFromFile(password, providedFileName string) {
	var found bool
	file, err := os.Open(providedFileName)
	if err != nil {
		log.Println("Failed to open file", providedFileName, "\nError:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if password == scanner.Text() {
			log.Println(redColor, "This password was FOUND in", providedFileName, "file", noneColor)
			found = true
			break
		}
	}

	if !found {
		log.Println("The password was NOT FOUND in", providedFileName)
	}
}

func HashCheck(Hash, HashValue string) {
	switch Hash {
	case "md5":
		withMD5HashRockyou(HashValue)
		withMD5HashRealhuman(HashValue)
	case "sha1":
		withSHA1HashRockyou(HashValue)
		withSHA1HashRealhuman(HashValue)
	case "sha256":
		withSHA256HashRockyou(HashValue)
		withSHA256HashRealhuman(HashValue)
		// case "bcrypt":
		// case "scrypt":
	}
}

func withMD5HashRockyou(hashvalue string) {
	dirName := "./leaked/rockyou/"
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
out:
	for _, v := range files {
		file, err := os.Open(dirName + v.Name())
		if err != nil {
			log.Fatal("Failed to open file:", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			hash := encodeMD5(line)
			if hashvalue == hash {
				log.Println(redColor, "The hash decrypts to", line, "and was found in rockyou password leak", noneColor)
				break out
			}
		}
	}
}

func withMD5HashRealhuman(hashvalue string) {
	dirName := "./leaked/realhuman/"
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
out:
	for _, v := range files {
		file, err := os.Open(dirName + v.Name())
		if err != nil {
			log.Fatal("Failed to open file:", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			hash := encodeMD5(line)
			if hashvalue == hash {
				log.Println(redColor, "The hash decrypts to", line, "and was found in realhuman password leak", noneColor)
				break out
			}
		}
	}
}

func withSHA1HashRealhuman(hashvalue string) {
	dirName := "./leaked/realhuman/"
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
out:
	for _, v := range files {
		file, err := os.Open(dirName + v.Name())
		if err != nil {
			log.Fatal("Failed to open file:", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			hash := encodeSHA1(line)
			if hashvalue == hash {
				log.Println(redColor, "The hash decrypts to", line, "and was found in realhuman password leak", noneColor)
				break out
			}
		}
	}
}

func withSHA1HashRockyou(hashvalue string) {
	dirName := "./leaked/rockyou/"
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
out:
	for _, v := range files {
		file, err := os.Open(dirName + v.Name())
		if err != nil {
			log.Fatal("Failed to open file:", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			hash := encodeSHA1(line)
			if hashvalue == hash {
				log.Println(redColor, "The hash decrypts to", line, "and was found in rockyou password leak", noneColor)
				break out
			}
		}
	}
}

func withSHA256HashRealhuman(hashvalue string) {
	dirName := "./leaked/realhuman/"
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
out:
	for _, v := range files {
		file, err := os.Open(dirName + v.Name())
		if err != nil {
			log.Fatal("Failed to open file:", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			hash := encodeSHA256(line)
			if hashvalue == hash {
				log.Println(redColor, "The hash decrypts to", line, "and was found in realhuman password leak", noneColor)
				break out
			}
		}
	}
}

func withSHA256HashRockyou(hashvalue string) {
	dirName := "./leaked/rockyou/"
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
out:
	for _, v := range files {
		file, err := os.Open(dirName + v.Name())
		if err != nil {
			log.Fatal("Failed to open file:", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			hash := encodeSHA256(line)
			if hashvalue == hash {
				log.Println(redColor, "The hash decrypts to", line, "and was found in rockyou password leak", noneColor)
				break out
			}
		}
	}
}

// Function to encode a string to MD5
func encodeMD5(s string) string {
	// Create a new MD5 hash
	hash := md5.New()

	// Write the string to the hash
	hash.Write([]byte(s))

	// Get the hash as a byte slice
	hashBytes := hash.Sum(nil)

	// Convert the byte slice to a hexadecimal string
	return hex.EncodeToString(hashBytes)
}

// Function to encode a string to SHA-1
func encodeSHA1(s string) string {
	// Create a new SHA-1 hash
	hash := sha1.New()

	// Write the string to the hash
	hash.Write([]byte(s))

	// Get the hash as a byte slice
	hashBytes := hash.Sum(nil)

	// Convert the byte slice to a hexadecimal string
	return hex.EncodeToString(hashBytes)
}

// Function to encode a string to SHA-256
func encodeSHA256(s string) string {
	// Create a new SHA-256 hash
	hash := sha256.New()

	// Write the string to the hash
	hash.Write([]byte(s))

	// Get the hash as a byte slice
	hashBytes := hash.Sum(nil)

	// Convert the byte slice to a hexadecimal string
	return hex.EncodeToString(hashBytes)
}
