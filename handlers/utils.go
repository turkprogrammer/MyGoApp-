package handlers

import (
	"fmt"
	"math/rand"
	"time"
)

// generateUniqueUsername генерирует уникальное имя пользователя
func generateUniqueUsername() string {
	rand.Seed(time.Now().UnixNano())
	adjectives := []string{"happy", "sad", "bright", "dark", "fast", "slow", "red", "blue", "green", "yellow"}
	nouns := []string{"cat", "dog", "bird", "tree", "car", "moon", "sun", "star", "fish", "book"}

	adjective := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	number := rand.Intn(1000)

	return fmt.Sprintf("%s_%s%d", adjective, noun, number)
}

// generateUniqueEmail генерирует уникальный адрес электронной почты
func generateUniqueEmail() string {
	rand.Seed(time.Now().UnixNano())
	domains := []string{"gmail.com", "yahoo.com", "outlook.com", "example.com", "domain.com"}

	username := generateUniqueUsername()
	domain := domains[rand.Intn(len(domains))]

	return fmt.Sprintf("%s@%s", username, domain)
}
