package cli

import (
	"flag"
	"fmt"

	"github.com/pnkj-kmr/infra-patch-manager/master"
	"github.com/pnkj-kmr/infra-patch-manager/remote"
)

// HandleRemote - handler for remote subcmd
func HandleRemote(cmd *flag.FlagSet) {
	remoteAll := cmd.Bool("all", false, "List of remotes <remotes.json>")
	remoteName := cmd.String("name", "", "Remote detail by -name <name>")
	remoteType := cmd.String("type", "", "Remote detail by -type <xy>")
	remoteStatus := cmd.Bool("status", false, "Remote detail with status")
	// getting a handler
	cliHandler := NewCLIHander(cmd, "")

	fmt.Println()
	if *remoteAll || *remoteType != "" || *remoteName != "" {
		remotes := cliHandler.GetRemotes(remoteName, remoteType)
		printRemotes(remotes, *remoteStatus)
	} else {
		cliHandler.DefaultHelp()
	}
}

func printRemotes(remotes []remote.Remote, s bool) {
	for _, r := range remotes {
		if s {
			updateRemoteStatus(r)
		}
		apps, _ := r.Apps()
		fmt.Printf("Name		: %s [%s]		%s\n", r.Name(), r.Type(), iif(s, iif(r.Status(), "OK", "NOT REACHABLE"), "STATUS NOT CHECKED"))
		fmt.Printf("Agent		: %s\n", r.AgentAddress())
		fmt.Printf("Applications	: %d\n", len(apps))
		for i, a := range apps {
			fmt.Printf("[%d]	%s	: %s\n", i+1, a.Name(), a.SourcePath())
		}
		fmt.Println()
	}
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
