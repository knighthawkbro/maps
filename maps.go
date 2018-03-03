package main

import (
	"bufio"
	"fmt"
	"log"
	"maps/array"
	"maps/list"
	"os"
	"regexp"
	"strings"
)

// Maps (private) -
type maps interface {
	Add(key, value interface{}) interface{}
	Remove(key interface{}) interface{}
	Get(key interface{}) interface{}
	Contains(key interface{}) bool
	GetKeys() string

	Size() int
}

func main() {
	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning driver function as an array...")
	fmt.Println("\n*************************************************")
	fmt.Println("")
	arr := array.New()
	driver(arr)

	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning driver function as a list...")
	fmt.Println("\n*************************************************")
	fmt.Println("")
	lst := list.New()
	driver(lst)

	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning seuss function as an array...")
	fmt.Println("\n*************************************************")
	fmt.Println("")
	words := array.New()
	suess(words)

	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning seuss function as a list...")
	fmt.Println("\n*************************************************")
	fmt.Println("")
	lwords := list.New()
	suess(lwords)
}

func suess(words maps) {
	file, err := os.Open("greenEggs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, word := range strings.Split(scanner.Text(), " ") {
			word = reg.ReplaceAllString(word, "")
			if word == "" {
				continue
			}
			if words.Contains(strings.ToLower(word)) {
				count, ok := words.Get(strings.ToLower(word)).(int)
				if !ok {
					continue
				}
				words.Add(strings.ToLower(word), count+1)
				continue
			}
			words.Add(strings.ToLower(word), 1)
		}
	}
	fmt.Println(words)
}

func driver(ages maps) {
	// creating map
	ages.Add("Maurice", 20)
	ages.Add("Sylvia", 18)
	fmt.Println("\nHere's the group\n", ages)

	// adding key not in map
	ages.Add("Leo", 7)
	fmt.Println("\nLeo is joining the group\n", ages)

	// adding key already in map (= replacing value)
	fmt.Println("\nSylvia just had a birthday")
	fmt.Println("Her old age was", ages.Add("Sylvia", 19))
	fmt.Println(ages)

	// printing keys
	fmt.Println("\nHere are the people in the group")
	fmt.Println(ages.GetKeys())

	// searching for key
	fmt.Printf("\nIs Sylvia in the group? ")
	if ages.Contains("Sylvia") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	fmt.Println(ages)

	// removing pair
	fmt.Println("\nSylvia has moved away")
	fmt.Println("Her age when she left", ages.Remove("Sylvia"))
	fmt.Println(ages)
	fmt.Printf("Is Sylvia in the group? ")
	if ages.Contains("Sylvia") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	// trying to remove pair that isn't in map
	fmt.Println("Trying to remove Sylvia again")
	fmt.Println(ages.Remove("Sylvia"))

	// getting value for key
	fmt.Println("\nHow old is Maurice?", ages.Get("Maurice"))

	// trying to add nil value
	fmt.Println("\nTrying to add Kyle (with nil age)")
	fmt.Println(ages.Add("Kyle", nil))
}
