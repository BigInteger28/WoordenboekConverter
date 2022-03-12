package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var klinkers string = "aeiouy"
var medeklinkers string = "bcdfghjklmnpqrstvwxz"
var letters string = "abcdefghijklmnopqrstuvwxyz"
var Woordenboek []string
var output []string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func maakWoordenboek() {
	file, err := os.Open("./woordenboek.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		Woordenboek = append(Woordenboek, scanner.Text())
	}
}

func shuffleWordsToFile() {
	file, err := os.Create("output.txt")
	check(err)
	for _, word := range output {
		fmt.Fprintln(file, word)
		check(err)
	}
	err = file.Close()
}

func naarWoord(nummer int) string {
	var woord string
	if nummer < 156 {
		woord += string(letters[nummer/6])
		woord += string(klinkers[nummer%6])
	} else if nummer < 4212 {
		nummer -= 156
		woord += string(letters[nummer/(26*6)])
		woord += string(klinkers[(nummer/26)%6])
		woord += string(letters[nummer%26])
	} else if nummer < 85488 {

	}
	woord += " "
	return woord
}

func zoekPositieNederlandsWoord(woord string) int {
	var positie int = -1
	for huidigWoord := 0; huidigWoord < len(Woordenboek); huidigWoord++ {
		if woord == Woordenboek[huidigWoord] {
			positie = huidigWoord
			goto result
		}
	}
result:
	return positie
}

func main() {
	var startTime int64 = time.Now().UnixMilli()
	maakWoordenboek()
	rand.Seed(time.Now().UnixNano())
	output = make([]string, 199359)
	var randPos int
	var lengte int = 199359
	for i := 0; i < 199359; i++ {
		randPos = rand.Intn(lengte)
		output[i] = Woordenboek[randPos]
		Woordenboek[randPos] = Woordenboek[lengte-1]
		Woordenboek = Woordenboek[:lengte-1]
		lengte--
	}
	shuffleWordsToFile()
	var verstrekenTijd int64 = time.Now().UnixMilli() - startTime
	fmt.Println(verstrekenTijd, "ms")
}
