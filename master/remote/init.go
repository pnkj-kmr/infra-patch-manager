package remote

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pnkj-kmr/infra-patch-manager/entity"
)

var _remotes []_remote

func init() {
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	// Open our jsonFile
	path := filepath.Join(wd, entity.ConfPath, "remotes.json")
	jsonFile, err := os.Open(path)
	if err != nil {
		return
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}
	// loading remotes from json
	json.Unmarshal(byteValue, &_remotes)
	// fmt.Println("Remote Conf	: conf/remotes.json")
	return
}

// GetRemotes - returns all available remotes
func GetRemotes() (remotes []Remote) {
	for _, r := range _remotes {
		rr, _ := NewRemote(r.RemoteName)
		remotes = append(remotes, rr)
	}
	return
}

// GetRemotesByType - returns available remotes by its type
func GetRemotesByType(t string) (remotes []Remote) {
	for _, r := range _remotes {
		if t == r.RemoteType {
			rr, _ := NewRemote(r.RemoteName)
			remotes = append(remotes, rr)
		}
	}
	return
}
