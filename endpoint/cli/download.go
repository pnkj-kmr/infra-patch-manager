package cli

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"text/tabwriter"

	"github.com/pnkj-kmr/infra-patch-manager/entity"
	"github.com/pnkj-kmr/infra-patch-manager/master"
	"github.com/pnkj-kmr/infra-patch-manager/master/remote"
)

// HandleDownload - handler for remote subcmd
func HandleDownload(cmd *flag.FlagSet) {
	remoteName := cmd.String("remote", "", "Remote detail by -remote <name>")
	downloadFile := cmd.String("f", "", "Download file by -f <xy>")
	targetPath := cmd.String("target", "", "Download file path by -target <xy>")
	// getting a handler
	cliHandler := NewCLIHander(cmd, "Use remote and f together")

	if *remoteName != "" && *downloadFile != "" {
		d := ""
		remotes := cliHandler.GetRemotes(remoteName, &d)
		downloadFunc(remotes, *downloadFile, *targetPath)
	} else {
		cliHandler.DefaultHelp()
	}
}

func downloadFunc(remotes []remote.Remote, downloadFile, targetPath string) {
	i, t := 0, len(remotes)
	type downloadChan struct {
		r      remote.Remote
		target string
		file   entity.File
		err    error
	}
	result := make(chan downloadChan, t)
	for _, r := range remotes {
		go func(r remote.Remote, downloadFile, targetPath string, result chan<- downloadChan) {
			rr, ff, ee := downloadNow(r, downloadFile, targetPath)
			result <- downloadChan{r: rr, target: targetPath, file: ff, err: ee}
		}(r, downloadFile, targetPath, result)
	}
	for ch := range result {
		downloadPrint(ch.r, ch.target, ch.file, ch.err)
		i++
		if i == t {
			close(result)
		}
	}
	fmt.Println()
	fmt.Println()
}

func downloadNow(r remote.Remote, downloadFile, targetPath string) (remote.Remote, entity.File, error) {
	pm, err := master.NewPatchMaster(r.Name(), false)
	var file entity.File
	if err == nil {
		r.UpdateStatus(true)
		f, e := pm.DownloadFileFromRemote(downloadFile, targetPath)
		err = e
		file = f
	}
	return r, file, err
}

func downloadPrint(r remote.Remote, target string, file entity.File, err error) {
	LoaderSkip()
	cwd, _ := os.Getwd()
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	format := "%v\t%v\t\t\t%v\t\n"
	fmt.Fprintf(tw, format, "Remote Name", fmt.Sprintf("%s [%s]", r.Name(), r.Type()), iif(r.Status(), greenText("...OK"), redText("...NOT REACHABLE")))
	if err == nil {
		fmt.Fprintf(tw, format, "Downloaded", fmt.Sprintf("Name - %s", file.Name()), greenText("DOWNLOADED"))
		fmt.Fprintf(tw, format, "", fmt.Sprintf("Size - %d", file.Size()), "")
		fmt.Fprintf(tw, format, "", fmt.Sprintf("Path - %s", filepath.Join(iif(target != "", target, cwd).(string), file.Name())), "")
	} else {
		fmt.Fprintf(tw, format, "Downloaded", "", redText("FAILED"))
		fmt.Fprintf(tw, format, "", fmt.Sprintf("%s", err.Error()), "")
	}
	tw.Flush()
}
