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
			applyFunc(cliHandler, remotes, appName, appType)
		} else {
			cliHandler.DefaultHelp()
		}
	} else {
		cliHandler.DefaultHelp()
	}
}

func applyFunc(cli CLI, remotes []remote.Remote, appName, appType *string) {
	fmt.Println()
	i, t := 0, len(remotes)
	type applyChan struct {
		r      remote.Remote
		exApps []remote.App
		apps   []remote.App
	}
	result := make(chan applyChan, t)
	for _, r := range remotes {
		go func(r remote.Remote, result chan<- applyChan) {
			existingApps := cli.GetRemoteApps(r, appName, appType)
			rr, apps := applyNow(r, existingApps)
			result <- applyChan{r: rr, exApps: existingApps, apps: apps}
		}(r, result)
	}
	for ch := range result {
		applyPrint(ch.r, ch.exApps, ch.apps)
		i++
		if i == t {
			close(result)
		}
	}
	fmt.Println()
}

func applyNow(r remote.Remote, existingApps []remote.App) (remote.Remote, []remote.App) {
	pm, err := master.NewPatchMaster(r.Name(), false)
	var rApps []remote.App
	if err == nil {
		apps, _ := pm.PatchTo(existingApps)
		if len(apps) > 0 {
			r.UpdateStatus(true)
		} else {
			r.UpdateStatus(false)
		}
		rApps = apps
	}
	return r, rApps
}

func applyPrint(r remote.Remote, ex []remote.App, apps []remote.App) {
	LoaderSkip()
	format := "%v\t%v\t%v\t\t\t%v\t\n"
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
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
	fmt.Println()
}
