package cli

import (
	"flag"
	"fmt"

	"github.com/pnkj-kmr/infra-patch-manager/master"
	"github.com/pnkj-kmr/infra-patch-manager/remote"
)

// HandleApply - handle apply
func HandleApply(cmd *flag.FlagSet) {
	remoteName := cmd.String("remote", "", "Remote by it's name")
	remoteType := cmd.String("remote-type", "", "Remote by it's type")
	remoteAll := cmd.Bool("remote-all", false, "All available remotes")
	appName := cmd.String("app", "", "Application by it's name")
	appType := cmd.String("app-type", "", "Application by it's type")
	appAll := cmd.Bool("app-all", false, "All available remote applications")
	// getting a handler
	cliHandler := NewCLIHander(cmd, "Try remote and app combination together.")

	if *remoteAll || *remoteType != "" || *remoteName != "" {
		remotes := cliHandler.GetRemotes(remoteName, remoteType)
		if *appAll || *appType != "" || *appName != "" {
			for _, r := range remotes {
				existingApps := cliHandler.GetRemoteApps(r, appName, appType)
				pm, err := master.NewPatchMaster(r.Name(), false)
				if err == nil {
					apps, err := pm.PatchTo(existingApps)
					if err != nil {
						r.UpdateStatus(false)
					} else {
						r.UpdateStatus(true)
					}
					printApplyWithApps(r, existingApps, apps)
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

func printApplyWithApps(r remote.Remote, ex []remote.App, apps []remote.App) {
	fmt.Println()
	fmt.Printf("Remote name	: %s [%s]		%s\n", r.Name(), r.Type(), iif(r.Status(), greenText("--- OK"), redText("--- NOT REACHABLE")))
	fmt.Printf("Applications	: %d [requested: %d]\n", len(apps), len(ex))
	if len(ex) == 0 {
		yellowPrintf("	No application found. To more refer conf/remotes.json\n\n")
	}
	for i, a := range apps {
		fmt.Printf("[%d]	%s	: [%s] %s		%s\n", i+1, a.Name(), a.Type(), a.SourcePath(), iif(a.Status(), greenText("APPLIED"), redText("NOT APPLIED")))
		for j, f := range a.GetFiles() {
			fmt.Printf("		: [%d] - %s [%d]	- %s\n", j+1, f.Path(), f.Size(), f.ModTime().Local().String())
		}
	}
}
