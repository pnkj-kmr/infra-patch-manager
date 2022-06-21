package cli

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/briandowns/spinner"
	"github.com/pnkj-kmr/infra-patch-manager/master/remote"
)

// CLI - helps to get the
type CLI interface {
	DefaultHelp()
	GetRemotes(*string, *string) []remote.Remote
	GetRemoteApps(remote.Remote, *string, *string) []remote.App
}

type _cli struct {
	cmd     *flag.FlagSet
	helpMsg string
}

// NewCLIHander - get a hander of cli
func NewCLIHander(c *flag.FlagSet, s string) CLI {
	c.Parse(os.Args[2:])
	return &_cli{c, s}
}

func (c *_cli) DefaultHelp() {
	fmt.Println("Infra-Patch-Manager subcommand <", c.cmd.Name(), "> holds below actions.", c.helpMsg)
	fmt.Println()
	c.cmd.PrintDefaults()
	fmt.Printf("\n\n")
}

func (c *_cli) GetRemotes(name, rtype *string) (r []remote.Remote) {
	if *name != "" {
		rr, err := remote.NewRemote(*name)
		if err != nil {
			fmt.Printf("\n\t%s\n\n", yellowText("- Given remote name does not exists. refer conf/remotes.json"))
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
			// fmt.Printf("\n\t%s\n\n", yellowText("- Given remote application name does not exists. refer conf/remotes.json"))
			// os.Exit(0)
			return a
		}
		a = append(a, app)
	} else if *apptype != "" {
		apps, err := r.AppByType(*apptype)
		if err != nil {
			fmt.Printf("\n\t%s\n\n", yellowText("- Invalid type. refer conf/remotes.json"))
			os.Exit(0)
		}
		a = apps
	} else {
		apps, err := r.Apps()
		if err != nil {
			fmt.Printf("\n\t%s\n\n", yellowText("- Internal error. refer conf/remotes.json"))
			os.Exit(0)
		}
		a = apps
	}
	return
}

// DefaultHelp - print all helps
func DefaultHelp() {
	format := "\t\t%v\t| %v\t\n"
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Printf("Infra-Patch-Manager contains the following subcommands set.\n\n")
	fmt.Fprintf(tw, format, greenText("remote"), "list or search a remote detail with reachablity")
	fmt.Fprintf(tw, format, greenText("rights"), "read/write rights check on a remote's application(s)")
	fmt.Fprintf(tw, format, greenText("upload"), "upload a patch to remote")
	fmt.Fprintf(tw, format, greenText("extract"), "untaring a tar.gz file on relative remote")
	fmt.Fprintf(tw, format, greenText("apply"), "applying a patch to relative remote application(s)")
	fmt.Fprintf(tw, format, greenText("verify"), "helps to validate an applied patch")
	fmt.Fprintf(tw, format, greenText("download"), "helps to download the patch from remote")
	fmt.Fprintf(tw, format, greenText("exec"), "helps to execute commands on remote(s)")
	tw.Flush()
	fmt.Print("\n\n")
}

// DefaultCheck - helps to check basic
func DefaultCheck() {
	if len(os.Args) < 2 {
		fmt.Printf("Infra-Patch-Manager contains the subcommands set. Follow help to know more.\n\n")
		format := "\t\t%v\t| %v\t\n"
		tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
		fmt.Fprintf(tw, format, greenText("help"), "to know more about subcommands")
		tw.Flush()
		fmt.Printf("\n\n")
		os.Exit(0)
	}
}

func defaultRemoteCheck(r []remote.Remote) {
	if len(r) == 0 {
		fmt.Printf("Infra-Patch-Manager contains the subcommands set.\n\n")
		fmt.Printf("\t%s\n\n", yellowText("- No Remote configured. check conf/remotes.json"))
		os.Exit(0)
	}
}

func iif(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

// Loader helps to show to running process status
func Loader() *spinner.Spinner {
	// spinSet := spinner.CharSets[26] // ...
	// spinSet := spinner.CharSets[36] // [>>]
	spinSet := []string{".", "..", "...", " ...", "  ...", "   ..."}
	s := spinner.New(spinSet, 100*time.Millisecond)
	// s.Prefix = "running "
	return s
}

// LoaderSkip helps to skip the loader from terminal
func LoaderSkip() {
	fmt.Println("\r                                        ")
}
