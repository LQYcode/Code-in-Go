package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	linesYesterday := "50"
	linesToday := "108"
	//linesMore := linesToday - linesYesterday
	//fmt.Println(linesMore)

	yesterday, err := strconv.Atoi(linesYesterday)
	if err != nil {
		log.Fatal(err)
	}

	today, err := strconv.Atoi(linesToday)
	if err != nil {
		log.Fatal(err)
	}

	lineMore := today - yesterday
	fmt.Println(lineMore)

}
