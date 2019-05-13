package lib

import (
	"bufio"
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

/*
UniqueLine shows the lines of unique.
*/
type UniqueLine struct {
	count int
	line  string
}

/*
DeleteLine shows the lines of deletion.
*/
type DeleteLine UniqueLine

/*
Context represents a targets.
*/
type Context struct {
	uniques []UniqueLine
	deletes []DeleteLine
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

func (args *Arguments) runUnique(scanner *bufio.Scanner, writer *bufio.Writer) error {
	var results = []string{}
	for scanner.Scan() {
		var line = scanner.Text()
		if flag, lineToDB := args.foundLine(line, results); !flag {
			writer.WriteString(line)
			writer.WriteString("\n")
			results = append(results, lineToDB)
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

func (args *Arguments) foundLine(line string, list []string) (flag bool, lineToDB string) {
	lineToDB = line
	if args.Options.IgnoreCase {
		lineToDB = strings.ToLower(line)
	}
	if args.Options.Adjacent {
		return args.Options.isFoundLineInAdjacentDB(lineToDB, list), lineToDB
	}
	return args.Options.isFoundLineInDB(lineToDB, list), lineToDB
}
