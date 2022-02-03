package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/pnkj-kmr/infra-patch-manager/entity"
	"github.com/pnkj-kmr/infra-patch-manager/master"
	"github.com/pnkj-kmr/infra-patch-manager/remote"
)

// HandleExtract - handler
func HandleExtract(cmd *flag.FlagSet) {
	remoteName := cmd.String("remote", "", "Remote by it's name")
	remoteType := cmd.String("remote-type", "", "Remote by it's type")
	remoteAll := cmd.Bool("remote-all", false, "All available remotes")
	filename := cmd.String("file", "", "File name to untar [i.e. --file patchXX.tar.gz]")
	// getting a handler
	cliHandler := NewCLIHander(cmd, "try remote v/s file combination")

	fmt.Println()
	if *filename == "" {
		cliHandler.DefaultHelp()
		os.Exit(0)
	}
	if *remoteAll || *remoteType != "" || *remoteName != "" {
		remotes := cliHandler.GetRemotes(remoteName, remoteType)
		extractToRemote(remotes, *filename)
	} else {
		cliHandler.DefaultHelp()
	}
}

func extractToRemote(allRemotes []remote.Remote, f string) {
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
	fmt.Printf("Name		: %s [%s]	-- %s\n", r.Name(), r.Type(), iif(r.Status(), "OK", "NOT REACHABLE"))
	fmt.Printf("Extract		: %s		%s\n", name, iif(ok, "EXTRACTED", "FAILED"))
	if ok {
		for i, f := range files {
			fmt.Printf("[%d]		  %s [%d] - %s\n", i+1, f.Name(), f.Size(), f.Path())
		}
	}
}
