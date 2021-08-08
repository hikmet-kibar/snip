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
