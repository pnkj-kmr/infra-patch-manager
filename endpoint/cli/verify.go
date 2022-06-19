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

	if *remoteAll || *remoteType != "" || *remoteName != "" {
		remotes := cliHandler.GetRemotes(remoteName, remoteType)
		if *appAll || *appType != "" || *appName != "" {
			verifyFunc(cliHandler, remotes, appName, appType)
		} else {
			cliHandler.DefaultHelp()
		}
	} else {
		cliHandler.DefaultHelp()
	}
}

func verifyFunc(cli CLI, remotes []remote.Remote, appName, appType *string) {
	fmt.Println()
	i, t := 0, len(remotes)
	type verifyChan struct {
		r      remote.Remote
		exApps []remote.App
		apps   []remote.App
	}
	result := make(chan verifyChan, t)
	for _, r := range remotes {
		go func(r remote.Remote, result chan<- verifyChan) {
			existingApps := cli.GetRemoteApps(r, appName, appType)
			rr, apps := verifyNow(r, existingApps)
			result <- verifyChan{r: rr, exApps: existingApps, apps: apps}
		}(r, result)
	}
	for ch := range result {
		verifyPrint(ch.r, ch.exApps, ch.apps)
		i++
		if i == t {
			close(result)
		}
	}
	fmt.Println()
}

func verifyNow(r remote.Remote, existingApps []remote.App) (remote.Remote, []remote.App) {
	pm, err := master.NewPatchMaster(r.Name(), false)
	var rApps []remote.App
	if err == nil {
		apps, _ := pm.VerifyFrom(existingApps)
		if len(apps) > 0 {
			r.UpdateStatus(true)
		} else {
			r.UpdateStatus(false)
		}
		rApps = apps
	}
	return r, rApps
}

func verifyPrint(r remote.Remote, ex []remote.App, apps []remote.App) {
	fmt.Println()
	format := "%v\t%v\t%v\t\t\t%v\t\n"
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Remote name", fmt.Sprintf("%s [%s]", r.Name(), r.Type()), "", iif(r.Status(), iif(len(ex) == len(apps), greenText("...OK"), yellowText("...PARTILLY VERIFIED")), redText("...NOT REACHABLE")))
	fmt.Fprintf(tw, format, "Applications", fmt.Sprintf("%d [requested: %d]", len(apps), len(ex)), "", "")
	if len(ex) == 0 {
		fmt.Fprintf(tw, format, "", yellowText("No application found. To more refer conf/remotes.json"), "", "")
	} else if len(apps) == 0 {
		fmt.Fprintf(tw, format, "", yellowText("No application(s) reachable"), "", "")
	}
	for _, a := range apps {
		fmt.Fprintf(tw, format, fmt.Sprintf("%s [%s]", a.Name(), a.Type()), a.SourcePath(), "", iif(a.Status(), greenText("VERIFIED"), redText("NOT VERIFIED")))
		for j, f := range a.GetFiles() {
			fmt.Fprintf(tw, format, "", fmt.Sprintf("[%d] %s [%d]", j+1, f.Path(), f.Size()), f.ModTime().Local().Format(time.Kitchen), "")
		}
	}
	tw.Flush()
	fmt.Println()
}
