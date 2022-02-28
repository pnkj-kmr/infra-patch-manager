package cli

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/pnkj-kmr/infra-patch-manager/entity"
	"github.com/pnkj-kmr/infra-patch-manager/master"
	"github.com/pnkj-kmr/infra-patch-manager/master/remote"
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
	format := "%v\t%v\t\t\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Remote name", fmt.Sprintf("%s [%s]", r.Name(), r.Type()), iif(r.Status(), greenText("...OK"), redText("...NOT REACHABLE")))
	fmt.Fprintf(tw, format, "Extract", name, iif(ok, greenText("EXTRACTED"), redText("FAILED")))
	if ok {
		for i, f := range files {
			fmt.Fprintf(tw, format, "", fmt.Sprintf("[%d]: %s [%d]", i+1, yellowText(f.Path()), f.Size()), "")
		}
	} else {
		fmt.Fprintf(tw, format, "", redText("Unable to extract"), "")
	}
	tw.Flush()
}

func printExtractList(remotes []remote.Remote) {
	fmt.Println()
	format := "%v\t%v\t\t\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	for _, r := range remotes {
		out := extractRemoteList(r)
		fmt.Fprintf(tw, format, "Remote name", fmt.Sprintf("%s [%s]", r.Name(), r.Type()), iif(r.Status(), greenText("...OK"), redText("...NOT REACHABLE")))
		if len(out) > 0 {
			for i, f := range out {
				fmt.Fprintf(tw, format, iif(i == 0, "List output", ""), fmt.Sprintf("[%d] %s", i+1, yellowText(f)), "")
			}
		} else {
			fmt.Fprintf(tw, format, "", redText("Unable to list uploaded files"), "")
		}
		fmt.Fprintf(tw, format, "", "", "")
	}
	tw.Flush()
	fmt.Println()
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
