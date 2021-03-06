package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/fatih/color"
	"github.com/go2c/optparse"
	"github.com/onodera-punpun/prt/config"
	"github.com/onodera-punpun/prt/ports"
	"github.com/onodera-punpun/prt/utils"
)

// Loc prints port locations
func Loc(args []string) {
	// Define valid arguments.
	argd := optparse.Bool("duplicate", 'd', false)
	argn := optparse.Bool("no-alias", 'n', false)
	argh := optparse.Bool("help", 'h', false)

	// Load config.
	conf := config.Load()

	// Parse arguments.
	vals, err := optparse.Parse(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	if *argh {
		fmt.Println("Usage: prt loc [arguments] [ports]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -d,   --duplicate       list duplicate ports as well")
		fmt.Println("  -n,   --no-alias        disable aliasing")
		fmt.Println("  -h,   --help            print help and exit")
		os.Exit(0)
	}

	// This command needs a value.
	if len(vals) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify a port!")
		os.Exit(1)
	}

	// Get all ports.
	all, err := ports.All()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var c []string
	var i int
	for _, p := range vals {
		// Continue if already checked.
		if utils.StringInList(p, c) {
			continue
		}
		// Add to checked ports.
		c = append(c, p)

		// Get port location.
		ll, err := ports.Loc(all, p)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		if !*argd {
			ll = []string{ll[0]}
		}

		var op string
		for _, l := range ll {
			// Alias if needed.
			if !*argn {
				l = ports.Alias(l)
			}

			// Print duplicate indentation.
			if *argd {
				// Reset indentation on new port
				if path.Base(l) != op {
					i = 0
				}
				op = path.Base(l)

				if i > 0 {
					color.Set(conf.DarkColor)
					fmt.Printf(strings.Repeat(conf.IndentChar, i))
					color.Unset()
				}
				i++
			}

			// Finally print the port.
			fmt.Println(l)
		}
	}
}
