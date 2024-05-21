package getflags

import (
	"dimploma/todos"
	"dimploma/variables"
	"errors"
	"flag"
	"strconv"
	"strings"
)

func CheckIfWebRequest(request bool, name, surname, country string) (bool, error) {

	if !request {
		return request, nil
	}

	countryId, err := strconv.Atoi(country)
	if err != nil || countryId < 0 || countryId > 224 {
		return request, errors.New("incorrect country id")
	}

	if name == "" || surname == "" {
		return request, errors.New("name and surname is not provided")
	}
	todos.SendRequest(name, surname, country)
	return true, nil
}

func ParseAllFlags() (variables.AllVariables, error) {
	var AllFlags variables.AllVariables

	flag.BoolVar(&AllFlags.CreateD, "D", false, "provide D flag if you want to created dict")
	flag.BoolVar(&AllFlags.Request, "q", false, "provide q flag if you want to send requset")
	flag.BoolVar(&AllFlags.Check, "check", false, "provide check flag if you want to check password strength")
	flag.BoolVar(&AllFlags.Expand, "expand", false, "same as check but additionally will ask you to provide -string of -name & -surname & year and create dictionary with provided string permutations for possible password")

	flag.StringVar(&AllFlags.Name, "name", "", "provide name of person you searching info about")
	flag.StringVar(&AllFlags.Surname, "surname", "", "provide surname of person you searching info about")
	flag.StringVar(&AllFlags.Year, "year", "", "provide birth of person you searching info about")
	flag.StringVar(&AllFlags.Country, "country", "0", "provide country of person you searching info about")
	flag.StringVar(&AllFlags.CheckPassword, "password", "", "provide password you would like to assess")
	flag.StringVar(&AllFlags.Filename, "f", "Passwords.txt", "provide filename to save dictionary into, if there is already exist file, it will be deleted and created again")
	flag.StringVar(&AllFlags.ProvidedFileName, "fc", "", "provide filename from which to take possible passwords, only works with -check and -hash")
	flag.StringVar(&AllFlags.Hash, "hash", "md5", "provide hashing method from list: md5, sha1, sha2, ntlm, lanman")
	flag.StringVar(&AllFlags.HashValue, "hv", "", "provide hash value to decrypt")

	flag.IntVar(&AllFlags.NumberOfWords, "N", 0, "provide number of word to appear in permutation")

	var strs string
	flag.StringVar(&strs, "s", "", "provide words separated by comma, without whitespaces")

	flag.BoolVar(&AllFlags.SpecialChar, "C", false, "provide flag to use special charachters when creating dictionary, CAREFULL dictionary with special charachers will grow exponentially, with 2 words will contain 19602 passwords, with 3 words 5821794 passwords")
	flag.Parse()

	AllFlags.StringVar = splitStrings(strs)

	return AllFlags, nil
}

func splitStrings(str string) []string {
	splitted := strings.Split(str, ",")
	return splitted
}
