package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	const numStudents = 200
	var wg sync.WaitGroup
	wg.Add(numStudents)

	ratings := make(chan int, numStudents)

	// rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < numStudents; i++ {
		go func(studentID int) {
			defer wg.Done()

			delay := time.Duration(rand.Intn(2000)) * time.Millisecond
			time.Sleep(delay)

			rating := rand.Intn(10) + 1
			fmt.Printf("Student %d rated: %d\n", studentID+1, rating)
			ratings <- rating
		}(i)
	}

	wg.Wait()

	close(ratings)

	var total, count int
	for r := range ratings {
		total += r
		count++
	}

	average := float64(total) / float64(count)
	fmt.Printf("\nAverage rating: %.2f\n", average)
}
