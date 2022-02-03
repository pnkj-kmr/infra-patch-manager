package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/pnkj-kmr/infra-patch-manager/remote"
)

// setting up the color for terminal output
var greenText func(a ...interface{}) string = color.New(color.FgHiGreen).SprintFunc()
var redText func(a ...interface{}) string = color.New(color.FgHiRed).SprintFunc()
var yellowText func(a ...interface{}) string = color.New(color.FgYellow).SprintFunc()
var yellowPrintf func(f string, a ...interface{}) = color.New(color.FgYellow).PrintfFunc()

// CLI - helps to get the
type CLI interface {
	DefaultHelp()
	GetRemotes(*string, *string) []remote.Remote
	GetRemoteApps(remote.Remote, *string, *string) []remote.App
}

// NewCLIHander - get a hander of cli
func NewCLIHander(c *flag.FlagSet, s string) CLI {
	c.Parse(os.Args[2:])
	return &_cli{c, s}
}

type _cli struct {
	cmd     *flag.FlagSet
	helpMsg string
}

func (c *_cli) DefaultHelp() {
	fmt.Println("Infra-Patch-Manager subcommand <", c.cmd.Name(), "> holds below actions.", c.helpMsg)
	c.cmd.PrintDefaults()
	fmt.Printf("\n\n")
}

func (c *_cli) GetRemotes(name, rtype *string) (r []remote.Remote) {
	if *name != "" {
		rr, err := remote.NewRemote(*name)
		if err != nil {
			yellowPrintf("Given remote name does not exists. refer conf/remotes.json\n\n")
			os.Exit(0)
		}
		r = append(r, rr)
	} else if *rtype != "" {
		r = remote.GetRemotesByType(*rtype)
	} else {
		r = remote.GetRemotes()
	}
	defaultRemoteCheck(r)
	return
}

func (c *_cli) GetRemoteApps(r remote.Remote, name, apptype *string) (a []remote.App) {
	if *name != "" {
		app, err := r.App(*name)
		if err != nil {
			yellowPrintf("\n\nGiven remote application name does not exists. refer conf/remotes.json\n\n")
			os.Exit(0)
		}
		a = append(a, app)
	} else if *apptype != "" {
		apps, err := r.AppByType(*apptype)
		if err != nil {
			yellowPrintf("\n\nInvalid type. refer conf/remotes.json\n\n")
			os.Exit(0)
		}
		a = apps
	} else {
		apps, err := r.Apps()
		if err != nil {
			yellowPrintf("Internal error. refer conf/remotes.json")
			os.Exit(0)
		}
		a = apps
	}
	return
}

// DefaultHelp - print all helps
func DefaultHelp() {
	fmt.Printf("Infra-Patch-Manager contains the following subcommands set.\n\n")
	fmt.Println(greenText("	remote"), "		| list or search a remote detail with reachablity")
	fmt.Println(greenText("	rights"), "		| read/write rights check on a remote's application(s)")
	fmt.Println(greenText("	upload"), "		| upload a patch to remote")
	fmt.Println(greenText("	extract"), "	| untaring a tar.gz file on relative remote")
	fmt.Println(greenText("	apply"), "		| applying a patch to relative remote application(s)")
	fmt.Println(greenText("	verify"), "		| helps to validate an applied patch")
	fmt.Println(greenText("	exec"), "		| Helps to execute cmd on remote(s)")
	fmt.Print("\n\n")
}

func defaultRemoteCheck(r []remote.Remote) {
	if len(r) == 0 {
		fmt.Printf("Infra-Patch-Manager contains the subcommands set.\n\n")
		yellowPrintf("	- No Remote configured. check conf/remotes.json\n\n\n")
		os.Exit(0)
	}
}

// DefaultCheck - helps to check basic
func DefaultCheck() {
	if len(os.Args) < 2 {
		fmt.Printf("Infra-Patch-Manager contains the subcommands set. Follow help to know more.\n\n")
		fmt.Printf("	-help		| to know more about subcommands\n\n")
		os.Exit(0)
	}
}

func iif(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
