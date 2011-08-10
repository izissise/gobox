package main

import (
	"os"
)

// Applet imports
import (
	"applets/echo"
	"applets/shell"
	"applets/telnetd"
	"applets/ls"
	"applets/rm"
	"applets/httpd"
)

// This map contains the mappings from callname
// to applet function.
var Applets map[string]Applet = map[string]Applet{
	"echo":    echo.Echo,
	"shell":   shell.Shell,
	"telnetd": telnetd.Telnetd,
	"ls":      ls.Ls,
	"rm":      rm.Rm,
	"httpd":   httpd.Httpd,
}

// Signature of applet functions.
// call is like os.Argv, and therefore contains the
// name of the applet itself in call[0].
// If the returned error is not nil, it is printed
// to stdout.
type Applet func(call []string) os.Error