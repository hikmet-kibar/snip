package app

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hikmet-kibar/snip/cmd/config"
)

// Lists all files of the .snip (or otherwise specified) directory
// TODO
//	- Have all names in a row instead of below one another?
func List(cfg config.Config) error {
	files, err := ioutil.ReadDir(cfg.Directory)
	if err != nil {
		return err
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
	return nil
}

func Remove(cfg config.Config) error {
	var err error
	return nil
}

// Prints the content of a snip
// TODO
//	- Use subdirectories to have smaller snips? like html>a>...
//	  (would reinvent UltiSnips :(
func Get(cfg config.Config) error {
	var err error

	snipPath := filepath.Join(cfg.Directory, cfg.Snip)
	snipExists := exists(snipPath)
	if !snipExists {
		return err
	}

	// Print snip content
	content, err := ioutil.ReadFile(snipPath)
	if err != nil {
		return err
	}
	fmt.Println(string(content))
	// fmt.Fprint(os.Stdout, ...

	return nil
}

// Creates a new file in the .snip directory and opens it in VIM
// TODO
//	- ONLY create the file. Implement a new function to EDIT snips!!
func New(cfg config.Config) error {
	var err error

	snipPath := filepath.Join(cfg.Directory, cfg.Snip)
	snipExists := exists(snipPath)
	if snipExists {
		return err
	}

	err = runCommand("vim", snipPath)
	if err != nil {
		return err
		// TODO: Should the created directories be deleted?
	}

	fmt.Println(snipPath)

	return nil
}

func runCommand(name string, args ...string) (string error) {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}
