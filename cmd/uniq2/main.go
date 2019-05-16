package main

import (
	"fmt"
	"os"

	flag "github.com/ogier/pflag"
	"github.com/tamada/uniq2/lib"
)

const VERSION = "1.0.0"

func printHelp(appName string) {
	fmt.Printf(`%s [OPTIONS] [INPUT [OUTPUT]]
OPTIONS
    -a, --adjacent        delete only adjacent duplicated lines.
    -d, --delete-lines    only prints deleted lines.
    -i, --ignore-case     case sensitive.
    -h, --help            print this message.

INPUT                     gives file name of input.  If argument is single dash ('-')
                          or absent, the program read strings from stdin.
OUTPUT                    represents the destination.
`, appName)
}

func perform(flags *flag.FlagSet, opts *lib.Options) int {
	var args, err = lib.NewArguments(opts, flags.Args()[1:])
	defer args.Close()
	if err == nil {
		err = args.Perform()
	}
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	return 0
}

func goMain() int {
	var flags, opts = buildFlagSet()
	var err = flags.Parse(os.Args)
	if err == nil {
		return perform(flags, opts)
	}
	fmt.Println(err.Error())

	return 0
}

func buildFlagSet() (*flag.FlagSet, *lib.Options) {
	var opts = lib.Options{}
	var flags = flag.NewFlagSet("uniq2", flag.ContinueOnError)
	flags.Usage = func() { printHelp("uniq2") }
	flags.BoolVarP(&opts.Adjacent, "adjacent", "a", false, "delete only the adjacent duplicate lines")
	flags.BoolVarP(&opts.DeleteLines, "delete-lines", "d", false, "only prints deleted lines")
	flags.BoolVarP(&opts.IgnoreCase, "ignore-case", "i", false, "case sensitive")
	return flags, &opts
}

func main() {
	// separates main function in order to run defers before exit.
	var exitStatus = goMain()
	os.Exit(exitStatus)
}
