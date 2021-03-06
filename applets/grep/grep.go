package grep

import (
	"flag"

	"fmt"
	"gobox/common"
	"io"
	"log"
	"os"
	"regexp"
)

var (
	flagSet  = flag.NewFlagSet("grep", flag.PanicOnError)
	helpFlag = flagSet.Bool("help", false, "Show this help")
)

func Grep(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() < 1 || *helpFlag {
		println("`grep` <pattern> [<file>...]")
		flagSet.PrintDefaults()
		return nil
	}

	pattern, err := regexp.Compile(flagSet.Arg(0))
	if err != nil {
		log.Fatalf("Invalid regular expression: %s\n", err)
	}

	if flagSet.NArg() == 1 {
		doGrep(pattern, os.Stdin, "<stdin>", false)
	} else {
		for _, fn := range flagSet.Args()[1:] {
			if fh, err := os.Open(fn); err == nil {
				func() {
					defer fh.Close()
					doGrep(pattern, fh, fn, flagSet.NArg() > 2)
				}()
			} else {
				log.Printf("Could not open file %s: %s\n", fn, err)
			}
		}
	}

	return nil
}

func doGrep(pattern *regexp.Regexp, fh io.Reader, fn string, print_fn bool) {
	buf := common.NewBufferedReader(fh)

	for {
		line, err := buf.ReadWholeLine()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Printf("Could not read from %s: %s\n", fn, err)
			return
		}
		if line == "" {
			continue
		}

		if pattern.MatchString(line) {
			if print_fn {
				fmt.Printf("%s:", fn)
			}
			fmt.Printf("%s\n", line)
		}
	}
}
