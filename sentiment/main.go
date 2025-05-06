package main

import (
	"fmt"
	"strings"
	"sync"
)

func analyzeSentiment(data string, resultChan chan <-string){
	defer wg.Done()
	if(strings.Contains(strings.ToLower(data), "happy")){
		resultChan <- "Positive"
	}else if(strings.Contains(strings.ToLower(data), "sad")){
		resultChan <- "Negative"
	}else{
		resultChan <- "Neutral"
	}
}
var wg sync.WaitGroup
func main(){
	reviews := []string{
		"I love this product! It's amazing! I'm happy",
		"This is the worst thing I've ever tasted. I'm sad",
		"I'm indifferent about this product. I'm neutral",
	}

	resultChan := make(chan string)

	for _, review := range reviews{
		wg.Add(1)
		go analyzeSentiment(review, resultChan)
	}
	
	
	results := make([]string, len(reviews))
	for i := range reviews{
		results[i] = <-resultChan
	}
	close(resultChan)
	for _, result := range results{
		fmt.Println(result)
	}
	wg.Wait()
}