package cli

import (
	"flag"
	"fmt"
	"strings"

	"github.com/pnkj-kmr/infra-patch-manager/master"
	"github.com/pnkj-kmr/infra-patch-manager/master/remote"
)

// HandleRemoteCmd - handler for remote subcmd
func HandleRemoteCmd(cmd *flag.FlagSet) {
	remoteName := cmd.String("remote", "", "Remote by it's name")
	remoteType := cmd.String("remote-type", "", "Remote by it's type")
	remoteAll := cmd.Bool("remote-all", false, "All available remotes")
	execCmd := cmd.String("c", "", "Executable command statement")
	appName := cmd.String("app", "", "Application by it's name")
	appType := cmd.String("app-type", "", "Application by it's type")
	appAll := cmd.Bool("app-all", false, "All available remote applications")
	start := cmd.Bool("start", false, "Start the requested applications")
	stop := cmd.Bool("stop", false, "Stop the requested applications")
	restart := cmd.Bool("restart", false, "Stop the requested applications")
	portCheck := cmd.Bool("check-port", false, "Check the netstat port status the requested applications")
	// getting a handler
	cliHandler := NewCLIHander(cmd, "")

	if *remoteAll || *remoteType != "" || *remoteName != "" {
		remotes := cliHandler.GetRemotes(remoteName, remoteType)
		if *execCmd != "" {
			fmt.Println()
			for _, r := range remotes {
				printRemoteCmd(r, map[string]string{"cmd": *execCmd})
				fmt.Println()
			}
		} else if (*appAll || *appType != "" || *appName != "") && (*start || *stop || *restart || *portCheck) {
			for _, r := range remotes {
				cmds := make(map[string]string)
				apps := cliHandler.GetRemoteApps(r, appName, appType)
				if *portCheck {
					for _, a := range apps {
						// TODO - OS dependent
						cmds[a.Name()] = fmt.Sprintf("netstat -aneop | grep :%s", a.AppPort())
					}
				} else if *stop {
					for _, a := range apps {
						// TODO - OS dependent
						cmds[a.Name()] = fmt.Sprintf("service stop %s", a.ServiceName())
					}
				} else if *restart {
					for _, a := range apps {
						// TODO - OS dependent
						cmds[a.Name()] = fmt.Sprintf("service restart %s", a.ServiceName())
					}
				} else if *start {
					for _, a := range apps {
						// TODO - OS dependent
						cmds[a.Name()] = fmt.Sprintf("service start %s", a.ServiceName())
					}
				}
				printRemoteCmd(r, cmds)
				fmt.Println()
			}
		} else {
			cliHandler.DefaultHelp()
		}
	} else {
		cliHandler.DefaultHelp()
	}
}

func printRemoteCmd(r remote.Remote, cmds map[string]string) {
	out := executeRemoteCmd(r, cmds)
	fmt.Printf("Remote name	: %s [%s]		%s\n", r.Name(), r.Type(), iif(r.Status(), greenText("--- OK"), redText("--- NOT REACHABLE")))
	for s, o := range out {
		fmt.Printf("Execute		: %s\n", s)
		if len(out) > 0 {
			fmt.Println(strings.Repeat("-", 60))
			fmt.Printf(yellowText(string(o)))
			fmt.Println(strings.Repeat("-", 60))
		}
	}
}

func executeRemoteCmd(r remote.Remote, cmds map[string]string) (out map[string][]byte) {
	out = make(map[string][]byte)
	pm, err := master.NewPatchMaster(r.Name(), false)
	if err == nil {
		for i, s := range cmds {
			v, err := pm.ExecuteCmdOnRemote(s)
			out[i+" -> "+s] = v
			if err != nil {
				r.UpdateStatus(false)
			} else {
				r.UpdateStatus(true)
			}
		}
	}
	return
}
