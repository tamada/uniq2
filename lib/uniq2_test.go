package lib

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func open(fileName string) *os.File {
	var file, _ = os.Open(fileName)
	return file
}

func TestSample(t *testing.T) {
	var testdata = []struct {
		message  string
		target   io.Reader
		fileFlag bool
	}{
		{"stdin", os.Stdin, true},
		{"file", open("../testdata/test1.txt"), true},
		{"buffer", &bytes.Buffer{}, false},
	}

	for _, td := range testdata {
		var _, ok = td.target.(*os.File)
		if ok != td.fileFlag {
			t.Errorf("conversion error: %v", td)
		}
	}
}

func deleteFile(path string) {
	os.Remove(path)
}

func TestOpenFile(t *testing.T) {
	var testdata = []struct {
		args       []string
		buildError bool
		inputPath  string
		outputPath string
	}{
		{[]string{"../testdata/test1.txt", "../testdata/dest1.txt"}, false, "../testdata/test1.txt", "../testdata/dest1.txt"},
		{[]string{"../testdata/not_exist.txt"}, true, "", ""},
	}
	for _, td := range testdata {
		var args, err = NewArguments(&Options{}, td.args)
		defer args.Close()
		defer deleteFile(td.outputPath)
		if (err == nil) == td.buildError {
			t.Errorf("%v: unexpected build arguments error: %v", td.args, err)
		}
		if err == nil {
			var inputFile = args.Input.(*os.File)
			var outputFile = args.Output.(*os.File)
			if inputFile.Name() != td.inputPath {
				t.Errorf("%v: input did not match, wont: %v, got: %v", td.args, td.inputPath, args.Input)
			}
			if outputFile.Name() != td.outputPath {
				t.Errorf("%v: output did not match, wont: %v, got: %v", td.args, td.outputPath, args.Output)
			}
		}
	}
}

func TestNewArguments(t *testing.T) {
	var testdata = []struct {
		args       []string
		buildError string
		wontInput  io.Reader
		wontOutput io.Writer
	}{
		{[]string{}, "", os.Stdin, os.Stdout},
		{[]string{"-"}, "", os.Stdin, os.Stdout},
		{[]string{"-", "-"}, "", os.Stdin, os.Stdout},
		{[]string{"-", "-", "-"}, fmt.Sprintf("too many arguments: %v", []string{"-", "-", "-"}), nil, nil},
	}

	for _, td := range testdata {
		var args, err = NewArguments(&Options{}, td.args)
		if err != nil && err.Error() != td.buildError {
			t.Errorf("%v: build arguments error: wont: %v, got: %v", td.args, td.buildError, err)
		}
		if err == nil {
			if args.Input != td.wontInput {
				t.Errorf("%v: input did not match, wont: %v, got: %v", td.args, td.wontInput, args.Input)
			}
			if args.Output != td.wontOutput {
				t.Errorf("%v: output did not match, wont: %v, got: %v", td.args, td.wontOutput, args.Output)
			}
		}
	}
}

func TestPerform(t *testing.T) {
	var testdata = []struct {
		opts   *Options
		input  string
		result string
	}{
		{&Options{}, "../testdata/test1.txt", "a1+a2+a3+a4+A1"},
		{&Options{Adjacent: true}, "../testdata/test1.txt", "a1+a2+a3+a4+a1+A1"},
		{&Options{IgnoreCase: true}, "../testdata/test1.txt", "a1+a2+a3+a4"},
		{&Options{Adjacent: true, IgnoreCase: true}, "../testdata/test1.txt", "a1+a2+a3+a4+a1"},
		{&Options{DeleteLines: true}, "../testdata/test1.txt", "a1+a2+a1"},
		{&Options{Adjacent: true, DeleteLines: true}, "../testdata/test1.txt", "a1+a2"},
		{&Options{IgnoreCase: true, DeleteLines: true}, "../testdata/test1.txt", "a1+a2+a1+A1"},
		{&Options{Adjacent: true, IgnoreCase: true, DeleteLines: true}, "../testdata/test1.txt", "a1+a2+A1"},
	}

	for _, td := range testdata {
		var inputFile, _ = os.Open(td.input)
		defer inputFile.Close()
		var output = bytes.Buffer{}
		var args = Arguments{Options: td.opts, Input: inputFile, Output: &output}
		args.Perform()
		var result = convertLnToPlus(output.String())
		if result != td.result {
			t.Errorf("test failed on option %v, wont: %s, got: %s", td.opts, td.result, result)
		}
	}
}

func convertLnToPlus(string string) string {
	var lines = strings.Split(strings.TrimSpace(string), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	return strings.Join(lines, "+")
}
