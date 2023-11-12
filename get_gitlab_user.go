package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// User is a struct that represents a GitLab user
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	State    string `json:"state"`
}

const perPage = 50 // Adjust the number of users per page as needed

func main() {
	// Read environment variables for GitLab token and URL
	token := os.Getenv("GITLAB_TOKEN")
	url := os.Getenv("GITLAB_URL")

	if token == "" || url == "" {
		log.Fatal("GITLAB_TOKEN or GITLAB_URL environment variables not set")
	}

	// Get all of the users from GitLab
	users := getallusrs(token, url)

	// Create a CSV file to write the users
	file, err := os.Create("users.csv")
	if err != nil {
		log.Fatal("Error creating CSV file:", err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header row
	header := []string{"ID", "Username", "Name", "Email", "State"}
	if err := writer.Write(header); err != nil {
		log.Fatal("Error writing header:", err)
	}

	// Write the users to the CSV file
	for _, user := range users {
		row := []string{
			strconv.Itoa(user.ID),
			user.Username,
			user.Name,
			user.Email,
			user.State,
		}
		if err := writer.Write(row); err != nil {
			log.Fatal("Error writing row:", err)
		}
	}
}

// getallusrs returns a slice of all users from GitLab
func getallusrs(token, url string) []User {
	// Create a slice to store all of the users
	var users []User

	// Get the total number of users from GitLab
	total := getTotalUsers(token, url)

	// Iterate over the pages of users
	for page := 1; page <= total; page++ {
		// Get the users for the current page
		usersOnPage := getUsers(token, url, page)

		// Append the users to the allUsers slice
		users = append(users, usersOnPage...)
	}

	return users
}

// getTotalUsers returns the total number of users in GitLab
func getTotalUsers(token, url string) int {
	// Make a GET request to the GitLab API
	req, err := http.NewRequest("GET", url+"/api/v4/users", nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}
	req.Header.Set("PRIVATE-TOKEN", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making request:", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	// Parse the JSON array of users
	var users []User
	if err := json.Unmarshal(body, &users); err != nil {
		log.Fatal("Error unmarshalling JSON:", err)
	}

	// Get the total number of users from the X-Total header
	total := resp.Header.Get("X-Total")
	if total == "" {
		log.Fatal("X-Total header not found")
	}

	// Convert the total to an integer
	totalInt, err := strconv.Atoi(total)
	if err != nil {
		log.Fatal("Error converting total to integer:", err)
	}

	return totalInt
}

// getUsers returns a slice of users for the specified page
func getUsers(token, url string, page int) []User {
	// Make a GET request to the GitLab API with the page and per_page parameters
	req, err := http.NewRequest("GET", url+"/api/v4/users", nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}
	req.Header.Set("PRIVATE-TOKEN", token)

	q := req.URL.Query()
	q.Add("page", strconv.Itoa(page))
	q.Add("per_page", strconv.Itoa(perPage))
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making request:", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	// Parse the JSON array of users
	var users []User
	if err := json.Unmarshal(body, &users); err != nil {
		log.Fatal("Error unmarshalling JSON:", err)
	}

	return users
}
