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

func closeImpl(value interface{}) {
	var closer, ok = value.(io.Closer)
	if ok {
		closer.Close()
	}
}

/*
Close finalize the files of Arguments.
*/
func (args *Arguments) Close() {
	closeImpl(args.Input)
	closeImpl(args.Output)
}

/*
NewArguments construct an instance of Arguments with the given parameters.
*/
func NewArguments(opts *Options, args []string) (*Arguments, error) {
	var arguments = Arguments{Options: opts}
	var input, output, err = parseCliArguments(args)
	arguments.Input = input
	arguments.Output = output
	return &arguments, err
}

func parseCliArguments(args []string) (*os.File, *os.File, error) {
	switch len(args) {
	case 0:
		return os.Stdin, os.Stdout, nil
	case 1:
		var input, err = createInput(args[0])
		return input, os.Stdout, err
	case 2:
		var input, output *os.File
		var err error
		input, err = createInput(args[0])
		if err == nil {
			output, err = createOutput(args[1])
		}
		return input, output, err
	}
	return nil, nil, fmt.Errorf("too many arguments: %v", args)
}

func createOutput(output string) (*os.File, error) {
	if output == "-" {
		return os.Stdout, nil
	}
	return os.OpenFile(output, os.O_CREATE|os.O_WRONLY, 0644)
}

func createInput(input string) (*os.File, error) {
	if input == "-" {
		return os.Stdin, nil
	}
	return os.Open(input)
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

func updateDatabase(line string, uniqFlag bool, entries []*entry) []*entry {
	for i, entry := range entries {
		if entry.line == line {
			entries[i].count[uniqFlag]++
		}
	}
	return entries
}

func upsertDatabase(line string, uniqFlag bool, entries []*entry) []*entry {
	if uniqFlag {
		return updateDatabase(line, uniqFlag, entries)
	}
	var entry = &entry{line: line, count: map[bool]int{uniqFlag: 1}}
	return append(entries, entry)
}

func (args *Arguments) runUnique(scanner *bufio.Scanner, writer *bufio.Writer) error {
	var entries = []*entry{}
	for scanner.Scan() {
		var line = scanner.Text()
		var uniqFlag, lineToDB = args.isUniqLine(line, entries)
		entries = upsertDatabase(lineToDB, uniqFlag, entries)
		if isPrint(uniqFlag, args.Options.DeleteLines) {
			writer.WriteString(line)
			writer.WriteString("\n")
		}
	}
	writer.Flush()
	return nil
}

func (opts *Options) match(readLine string, lineOfDB *entry) bool {
	return readLine == lineOfDB.line
}

func (opts *Options) isFoundLineInAdjacentDB(line string, list []*entry) bool {
	if len(list) == 0 {
		return false
	}
	return opts.match(line, list[len(list)-1])
}

func (opts *Options) isFoundLineInDB(line string, list []*entry) bool {
	for _, lineInList := range list {
		if line == lineInList.line {
			return true
		}
	}
	return false
}

func (args *Arguments) isUniqLine(line string, list []*entry) (flag bool, lineToDB string) {
	lineToDB = line
	if args.Options.IgnoreCase {
		lineToDB = strings.ToLower(line)
	}
	if args.Options.Adjacent {
		return args.Options.isFoundLineInAdjacentDB(lineToDB, list), lineToDB
	}
	return args.Options.isFoundLineInDB(lineToDB, list), lineToDB
}
