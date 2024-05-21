package todos

import (
	"dimploma/variables"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sync"
)

const (
	findName             string = `<h1 class="name">(.*?)</h1>`
	findID               string = `<a\s+href="\./\?id=([^"]+)"`
	findPersonBirthData  string = `Дата рождения: (.*?)<br>`
	findPersonGenderData string = `Пол: (.*?)<br>`
	findPersonCityData   string = `Город: (.*?)<br>`
	findPersonVkID       string = `target="_blank" >(.*?)</a>`
)

func SendRequest(name, surname, country string) {

	// URL to fetch
	url := fmt.Sprintf("https://sociumin.com/search.php?q=%s+%s&countryID=%s", name, surname, country)

	// Create a GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set request headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:120.0) Gecko/20100101 Firefox/120.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	//resp body to string conversion
	bodyStr := string(body)
	fmt.Println(bodyStr)

	//regular expressions to find ids
	re := regexp.MustCompile(findID)
	matches := re.FindAllStringSubmatch(bodyStr, -1)

	IdSlice := make([]string, 0, 10)

	for _, match := range matches {
		if len(match) > 1 {
			IdSlice = append(IdSlice, match[1])
		}
	}

	fmt.Print("\tInformation about people found:\n\n")
	personsData := make([]variables.PersonWebData, len(IdSlice))
	wg := new(sync.WaitGroup)
	for i := range IdSlice {
		i := i
		if IdSlice[i] != "" {
			wg.Add(1)
			go func(i int, wg *sync.WaitGroup) {
				defer wg.Done()
				personsData[i] = SendRequestWithId(IdSlice[i])
			}(i, wg)
		}
	}
	wg.Wait()
	for i, v := range IdSlice {
		if v != "" {
			fmt.Println("ID:", v)
			fmt.Printf("  Name:        %s\n", personsData[i].PersonName)
			fmt.Printf("  Birth date:  %s\n", personsData[i].BirthDate)
			fmt.Printf("  City:        %s\n", personsData[i].City)
			fmt.Printf("  Gender:      %s\n", personsData[i].Gender)
			fmt.Printf("  VK ID:       %s\n", personsData[i].VkID)
			fmt.Println()
		}
	}
}

func SendRequestWithId(id string) variables.PersonWebData {
	var PersonData variables.PersonWebData
	// URL to fetch
	url := fmt.Sprintf("https://sociumin.com/?id=%s", id)

	// Create a GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return PersonData
	}

	// Set request headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:120.0) Gecko/20100101 Firefox/120.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return PersonData
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return PersonData
	}
	//resp body to string conversion
	bodyStr := string(body)
	// fmt.Println(bodystr)

	re := regexp.MustCompile(findName)
	name := re.FindStringSubmatch(bodyStr)
	if len(name) > 1 {
		PersonData.PersonName = name[1]
	}

	//regular expressions to find birthDay
	re = regexp.MustCompile(findPersonBirthData)
	birthDay := re.FindStringSubmatch(bodyStr)
	if len(birthDay) > 1 {
		PersonData.BirthDate = birthDay[1]
	}

	//regular expressions to find city
	re = regexp.MustCompile(findPersonCityData)
	city := re.FindStringSubmatch(bodyStr)
	if len(city) > 1 {
		PersonData.City = city[1]
	}

	//regular expressions to find gender
	re = regexp.MustCompile(findPersonGenderData)
	gender := re.FindStringSubmatch(bodyStr)
	if len(gender) > 1 {
		PersonData.Gender = gender[1]
	}

	//regular expressions to find vkid
	re = regexp.MustCompile(findPersonVkID)
	vkid := re.FindStringSubmatch(bodyStr)
	if len(vkid) > 1 {
		PersonData.VkID = vkid[1]
	}
	return PersonData
}
