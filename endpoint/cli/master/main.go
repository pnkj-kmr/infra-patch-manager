package main

import (
	"flag"
	"os"

	"github.com/pnkj-kmr/infra-patch-manager/endpoint/cli"
)

func main() {
	// Helps to detail out available remotes details
	remoteCmd := flag.NewFlagSet("remote", flag.ExitOnError)
	rightsCmd := flag.NewFlagSet("rights", flag.ExitOnError)
	uploadCmd := flag.NewFlagSet("upload", flag.ExitOnError)
	extractCmd := flag.NewFlagSet("extract", flag.ExitOnError)
	applyCmd := flag.NewFlagSet("apply", flag.ExitOnError)
	verifyCmd := flag.NewFlagSet("verify", flag.ExitOnError)
	execCmd := flag.NewFlagSet("exec", flag.ExitOnError)

	cli.DefaultCheck()
	switch os.Args[1] {
	case "remote":
		cli.HandleRemote(remoteCmd)
	case "rights":
		cli.HandleRights(rightsCmd)
	case "upload":
		cli.HandleUpload(uploadCmd)
	case "extract":
		cli.HandleExtract(extractCmd)
	case "apply":
		cli.HandleApply(applyCmd)
	case "verify":
		cli.HandleVerify(verifyCmd)
	case "exec":
		cli.HandleRemoteCmd(execCmd)
	default:
		cli.DefaultHelp()
	}
}
