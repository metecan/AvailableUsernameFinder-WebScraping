package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	var available []string
	// Check per 500 time
	for i := 0; i < 500; i++ {
		randomdata := random(4)

		status, name := findPositive(randomdata)
		if status == 0 {
			available = append(available, name)
		}
		fmt.Println("#", i, ", Status: ", status)
		time.Sleep(1500 * time.Millisecond)
	}
	fmt.Println("Available: ", available)
}
func findPositive(userName string) (int, string) {
	counter := 0
	c := colly.NewCollector()

	c.OnHTML("title", func(e *colly.HTMLElement) {
		title := e.Attr("title")
		counter++
		fmt.Printf("title: %q\n", e.Text)
		c.Visit(e.Request.AbsoluteURL(title))

	})
	
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Checking...", r.URL.String())
	})
	c.Visit("https://github.com/" + userName)
	return counter, userName
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func random(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
