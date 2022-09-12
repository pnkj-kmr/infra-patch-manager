package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/pnkj-kmr/infra-patch-manager/master"
	"github.com/pnkj-kmr/infra-patch-manager/master/remote"
)

// HandleRemoteCmd - handler for remote subcmd
func HandleRemoteCmd(cmd *flag.FlagSet) {
	passcode := cmd.String("p", "", "Passcode for remote agent")
	remoteName := cmd.String("remote", "", "Remote by it's name")
	remoteType := cmd.String("remote-type", "", "Remote by it's type")
	remoteAll := cmd.Bool("remote-all", false, "All available remotes")
	execCmd := cmd.String("c", "", "Executable command statement")
	appName := cmd.String("app", "", "Application by it's name")
	appType := cmd.String("app-type", "", "Application by it's type")
	appAll := cmd.Bool("app-all", false, "All available remote applications")
	start := cmd.Bool("start", false, "Start the requested applications")
	stop := cmd.Bool("stop", false, "Stop the requested applications")
	restart := cmd.Bool("restart", false, "Restart the requested applications")
	status := cmd.Bool("status", false, "Check status of requested applications")
	portCheck := cmd.Bool("port", false, "Check the netstat port status the requested applications")

	// getting a handler
	cliHandler := NewCLIHander(cmd, "Always use passcode [i.e. -p ABC] with others combination")

	if *passcode == "" {
		// setting up the message to secure the passcode use case
		cliHandler.DefaultHelp()
		os.Exit(0)
	}

	if *remoteAll || *remoteType != "" || *remoteName != "" {
		remotes := cliHandler.GetRemotes(remoteName, remoteType)
		fmt.Println()
		if *execCmd != "" {
			executeFunc(remotes, execCmd, passcode)
		} else if (*appAll || *appType != "" || *appName != "") && (*start || *stop || *restart || *portCheck || *status) {
			executeFunc2(cliHandler, remotes, passcode, appName, appType, *start, *stop, *status, *restart, *portCheck)
		} else {
			cliHandler.DefaultHelp()
		}
	} else {
		cliHandler.DefaultHelp()
	}
}

func executeFunc(remotes []remote.Remote, execCmd, passcode *string) {
	i, t := 0, len(remotes)
	type execChan struct {
		r   remote.Remote
		out map[string][]byte
		err error
	}
	result := make(chan execChan, t)
	for _, r := range remotes {
		go func(r remote.Remote, result chan<- execChan) {
			out, err := executeCmd(r, map[string]string{"cmd": *execCmd}, passcode)
			result <- execChan{r, out, err}
		}(r, result)
	}
	for ch := range result {
		executePrint(ch.r, ch.out, ch.err)
		i++
		if i == t {
			close(result)
		}
	}
	fmt.Println()
}

func executeFunc2(cli CLI, remotes []remote.Remote, passcode, appName, appType *string, start, stop, status, restart, portCheck bool) {
	i, t := 0, len(remotes)
	type execChan struct {
		r   remote.Remote
		out map[string][]byte
		err error
	}
	result := make(chan execChan, t)
	for _, r := range remotes {
		go func(r remote.Remote, result chan<- execChan) {
			cmds := make(map[string]string)
			apps := cli.GetRemoteApps(r, appName, appType)
			if portCheck {
				for _, a := range apps {
					// TODO - OS dependent - added for linux
					cmds[a.Name()] = fmt.Sprintf("netstat -aenop | grep :%s", a.AppPort())
				}
			} else if stop {
				for _, a := range apps {
					// TODO - OS dependent - added for linux
					cmds[a.Name()] = fmt.Sprintf("service %s stop", a.ServiceName())
				}
			} else if restart {
				for _, a := range apps {
					// TODO - OS dependent - added for linux
					cmds[a.Name()] = fmt.Sprintf("service %s restart", a.ServiceName())
				}
			} else if start {
				for _, a := range apps {
					// TODO - OS dependent - added for linux
					cmds[a.Name()] = fmt.Sprintf("service %s start", a.ServiceName())
				}
			} else if status {
				for _, a := range apps {
					// TODO - OS dependent - added for linux
					cmds[a.Name()] = fmt.Sprintf("service %s status", a.ServiceName())
				}
			}
			out, err := executeCmd(r, cmds, passcode)
			result <- execChan{r, out, err}
		}(r, result)
	}
	for ch := range result {
		executePrint(ch.r, ch.out, ch.err)
		i++
		if i == t {
			close(result)
		}
	}
	fmt.Println()
}

func executeCmd(r remote.Remote, cmds map[string]string, passcode *string) (out map[string][]byte, err error) {
	out = make(map[string][]byte)
	pm, err := master.NewPatchMaster(r.Name(), false)
	pcode := *passcode
	if err == nil {
		for i, s := range cmds {
			v, e := pm.ExecuteCmdOnRemote(s, pcode)
			out[fmt.Sprintf("%s	[passcode: %s] -> %s", i, pcode, s)] = v
			if e != nil {
				err = e
				r.UpdateStatus(false)
			} else {
				r.UpdateStatus(true)
			}
		}
	}
	return
}

func executePrint(r remote.Remote, out map[string][]byte, err error) {
	LoaderSkip()
	var e string
	if err != nil {
		e = err.Error()
	}
	status := iif(e != "", redText("...INVALID"), iif(r.Status(), greenText("...OK"), redText("...NOT REACHABLE")))
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	format := "%v\t%v\t\t\t%v\t\n"
	fmt.Fprintf(tw, format, "Remote Name:", fmt.Sprintf("%s [%s]", r.Name(), r.Type()), status)
	for s, o := range out {
		fmt.Fprintf(tw, "%v\n", s)
		fmt.Fprintf(tw, "%v\n", strings.Repeat("-", 60))
		if len(o) > 0 {
			fmt.Fprintf(tw, "%v", yellowText(string(o)))
		} else if e != "" {
			fmt.Fprintf(tw, "%v\n", redText(e))
		} else {
			fmt.Fprintf(tw, "%v\n", redText("NO OUTPUT RECEIVED"))
		}
		fmt.Fprintf(tw, "%v\n\n", strings.Repeat("-", 60))
	}
	tw.Flush()
}
