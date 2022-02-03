package cli

import (
	"flag"
	"fmt"

	"github.com/pnkj-kmr/infra-patch-manager/master"
	"github.com/pnkj-kmr/infra-patch-manager/remote"
)

// HandleVerify - handler
func HandleVerify(cmd *flag.FlagSet) {
	remoteName := cmd.String("remote", "", "Remote by it's name")
	remoteType := cmd.String("remote-type", "", "Remote by it's type")
	remoteAll := cmd.Bool("remote-all", false, "All available remotes")
	appName := cmd.String("app", "", "Application by it's name")
	appType := cmd.String("app-type", "", "Application by it's type")
	appAll := cmd.Bool("app-all", false, "All available remote applications")
	// getting a handler
	cliHandler := NewCLIHander(cmd, "Try remote and app combination together.")

	fmt.Println()
	if *remoteAll || *remoteType != "" || *remoteName != "" {
		remotes := cliHandler.GetRemotes(remoteName, remoteType)
		if *appAll || *appType != "" || *appName != "" {
			for _, r := range remotes {
				existingApps := cliHandler.GetRemoteApps(r, appName, appType)
				pm, err := master.NewPatchMaster(r.Name(), false)
				if err == nil {
					apps, err := pm.VerifyFrom(existingApps)
					if err != nil {
						r.UpdateStatus(false)
					} else {
						r.UpdateStatus(true)
					}
					printVerifyWithApps(r, existingApps, apps)
				}
				fmt.Println()
			}
		} else {
			cliHandler.DefaultHelp()
		}
	} else {
		cliHandler.DefaultHelp()
	}
}

func printVerifyWithApps(r remote.Remote, ex []remote.App, apps []remote.App) {
	fmt.Printf("Name		: %s [%s]		%s\n", r.Name(), r.Type(), iif(r.Status(), "OK", "NOT REACHABLE"))
	fmt.Printf("Applications	: %d [requested: %d]\n", len(apps), len(ex))
	if len(ex) == 0 {
		fmt.Printf("	No application found. To more refer conf/remotes.json")
	}
	for i, a := range apps {
		fmt.Printf("[%d]	%s	: %s		%s\n", i+1, a.Name(), a.SourcePath(), iif(a.Status(), "VERIFIED", "NOT VERIFIED"))
		for j, f := range a.GetFiles() {
			fmt.Printf("		: [%d] - %s [%d]	- %s\n", j+1, f.Path(), f.Size(), f.ModTime().Local().String())
		}
	}
}
