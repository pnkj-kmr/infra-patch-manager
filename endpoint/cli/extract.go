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
			extractListFunc(remotes)
		} else if *filename != "" {
			remotes := cliHandler.GetRemotes(remoteName, remoteType)
			extractFunc(remotes, *filename)
		} else {
			cliHandler.DefaultHelp()
		}
	} else {
		cliHandler.DefaultHelp()
	}
}

func extractFunc(remotes []remote.Remote, f string) {
	fmt.Println()
	i, t := 0, len(remotes)
	type extractChan struct {
		r   remote.Remote
		f   string
		out []entity.Entity
		ok  bool
	}
	result := make(chan extractChan, t)
	for _, r := range remotes {
		go func(r remote.Remote, result chan<- extractChan) {
			r, out, ok := extractFile(r, f)
			result <- extractChan{r, f, out, ok}
		}(r, result)
	}
	for ch := range result {
		extractPrint(ch.r, ch.f, ch.out, ch.ok)
		i++
		if i == t {
			close(result)
		}
	}
	fmt.Println()
}

func extractFile(r remote.Remote, f string) (remote.Remote, []entity.Entity, bool) {
	pm, err := master.NewPatchMaster(r.Name(), false)
	var out []entity.Entity
	var ok bool
	if err == nil {
		files, status, err := pm.ExtractFileToRemote("", f)
		if err != nil {
			r.UpdateStatus(false)
		} else {
			r.UpdateStatus(true)
		}
		out = files
		ok = status
	}
	return r, out, ok
}

func extractPrint(r remote.Remote, name string, files []entity.Entity, ok bool) {
	LoaderSkip()
	format := "%v\t%v\t\t\t%v\t\n"
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
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
	fmt.Println()
}

func extractListFunc(remotes []remote.Remote) {
	fmt.Println()
	i, t := 0, len(remotes)
	type extractChan struct {
		r   remote.Remote
		out []string
	}
	result := make(chan extractChan, t)
	for _, r := range remotes {
		go func(r remote.Remote, result chan<- extractChan) {
			r, out := extractList(r)
			result <- extractChan{r, out}
		}(r, result)
	}
	for ch := range result {
		extractListPrint(ch.r, ch.out)
		i++
		if i == t {
			close(result)
		}
	}
	fmt.Println()
}

func extractList(r remote.Remote) (remote.Remote, []string) {
	pm, err := master.NewPatchMaster(r.Name(), false)
	var out []string
	if err == nil {
		out, err = pm.ListAvailablePatches()
		if err == nil {
			r.UpdateStatus(true)
		}
	}
	return r, out
}

func extractListPrint(r remote.Remote, out []string) {
	LoaderSkip()
	format := "%v\t%v\t\t\t%v\t\n"
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Remote name", fmt.Sprintf("%s [%s]", r.Name(), r.Type()), iif(r.Status(), greenText("...OK"), redText("...NOT REACHABLE")))
	if len(out) > 0 {
		for i, f := range out {
			fmt.Fprintf(tw, format, iif(i == 0, "List output", ""), fmt.Sprintf("[%d] %s", i+1, yellowText(f)), "")
		}
	} else {
		fmt.Fprintf(tw, format, "", redText("Unable to list uploaded files"), "")
	}
	fmt.Fprintf(tw, format, "", "", "")
	tw.Flush()
}
