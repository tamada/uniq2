package uniq2

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
Arguments shows data source and destination.
*/
type Arguments struct {
	input  io.Reader
	output io.Writer
}

/*
NewArguments creates an instance of Arguments from given args.
*/
func NewArguments(args []string) (*Arguments, error) {
	input, output, err := parseCliArguments(args)
	if err != nil {
		return nil, err
	}
	return &Arguments{input: input, output: output}, nil
}

/*
Close closes data source and destination.
*/
func (args *Arguments) Close() {
	closeImpl(args.input)
	closeImpl(args.output)
}

func closeImpl(stream interface{}) {
	closer, ok := stream.(io.Closer)
	if ok {
		closer.Close()
	}
}

/*
Perform executes Uniq2 by following the given Parameters.
*/
func (args *Arguments) Perform(opts *Parameters) error {
	uniqer := opts.BuildUniqer()
	return args.performImpl(uniqer)

}

func (args *Arguments) performImpl(uniqer Uniqer) error {
	reader := bufio.NewReader(args.input)
	writer := bufio.NewWriter(args.output)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		if uniqer.StreamLine(line) {
			writer.WriteString(line)
			writer.WriteString("\n")
		}
	}
	writer.Flush()
	return nil
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

func parseCliArguments(args []string) (io.ReadCloser, io.WriteCloser, error) {
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
