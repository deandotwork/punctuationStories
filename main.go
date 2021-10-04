package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Counts struct {
}

func wordCount(rdr io.Reader) map[string]float64 {
	counts := map[string]float64{}
	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		word := scanner.Text()
		counts[word]++
	}
	return counts
}

func totalWords() float64 {
	scanner := bufio.NewScanner(strings.NewReader(srcFile))
	scanner.Split(bufio.ScanWords)
	words := 0.0
	for scanner.Scan() {
		words++
	}
	return words
}

func problem1() {
	got := 1
	want := 1

	assert(got == want)
}

var srcFile string

func main() {
	srcFile, err := os.Open("pride.txt")
	if err != nil {
		log.Fatalln("Sorry, there was a problem opening the file", srcFile)
	}
	defer srcFile.Close()

	counts := wordCount(srcFile)
	fmt.Println(counts)
	spaces := counts[" "]
	commas := counts[","]
	periods := counts["."]
	quotes := counts["\""] / 2
	fmt.Println("commas", commas)
	fmt.Println("spaces", spaces)

	for i := 0.0; i < commas/100; i++ {
		fmt.Printf(",")
	}
	fmt.Println("periods", periods)
	var ratio float64
	if commas >= periods {
		ratio = 1 / (periods / commas)
		//	fmt.Printf("ratio: %1.2f\n", ratio)
	} else if commas <= periods {
		ratio = 1 / (commas / periods)
		//	fmt.Printf("ratio: %1.2f\n", ratio)
	}
	// This measures the pause and pace of a story. Only relevant in relation to another txt?
	area := (commas * periods) / totalWords()
	bookPages := totalWords()
	fmt.Println("Number of Pages:\n", bookPages)
	fmt.Printf("The pause and pace is: %4.3f\n", area)
	fmt.Printf("With a ratio of commas to periods: 1:%1.2f\n", ratio)
	fmt.Println("Questions ? :", counts["?"])
	fmt.Println("Exclamations ! :", counts["!"])
	fmt.Println(quotes)
	fmt.Printf("Number of Quotes: %1.0f\n\n", quotes)

	problem1()
}
