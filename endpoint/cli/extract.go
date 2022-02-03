package cli

import (
	"flag"
	"fmt"
	"strings"

	"github.com/pnkj-kmr/infra-patch-manager/entity"
	"github.com/pnkj-kmr/infra-patch-manager/master"
	"github.com/pnkj-kmr/infra-patch-manager/remote"
)

// HandleExtract - handler
func HandleExtract(cmd *flag.FlagSet) {
	remoteName := cmd.String("remote", "", "Remote by it's name")
	remoteType := cmd.String("remote-type", "", "Remote by it's type")
	remoteAll := cmd.Bool("remote-all", false, "All available remotes")
	filename := cmd.String("f", "", "File name to untar [i.e. -f patchXX.tar.gz]")
	listfiles := cmd.Bool("list", false, "List out all uploaded files")
	// getting a handler
	cliHandler := NewCLIHander(cmd, "Try remote v/s [file or list] combination")

	if *remoteAll || *remoteType != "" || *remoteName != "" {
		if *listfiles {
			remotes := cliHandler.GetRemotes(remoteName, remoteType)
			printExtractList(remotes)
		} else if *filename != "" {
			remotes := cliHandler.GetRemotes(remoteName, remoteType)
			extractToRemote(remotes, *filename)
		} else {
			cliHandler.DefaultHelp()
		}
	} else {
		cliHandler.DefaultHelp()
	}
}

func extractToRemote(allRemotes []remote.Remote, f string) {
	fmt.Println()
	for _, r := range allRemotes {
		pm, err := master.NewPatchMaster(r.Name(), false)
		if err == nil {
			files, ok, err := pm.ExtractFileToRemote("", f)
			if err != nil {
				r.UpdateStatus(false)
			} else {
				r.UpdateStatus(true)
			}
			printRemoteExtract(r, f, files, ok)
			fmt.Println()
		}
	}
	fmt.Println()
}

func printRemoteExtract(r remote.Remote, name string, files []entity.Entity, ok bool) {
	fmt.Printf("Remote name	: %s [%s]	%s\n", r.Name(), r.Type(), iif(r.Status(), greenText("--- OK"), redText("--- NOT REACHABLE")))
	fmt.Printf("Extract		: %s		%s\n", name, iif(ok, greenText("EXTRACTED"), redText("FAILED")))
	if ok {
		for i, f := range files {
			fmt.Printf("[%d]		  %s [%d] - %s\n", i+1, yellowText(f.Name()), f.Size(), f.Path())
		}
	}
}

func printExtractList(remotes []remote.Remote) {
	fmt.Println()
	for _, r := range remotes {
		out := extractRemoteList(r)
		fmt.Printf("Remote name	: %s [%s]		%s\n", r.Name(), r.Type(), iif(r.Status(), greenText("--- OK"), redText("--- NOT REACHABLE")))
		fmt.Println("List output:")
		if len(out) > 0 {
			fmt.Println(strings.Repeat("-", 60))
			for i, f := range out {
				fmt.Printf("[%d]	%s\n", i+1, yellowText(f))
			}
			fmt.Println(strings.Repeat("-", 60))
		}
		fmt.Println()
	}
}

func extractRemoteList(r remote.Remote) (out []string) {
	pm, err := master.NewPatchMaster(r.Name(), false)
	if err == nil {
		out, err = pm.ListAvailablePatches()
		if err == nil {
			r.UpdateStatus(true)
		}
	}
	return
}
