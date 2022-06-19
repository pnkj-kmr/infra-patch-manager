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
		s := cli.Loader()
		s.Start()
		cli.HandleRemote(remoteCmd)
		s.Stop()
	case "rights":
		s := cli.Loader()
		s.Start()
		cli.HandleRights(rightsCmd)
		s.Stop()
	case "upload":
		s := cli.Loader()
		s.Start()
		cli.HandleUpload(uploadCmd)
		s.Stop()
	case "extract":
		s := cli.Loader()
		s.Start()
		cli.HandleExtract(extractCmd)
		s.Stop()
	case "apply":
		s := cli.Loader()
		s.Start()
		cli.HandleApply(applyCmd)
		s.Stop()
	case "verify":
		s := cli.Loader()
		s.Start()
		cli.HandleVerify(verifyCmd)
		s.Stop()
	case "exec":
		s := cli.Loader()
		s.Start()
		cli.HandleRemoteCmd(execCmd)
		s.Stop()
	default:
		cli.DefaultHelp()
	}
}
