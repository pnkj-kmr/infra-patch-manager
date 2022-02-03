package cli

import (
	"flag"
	"fmt"
	"strings"

	"github.com/pnkj-kmr/infra-patch-manager/master"
	"github.com/pnkj-kmr/infra-patch-manager/remote"
)

// HandleRemoteCmd - handler for remote subcmd
func HandleRemoteCmd(cmd *flag.FlagSet) {
	remoteName := cmd.String("remote", "", "Remote by it's name")
	remoteType := cmd.String("remote-type", "", "Remote by it's type")
	remoteAll := cmd.Bool("remote-all", false, "All available remotes")
	execCmd := cmd.String("c", "", "Executable command statement")
	// getting a handler
	cliHandler := NewCLIHander(cmd, "Try remote(s) with c(command) together")

	fmt.Println()
	if *remoteAll || *remoteType != "" || *remoteName != "" {
		remotes := cliHandler.GetRemotes(remoteName, remoteType)
		if *execCmd != "" {
			printRemoteCmd(remotes, *execCmd)
		} else {
			cliHandler.DefaultHelp()
		}
	} else {
		cliHandler.DefaultHelp()
	}
}

func printRemoteCmd(remotes []remote.Remote, s string) {
	for _, r := range remotes {
		out := executeRemoteCmd(r, s)
		fmt.Printf("Name		: %s [%s]		%s\n", r.Name(), r.Type(), iif(r.Status(), "OK", "NOT REACHABLE"))
		fmt.Printf("Execute		: %s\n", s)
		fmt.Println("Output		:")
		if len(out) > 0 {
			fmt.Println(strings.Repeat("-", 60))
			fmt.Printf(string(out))
			fmt.Println(strings.Repeat("-", 60))
		}
		fmt.Println()
	}
}

func executeRemoteCmd(r remote.Remote, s string) (out []byte) {
	pm, err := master.NewPatchMaster(r.Name(), false)
	if err == nil {
		v, err := pm.ExecuteCmdOnRemote(s)
		out = v
		if err != nil {
			r.UpdateStatus(false)
		} else {
			r.UpdateStatus(true)
		}
	}
	return
}
