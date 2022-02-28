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

// HandleUpload - handler for remote subcmd
func HandleUpload(cmd *flag.FlagSet) {
	remoteName := cmd.String("remote", "", "Remote by it's name")
	remoteType := cmd.String("remote-type", "", "Remote by it's type")
	remoteAll := cmd.Bool("remote-all", false, "All available remotes")
	path := cmd.String("src", "", "File path which will be uploaded")
	// getting a handler
	cliHandler := NewCLIHander(cmd, "Try remote with src input")

	var f entity.File
	if *path != "" {
		f = checkPath(path)
	} else {
		cliHandler.DefaultHelp()
		os.Exit(0)
	}
	if *remoteAll || *remoteType != "" || *remoteName != "" {
		remotes := cliHandler.GetRemotes(remoteName, remoteType)
		uploadToRemotes(remotes, f)
	} else {
		cliHandler.DefaultHelp()
	}
}

func checkPath(path *string) entity.File {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error	: Internal\n")
		os.Exit(0)
	}
	f, err := entity.NewFile(*path, wd)
	if err != nil {
		fmt.Printf("Error	: %s - not exists - %v\n", *path, err)
		os.Exit(0)
	}
	if f.IsDir() {
		fmt.Printf("Error	: %s - cannot upload\n", *path)
		os.Exit(0)
	}
	return f
}

func uploadToRemotes(allRemotes []remote.Remote, f entity.File) {
	fmt.Println()
	for _, r := range allRemotes {
		pm, err := master.NewPatchMaster(r.Name(), false)
		if err == nil {
			uploaded, ok, err := pm.UploadFileToRemote(f)
			if err != nil {
				r.UpdateStatus(false)
			} else {
				r.UpdateStatus(true)
			}
			printRemoteUpload(r, f, uploaded, ok)
			fmt.Println()
		}
	}
	fmt.Println()
}

func printRemoteUpload(r remote.Remote, in entity.File, out entity.Entity, ok bool) {
	format := "%v\t%v\t\t\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Remote name", fmt.Sprintf("%s [%s]", r.Name(), r.Type()), iif(r.Status(), greenText("...OK"), redText("...NOT REACHABLE")))
	fmt.Fprintf(tw, format, "Requested", fmt.Sprintf("%s [%d]", in.Name(), in.Size()), "")
	if ok {
		fmt.Fprintf(tw, format, "Uploaded", fmt.Sprintf("Name - %s", out.Name()), iif(ok, greenText("UPLOADED"), redText("FAILED")))
		fmt.Fprintf(tw, format, "", fmt.Sprintf("Size - %d", out.Size()), "")
		fmt.Fprintf(tw, format, "", fmt.Sprintf("Path - %s", out.Path()), "")
	} else {
		fmt.Fprintf(tw, format, "Uploaded", "", iif(ok, greenText("UPLOADED"), redText("FAILED")))
	}
	tw.Flush()
}
