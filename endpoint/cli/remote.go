package cli

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/pnkj-kmr/infra-patch-manager/master"
	"github.com/pnkj-kmr/infra-patch-manager/master/remote"
)

// HandleRemote - handler for remote subcmd
func HandleRemote(cmd *flag.FlagSet) {
	remoteAll := cmd.Bool("all", false, "List of remotes <remotes.json>")
	remoteName := cmd.String("name", "", "Remote detail by -name <name>")
	remoteType := cmd.String("type", "", "Remote detail by -type <xy>")
	remoteStatus := cmd.Bool("ping", false, "Remote detail with ping status")
	// getting a handler
	cliHandler := NewCLIHander(cmd, "")

	if *remoteAll || *remoteType != "" || *remoteName != "" {
		remotes := cliHandler.GetRemotes(remoteName, remoteType)
		printRemotes(remotes, *remoteStatus)
	} else {
		cliHandler.DefaultHelp()
	}
}

func printRemotes(remotes []remote.Remote, s bool) {
	fmt.Println()
	format := "%v\t%v\t\t\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	for _, r := range remotes {
		if s {
			updateRemoteStatus(r)
		}
		apps, _ := r.Apps()
		fmt.Fprintf(tw, format, "Remote Name", fmt.Sprintf("%s [%s]", r.Name(), r.Type()), iif(s, iif(r.Status(), greenText("...OK"), redText("...NOT REACHABLE")), yellowText("...STATUS NOT CHECKED")))
		fmt.Fprintf(tw, format, "Agent Address", r.AgentAddress(), "")
		fmt.Fprintf(tw, format, "Applications", len(apps), "")
		for i, a := range apps {
			fmt.Fprintf(tw, format, "", fmt.Sprintf("[%d] %s [%s] : %s", i+1, a.Name(), a.Type(), a.SourcePath()), "")
		}
		fmt.Fprintf(tw, format, "", "", "")
	}
	tw.Flush()
}

func updateRemoteStatus(r remote.Remote) {
	pm, err := master.NewPatchMaster(r.Name(), false)
	if err == nil {
		ok, err := pm.Ping()
		if err != nil {
			r.UpdateStatus(false)
		} else {
			r.UpdateStatus(ok)
		}
	}
}
