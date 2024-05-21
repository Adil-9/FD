package main

import (
	"dimploma/check"
	"dimploma/dictionary"
	"dimploma/getflags"
	"fmt"
	"log"
)

func main() {
	AllFlags, err := getflags.ParseAllFlags()
	if err != nil {
		log.Print(err.Error())
		return
	}
	var temp int
	if AllFlags.Check {
		temp++
	}
	if AllFlags.CreateD {
		temp++
	}
	if AllFlags.Expand {
		temp++
	}
	if AllFlags.Request {
		temp++
	}
	if temp > 1 {
		log.Print("you are only allowed to set one option -D, -check, -q, -expand at a time")
		return
	}
	if AllFlags.Request {
		_, err := getflags.CheckIfWebRequest(AllFlags.Request, AllFlags.Name, AllFlags.Surname, AllFlags.Country)
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
	if AllFlags.CreateD {
		if len(AllFlags.FormattedStrings) > 12 {
			log.Print("Too many words in a string")
			return
		}

		dictionary.CreateDict(AllFlags)
	}
	if AllFlags.Check {
		if AllFlags.CheckPassword == "" || len(AllFlags.CheckPassword) < 6 {
			log.Print("password is not provided or too short, min length 6")
			return
		}
		check.PasswordChecker(AllFlags.CheckPassword)
		if AllFlags.ProvidedFileName != "" {
			check.PasswordCheckerFromFile(AllFlags.CheckPassword, AllFlags.ProvidedFileName)
		}
	}
	if AllFlags.Expand {
		if AllFlags.CheckPassword == "" || len(AllFlags.CheckPassword) < 6 {
			log.Print("password is not provided or too short, min length 6")
			return
		}
		if len(AllFlags.StringVar) <= 2 && (AllFlags.Name == "" || AllFlags.Surname == "" || AllFlags.Year == "") {
			log.Print("Provide -string with 3 words minimum or -year, -name, -surname")
			return
		}
		check.PasswordChecker(AllFlags.CheckPassword)
		if AllFlags.Name != "" {
			temp := []string{AllFlags.Name, AllFlags.Surname, AllFlags.Year}
			AllFlags.StringVar = temp
			dictionary.CreateDict(AllFlags)
			check.CheckInCreatedDict(AllFlags.CheckPassword)
		} else {
			dictionary.CreateDict(AllFlags)
			check.CheckInCreatedDict(AllFlags.CheckPassword)
		}
	}
	if AllFlags.Hash != "" {
		if AllFlags.Hash == "md5" || AllFlags.Hash == "sha1" || AllFlags.Hash == "sha2" || AllFlags.Hash == "ntml" || AllFlags.Hash == "lanman" {
			if AllFlags.HashValue != "" {
				check.HashCheck(AllFlags.Hash, AllFlags.HashValue)
			}
		}
	}
}


