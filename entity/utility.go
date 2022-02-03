package entity

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// EnableAgentMode - basic setup for agent to run
func EnableAgentMode() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
	// Creating the default folders for applications
	for _, d := range []string{resourceDir, assetsDir, patchDir, rollbackDir} {
		_, err = CreateDirectoryIfNotExists(d)
		if err != nil {
			log.Fatal(filepath.Join(wd, d), "ERROR:", err)
		}
	}
}

// CreateDirectoryIfNotExists helps to create a dir if not exists and returns the dir path
func CreateDirectoryIfNotExists(path string) (string, error) {
	// creating the directories if not exists
	cd, err := NewDir(path)
	if err != nil {
		if os.IsNotExist(err) {
			d, err := NewDir(filepath.Dir(path))
			if err != nil {
				log.Println("Cannot load directory information: ", path, "--", filepath.Dir(path), err)
				return "", err
			}
			p := strings.Replace(path, filepath.Dir(path), "", 1)
			log.Println("Directory: ", d.Path(), p)
			err = d.Create(p)
			if err != nil {
				log.Println("Cannot create directory: ", d.Path(), p, err)
				return "", err
			}
			newpath := filepath.Join(d.Path(), p)
			log.Println("CREATE: if not exists - dir -", newpath)
			return newpath, err
		}
		return "", err
	}
	return cd.Path(), nil
}

// RandomString helps to generate random charactor with n length
func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}

// RandomStringWithTime helps to generate random charactor with n length with time suffix
func RandomStringWithTime(n int, prefix string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return prefix + string(b) + "__" + strconv.FormatInt(time.Now().Unix(), 10)
}

// Ping helps to validate the ping-pong echo
// input as "PING" output will be "PONG" else ""
func Ping(in string) (out string) {
	if ok := strings.EqualFold(in, "PING"); ok {
		out = "PONG"
	}
	return
}
