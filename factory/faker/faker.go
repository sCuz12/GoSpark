package faker

import (
	"math/rand"
	"strings"
	"time"
)


var selectedName string

func Name() string {
	names := []string{
        "John", "Jane", "Alice", "Bob",
        "Emma", "Michael", "Sophia",
    
	}
	//select randomly
	randomName := names[rand.Intn(len(names))]

	selectedName = randomName

    return randomName
}


func LastName() string {
	lastNames := []string{
		"Smith", "Johnson", "Williams", "Brown", "Jones",
		"Garcia", "Miller", "Davis", "Rodriguez", "Martinez",
		"Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson",
		"Thomas", "Taylor", "Moore", "Jackson", "Martin",
		"Lee", "Perez", "Thompson", "White", "Harris",
	}
	//select randomly
	randomName := lastNames[rand.Intn(len(lastNames))]

	selectedName = randomName

    return randomName	
}

// Email generates a random email
func Email() string {
	var mailName string 

    domains := []string{"example.com", "mail.com", "test.com"}

	if selectedName == "" {
		mailName = strings.TrimSpace(Name())
	} else {
		mailName = selectedName
	}

    return mailName + "@" + domains[rand.Intn(len(domains))]
}

// RememberToken generates a random string of length n
func RememberToken(n int) string {
    const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    s := make([]byte, n)
    for i := range s {
        s[i] = letters[rand.Intn(len(letters))]
    }
    return string(s)
}


func Random() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

    b := make([]byte, 10)
    for i := range b {
        b[i] = charset[rand.Intn(len(charset))]
    }
    return string(b)
}


func RandomBool() bool {
    return rand.Intn(2) == 0
}

func RandomDate() time.Time {
    min := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
    max := time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC).Unix()
    delta := max - min

    sec := rand.Int63n(delta) + min
    return time.Unix(sec, 0)
}

func RandomInt(min, max int) int {
    return rand.Intn(max-min+1) + min
}

func RandomFloat(min, max float64) float64 {
    return min + rand.Float64()*(max-min)
}