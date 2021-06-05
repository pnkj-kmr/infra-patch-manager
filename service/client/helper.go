package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/pnkj-kmr/infra-patch-manager/service"
	"github.com/pnkj-kmr/infra-patch-manager/utility"
)

// getRemotes - load func for remotes r
// func loads the remote configration from config file
func getRemotes() (r []service.Remote, err error) {
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
