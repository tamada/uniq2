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
const VERSION = "1.1.1"

func helpMessage(appName string) string {
	return fmt.Sprintf(`%s version %s
%s [OPTIONS] [INPUT [OUTPUT]]
OPTIONS
    -a, --adjacent        delete only adjacent duplicated lines.
    -d, --delete-lines    only prints deleted lines.
    -i, --ignore-case     case sensitive.
    -h, --help            print this message.

INPUT                     gives file name of input.  If argument is single dash ('-')
                          or absent, the program read strings from stdin.
OUTPUT                    represents the destination.`, appName, VERSION, appName)
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

func goMain(args []string) int {
	// defer profile.Start(profile.ProfilePath(".")).Stop()

	var flags, opts = buildFlagSet()
	var err = flags.Parse(args)
	if err != nil {
		return printError(err, 1)
	}
	if opts.helpFlag {
		return printError(fmt.Errorf(helpMessage("uniq2")), 0)
	}
	return perform(flags, opts.params)
}

type options struct {
	params   *uniq2.Parameters
	helpFlag bool
}

func buildFlagSet() (*flag.FlagSet, *options) {
	var opts = options{params: &uniq2.Parameters{}}
	var flags = flag.NewFlagSet("uniq2", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage("uniq2")) }
	flags.BoolVarP(&opts.params.Adjacent, "adjacent", "a", false, "delete only the adjacent duplicate lines")
	flags.BoolVarP(&opts.params.DeleteLines, "delete-lines", "d", false, "only prints deleted lines")
	flags.BoolVarP(&opts.params.IgnoreCase, "ignore-case", "i", false, "case sensitive")
	flags.BoolVarP(&opts.helpFlag, "help", "h", false, "print this message")
	return flags, &opts
}

func main() {
	var exitStatus = goMain(os.Args)
	os.Exit(exitStatus)
}
