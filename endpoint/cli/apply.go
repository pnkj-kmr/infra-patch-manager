package cli

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/pnkj-kmr/infra-patch-manager/master"
	"github.com/pnkj-kmr/infra-patch-manager/master/remote"
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
					apps, _ := pm.PatchTo(existingApps)
					if len(apps) > 0 {
						r.UpdateStatus(true)
					} else {
						r.UpdateStatus(false)
					}
					printApplyWithApps(r, existingApps, apps)
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

func printApplyWithApps(r remote.Remote, ex []remote.App, apps []remote.App) {
	fmt.Println()
	format := "%v\t%v\t%v\t\t\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Remote name", fmt.Sprintf("%s [%s]", r.Name(), r.Type()), "", iif(r.Status(), iif(len(ex) == len(apps), greenText("...OK"), yellowText("...PARTILLY APPLIED")), redText("...NOT REACHABLE")))
	fmt.Fprintf(tw, format, "Applications", fmt.Sprintf("%d [requested: %d]", len(apps), len(ex)), "", "")
	if len(ex) == 0 {
		fmt.Fprintf(tw, format, "", yellowText("No application found. To more refer conf/remotes.json"), "", "")
	} else if len(apps) == 0 {
		fmt.Fprintf(tw, format, "", yellowText("No application(s) reachable"), "", "")
	}
	for _, a := range apps {
		fmt.Fprintf(tw, format, fmt.Sprintf("%s [%s]", a.Name(), a.Type()), a.SourcePath(), "", iif(a.Status(), greenText("APPLIED"), redText("NOT APPLIED")))
		for j, f := range a.GetFiles() {
			fmt.Fprintf(tw, format, "", fmt.Sprintf("[%d] %s [%d]", j+1, f.Path(), f.Size()), f.ModTime().Local().Format(time.Kitchen), "")
		}
	}
	tw.Flush()
}
