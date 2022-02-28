package cli

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/pnkj-kmr/infra-patch-manager/master"
	"github.com/pnkj-kmr/infra-patch-manager/master/remote"
)

// HandleRights - hangle all remote-app rights related detail
func HandleRights(cmd *flag.FlagSet) {
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
					apps, err := pm.RightsCheckFor(existingApps)
					if err != nil {
						r.UpdateStatus(false)
					} else {
						r.UpdateStatus(true)
					}
					printRemoteWithApps(r, existingApps, apps)
				}
				fmt.Println()
			}
			fmt.Println()
		} else {
			cliHandler.DefaultHelp()
		}
	} else {
		cliHandler.DefaultHelp()
	}
}

func printRemoteWithApps(r remote.Remote, ex []remote.App, apps []remote.App) {
	fmt.Println()
	format := "%v\t%v\t%v\t\t\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Remote name", fmt.Sprintf("%s [%s]", r.Name(), r.Type()), "", iif(r.Status(), greenText("...OK"), redText("...NOT REACHABLE")))
	fmt.Fprintf(tw, format, "Applications", fmt.Sprintf("%d [requested: %d]", len(apps), len(ex)), "", "")
	if len(ex) == 0 {
		fmt.Fprintf(tw, format, "", yellowText("No application found. To more refer conf/remotes.json"), "", "")
	} else if len(apps) == 0 {
		fmt.Fprintf(tw, format, "", yellowText("No application(s) reachable"), "", "")
	}
	for i, a := range apps {
		fmt.Fprintf(tw, format, "", fmt.Sprintf("[%d] %s [%s]", i+1, a.Name(), a.Type()), a.SourcePath(), iif(a.Status(), greenText("OK"), redText("NO R/W RIGHTS")))
	}
	tw.Flush()
}
