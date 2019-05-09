package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func foundIn(line string, list []string) bool {
	for _, lineInList := range list {
		if line == lineInList {
			return true
		}
	}
	return false
}

/*
Uniq deletes duplicated items from the given slice.
*/
func Uniq(lines []string) []string {
	var results = []string{}
	for _, line := range lines {
		if !foundIn(line, results) {
			results = append(results, line)
		}
	}
	return results
}

func printAll(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}

func readFromStdin() ([]string, error) {
	var lines = []string{}
	var s = bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines, s.Err()
}

func readFromFile(fileName string) ([]string, error) {
	var lines = []string{}
	var file, err = os.Open(fileName)
	if err != nil {
		return lines, err
	}
	var fin = bufio.NewScanner(file)
	for fin.Scan() {
		var text = strings.Split(fin.Text(), "\n")
		lines = append(lines, text...)
	}
	return lines, nil
}

func execStdin() {
	var lines, err = readFromStdin()
	if err != nil {
		fmt.Println(err)
	} else {
		var results = Uniq(lines)
		printAll(results)
	}
}

func execEachFile(args []string) {
	for _, file := range args {
		var lines, err = readFromFile(file)
		if err != nil {
			fmt.Println(err)
		} else {
			printAll(Uniq(lines))
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		execStdin()
	} else {
		execEachFile(os.Args[1:])
	}
}
