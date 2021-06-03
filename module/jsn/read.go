package jsn

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/pnkj-kmr/patch/utility"
)

// RemoteApp defines the app basic details
type RemoteApp struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Service string `json:"service"`
	Type    string `json:"apptype"`
	Status  bool   `json:"status"`
}

// Remote defines the server basic details
type Remote struct {
	Name    string      `json:"name"`
	Address string      `json:"address"`
	Apps    []RemoteApp `json:"apps"`
	Status  bool        `json:"status"`
}

// GetRemotes - load func for remotes r
func GetRemotes() (r []Remote, err error) {
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	// Open our jsonFile
	path := filepath.Join(wd, utility.ConfDirectory, "remotes.json")
	jsonFile, err := os.Open(path)
	if err != nil {
		return
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}

	json.Unmarshal(byteValue, &r)
	log.Println("DEFAULT REMOTE CONFIGURATION FILE", path)
	return
}
