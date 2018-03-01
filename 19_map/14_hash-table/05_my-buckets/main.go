package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

const NumBuckets = 12

// Hash func which provides even distribution
// accross the buckets
func hashBucket(word string, numBuckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % numBuckets
}

// Simple hash func
// func hashBucket(word string, numBuckets int) int {
// 	return int(word[0]) % numBuckets
// }

func main() {
	// Get the page
	resp, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Scan the page

	// scanner init
	sc := bufio.NewScanner(resp.Body)
	defer resp.Body.Close()

	// setting the split func
	sc.Split(bufio.ScanWords)

	// buckets setup
	// {
	//     bucket0 int: {word string: count int},
	//     bucket1 int: {word string: count int},
	//     bucket2 int: {word string: count int},
	//     ...
	//     bucketN int: {word string: count int},
	// }
	buckets := make(map[int]map[string]int)
	for i := 0; i < NumBuckets; i++ {
		buckets[i] = make(map[string]int)
	}

	// fill the buckets
	for sc.Scan() {
		word := sc.Text()
		hash := hashBucket(word, NumBuckets)
		buckets[hash][word]++
	}

	// Print number of items in buckets
	for k, v := range buckets {
		fmt.Println(k, " -- ", len(v))
	}
}
