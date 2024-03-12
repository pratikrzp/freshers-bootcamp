package main

import (
	"fmt"
	"sync"
)

func StringRepeatCounter() {
	inputStrings := []string{"quickk", "kbrown", "fox", "lazy", "dog"}

	resultChannel := make(chan map[string]int, len(inputStrings))

	var wg sync.WaitGroup
	defer wg.Wait()

	for _, str := range inputStrings {
		wg.Add(1)
		go func(str string) {
			defer wg.Done()
			processString(str, resultChannel)
		}(str)
	}

	finalOccurrences := make(map[string]int)
	go func() {
		for result := range resultChannel {
			for key, value := range result {
				finalOccurrences[string(key)] += value
			}
		}
	}()
	wg.Wait()
	close(resultChannel)

	fmt.Println(finalOccurrences)
}

func processString(str string, resultChannel chan<- map[string]int) {
	occurrences := make(map[string]int)
	for _, c := range str {
		occurrences[string(c)]++
	}
	resultChannel <- occurrences
}
