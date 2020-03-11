package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
	"github.com/tamada/uniq2"
)

/*
VERSION shows version of uniq2.
*/
const VERSION = "1.0.3"

func printHelp(appName string) {
	fmt.Printf(`%s [OPTIONS] [INPUT [OUTPUT]]
OPTIONS
    -a, --adjacent        delete only adjacent duplicated lines.
    -d, --delete-lines    only prints deleted lines.
    -i, --ignore-case     case sensitive.
	-s, --show-counts     show counts.
    -h, --help            print this message.

INPUT                     gives file name of input.  If argument is single dash ('-')
                          or absent, the program read strings from stdin.
OUTPUT                    represents the destination.
`, appName)
}

func printError(err error, statusCode int) int {
	if err == nil {
		return 0
	}
	fmt.Println(err.Error())
	return statusCode
}

func perform(flags *flag.FlagSet, opts *uniq2.Parameters) int {
	var args, err = uniq2.NewArguments(flags.Args()[1:])
	if err != nil {
		return printError(err, 1)
	}
	defer args.Close()
	err = args.Perform(opts)
	return printError(err, 2)
}

func goMain() int {
	// defer profile.Start(profile.ProfilePath(".")).Stop()

	var flags, opts = buildFlagSet()
	var err = flags.Parse(os.Args)
	if err == nil {
		return perform(flags, opts)
	}
	fmt.Println(err.Error())

	return 0
}

func buildFlagSet() (*flag.FlagSet, *uniq2.Parameters) {
	var opts = uniq2.Parameters{}
	var flags = flag.NewFlagSet("uniq2", flag.ContinueOnError)
	flags.Usage = func() { printHelp("uniq2") }
	flags.BoolVarP(&opts.Adjacent, "adjacent", "a", false, "delete only the adjacent duplicate lines")
	flags.BoolVarP(&opts.DeleteLines, "delete-lines", "d", false, "only prints deleted lines")
	flags.BoolVarP(&opts.IgnoreCase, "ignore-case", "i", false, "case sensitive")
	return flags, &opts
}

func main() {
	var exitStatus = goMain()
	os.Exit(exitStatus)
}
