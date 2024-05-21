package variables

type PersonWebData struct {
	PersonName string
	BirthDate  string
	Gender     string
	City       string
	VkID       string
}

type AllVariables struct {
	CreateD          bool
	Request          bool
	Name             string
	Surname          string
	Year             string
	Country          string
	StringVar        []string
	FormattedStrings [][]string
	SpecialChar      bool
	NumberOfWords    int
	Check            bool
	CheckPassword    string
	Expand           bool
	Filename         string
	ProvidedFileName string
	HashValue        string
	Hash             string
}
