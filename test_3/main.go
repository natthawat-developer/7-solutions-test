package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func FetchBaconIpsum(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text", nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func splitWords(r rune) bool {
	
	return r != '-' && (r < 'A' || r > 'Z') && (r < 'a' || r > 'z')
}

func ExtractMeatWords(text string) map[string]int {
	var meatCount sync.Map
	var wg sync.WaitGroup

	words := strings.FieldsFunc(strings.ToLower(text), splitWords)

	workerCount := 2 * runtime.NumCPU()
	wordChan := make(chan string, len(words))

	worker := func() {
		defer wg.Done()
		for word := range wordChan {
			actual, loaded := meatCount.LoadOrStore(word, new(int32))
			if loaded {
				atomic.AddInt32(actual.(*int32), 1)
			} else {
				atomic.StoreInt32(actual.(*int32), 1)
			}
		}
	}

	wg.Add(workerCount)
	for range workerCount {
		go worker()
	}

	for _, word := range words {
		wordChan <- word
	}
	close(wordChan)

	wg.Wait()

	result := make(map[string]int)
	meatCount.Range(func(key, value any) bool {
		result[key.(string)] = int(atomic.LoadInt32(value.(*int32)))
		return true
	})

	return result
}

func main() {
	app := fiber.New()

	app.Get("/beef/summary", func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		text, err := FetchBaconIpsum(ctx)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch data"})
		}

		meatCount := ExtractMeatWords(text)

		return c.JSON(fiber.Map{"beef": meatCount})
	})

	log.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}
