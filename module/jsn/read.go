package jsn

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// RemoteApp defines the app basic details
type RemoteApp struct {
	AppType string `json:"apptype"`
	Name    string `json:"name"`
	Path    string `json:"path"`
	Service string `json:"service"`
}

// Remote defines the server basic details
type Remote struct {
	Name    string      `json:"name"`
	Address string      `json:"address"`
	Apps    []RemoteApp `json:"apps"`
}

// GetRemotes - load func for remotes r
func GetRemotes() (r []Remote, err error) {
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	// Open our jsonFile
	jsonFile, err := os.Open(filepath.Join(wd, "conf", "remotes.json"))
	if err != nil {
		return
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}

	json.Unmarshal(byteValue, &r)
	return
}
