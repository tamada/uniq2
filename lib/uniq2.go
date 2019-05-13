package lib

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
Options represents option parameter values.
*/
type Options struct {
	Adjacent    bool
	ShowCounts  bool
	DeleteLines bool
	IgnoreCase  bool
}

/*
Arguments represents the command line arguments.
*/
type Arguments struct {
	Options *Options
	Input   io.Reader
	Output  io.Writer
}

type entry struct {
	line  string
	count map[bool]int
}

/*
Close finalize the files of Arguments.
*/
func (args *Arguments) Close() {
	var inputFile, ok1 = args.Input.(*os.File)
	if ok1 {
		inputFile.Close()
	}
	var outputFile, ok2 = args.Output.(*os.File)
	if ok2 {
		outputFile.Close()
	}
}

/*
Perform reads files from args.Input and writes result to args.Output.
*/
func (args *Arguments) Perform() error {
	var scanner = bufio.NewScanner(args.Input)
	var writer = bufio.NewWriter(args.Output)

	return args.runUnique(scanner, writer)
}

func isPrint(uniqFlag bool, deleteLineFlag bool) bool {
	return !uniqFlag && !deleteLineFlag ||
		uniqFlag && deleteLineFlag
}

func (args *Arguments) runUnique(scanner *bufio.Scanner, writer *bufio.Writer) error {
	var results = []string{}
	for scanner.Scan() {
		var line = scanner.Text()
		var uniqFlag, lineToDB = args.isUniqLine(line, results)
		fmt.Printf("line: %s, uniq: %v, deletes: %v\n", line, uniqFlag, args.Options.DeleteLines)
		if !uniqFlag {
			results = append(results, lineToDB)
		}
		if isPrint(uniqFlag, args.Options.DeleteLines) {
			writer.WriteString(line)
			writer.WriteString("\n")
		}
	}
	writer.Flush()
	return nil
}

func (opts *Options) match(readLine, lineOfDB string) bool {
	return readLine == lineOfDB
}

func (opts *Options) isFoundLineInAdjacentDB(line string, list []string) bool {
	if len(list) == 0 {
		return false
	}
	return opts.match(line, list[len(list)-1])
}

func (opts *Options) isFoundLineInDB(line string, list []string) bool {
	for _, lineInList := range list {
		if line == lineInList {
			return true
		}
	}
	return false
}

func (args *Arguments) isUniqLine(line string, list []string) (flag bool, lineToDB string) {
	lineToDB = line
	if args.Options.IgnoreCase {
		lineToDB = strings.ToLower(line)
	}
	if args.Options.Adjacent {
		return args.Options.isFoundLineInAdjacentDB(lineToDB, list), lineToDB
	}
	return args.Options.isFoundLineInDB(lineToDB, list), lineToDB
}
