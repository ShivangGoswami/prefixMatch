package main

import (
	"flag"
	"log"

	fr "github.com/ShivangGoswami/prefixesMatch/filereader"
	mat "github.com/ShivangGoswami/prefixesMatch/matcher"
)

func main() {
	//Init Variables
	var preflixFilePath, inputString string
	//flag declaration
	flag.StringVar(&preflixFilePath, "prefix-source-file", "", "source file for getting the prefixes")
	flag.StringVar(&inputString, "input-string", "", "Input String that has to matched aganist all availble prefixes")
	//flag parsing
	flag.Parse()

	//error check for input string
	checkEmpty(inputString, "Input String")
	checkEmpty(preflixFilePath, "Prefix File Path")

	//Read Data Value from source files
	prefixDataStore := fr.ReadFile(preflixFilePath) //"./source/sample_prefixes.txt")
	log.Println("Prefix source contains:", len(prefixDataStore), "Input string:", inputString)
	result := mat.Matcher(prefixDataStore, inputString)
	if result.Prefix == "" {
		log.Println("Reuslt: No Match found!")
	} else {
		log.Println("Result: " + result.Prefix + " is the longest prefix found!")
	}
}

func checkEmpty(input string, flag string) {
	if input == "" {
		log.Fatal(flag + " can't be empty!")
	}
}
