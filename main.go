package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
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
	cmd, err := exec.Command("sh", "counter.sh").Output()
	if err != nil {
		panic(err)
	}
	// TODO: Is there a better way to convert a slice of uints ([]uint8) to float64?
	f := string(cmd)
	words, err := strconv.ParseFloat(f, 64)
	if err != nil {
		panic(err)
	}
	return words
}
func main() {
	srcFile, err := os.Open("file.txt")
	if err != nil {
		log.Fatalln("Sorry, there was a problem opening the file", srcFile)
	}
	defer srcFile.Close()

	counts := wordCount(srcFile)
	spaces := counts[" "]
	commas := counts[","]
	periods := counts["."]
	quotes := counts["\""] / 2
	fmt.Println("commas", commas)
	fmt.Println("spaces", spaces)

	for i := 0.0; i < commas/10; i++ {
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
	fmt.Printf("The pause and pace is: %1.3f\n", area)
	fmt.Printf("With a ratio of commas to periods: 1:%1.2f\n", ratio)
	fmt.Println("Questions ? :", counts["?"])
	fmt.Println("Exclamations ! :", counts["!"])
	fmt.Printf("Number of Quotes: %1.0f\n\n", quotes)

}
