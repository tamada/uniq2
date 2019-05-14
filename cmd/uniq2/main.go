package main

import (
	"fmt"
	"os"

	"github.com/tamada/uniq2/lib"
	"github.com/urfave/cli"
)

const VERSION = "0.2.0"

func printHelp(app *cli.App) {
	fmt.Printf(`%s
OPTIONS
    -a, --adjacent        delete only adjacent duplicated lines.
    -c, --show-counts     show counts of deleted lines.
    -d, --delete-lines    only prints deleted lines.
    -i, --ignore-case     case sensitive.
    -h, --help            print this message.

INPUT                     gives file name of input.  If argument is single dash ('-')
                          or absent, the program read strings from stdin.
OUTPUT                    represents the destination.
`, app.Usage)
}

func buildFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "adjacent, a",
			Usage: "delete only adjacent duplicated lines.",
		},
		cli.BoolFlag{
			Name:  "show-counts, c",
			Usage: "show counts of deleted lines.",
		},
		cli.BoolFlag{
			Name:  "delete-lines, d",
			Usage: "only prints deleted lines.",
		},
		cli.BoolFlag{
			Name:  "ignore-case, i",
			Usage: "case sensitive",
		},
	}
}

func constructApp() *cli.App {
	var app = cli.NewApp()
	app.Name = "uniq2"
	app.Usage = "Eliminates duplicated lines"
	app.UsageText = "uniq2 [OPTIONS] [INPUT [OUTPUT]]"
	app.Version = VERSION
	app.Flags = buildFlags()
	app.Action = func(c *cli.Context) error {
		return action(app, c)
	}
	return app
}

func parseOptions(c *cli.Context) *lib.Options {
	return &lib.Options{
		Adjacent:    c.Bool("adjacent"),
		ShowCounts:  c.Bool("show-counts"),
		DeleteLines: c.Bool("delete-lines"),
		IgnoreCase:  c.Bool("ignore-case"),
	}
}

func perform(args *lib.Arguments) error {
	defer args.Close()
	return args.Perform()
}

func action(app *cli.App, c *cli.Context) error {
	var options = parseOptions(c)
	if c.Bool("help") {
		printHelp(app)
		return nil
	}
	var args, err = lib.NewArguments(options, c.Args())
	if err != nil {
		return err
	}
	return perform(args)
}

func goMain() int {
	var app = constructApp()
	var err = app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	return 0
}

func main() {
	// separates main function in order to run defers before exit.
	var exitStatus = goMain()
	os.Exit(exitStatus)
}
