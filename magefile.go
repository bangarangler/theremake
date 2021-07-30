// +build mage

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

// A build step that requires additional params, or platform specific steps for example
func BuildClient() error {
	fmt.Println("running npm run build for client...")
	os.Chdir("./remake-client")
	defer os.Chdir("..")
	err := sh.Run("npm", "run", "build")
	return err
}

func InstallClient() error {
	fmt.Println("running npm install for client...")
	os.Chdir("./remake-client")
	defer os.Chdir("..")
	err := sh.Run("npm", "install")
	return err
}

func StartProd() error {
	mg.Deps(CheckLanguageVersion)
	mg.Deps(InstallClient)
	mg.Deps(BuildClient)
	mg.Deps(PostBuild)
	fmt.Println("Starting Prod...")

	err := sh.Run("caddy", "run", "--config", "./remake-client/Caddyfile")

	return err
}

func Start() {
	fmt.Println("Starting Prod for jonathandain.dev...")
	mg.Deps(StopProd)
	mg.Deps(Clean)
	mg.Deps(StartProd)
}

func PostBuild() error {
	var err error
	fmt.Println("moving robots.txt into build dir...")
	err = sh.Run("cp", "-r", "./remake-client/robots.txt", "./remake-client/build/")
	return err
}

func StopProd() error {
	var err error
	fmt.Println("Stopping Caddy...")
	err = sh.Run("caddy", "stop")
	return err
}

func Clean() error {
	var err error
	fmt.Println("Cleaning...")
	fmt.Println("deleteing all node_modules and build dir from client...")
	os.Chdir("./remake-client")
	defer os.Chdir("..")
	err = sh.Run("rm", "-rf", "node_modules")
	err = sh.Run("rm", "-rf", "build")
	return err
}

func CheckLanguageVersion() error {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	sh.Run("node", "-v") // print node version to console
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	input := string(out[1:])                // remove the v
	input = strings.TrimSuffix(input, "\n") // remove the \n so the bytes are a match further down

	b, err := ioutil.ReadFile(".tool-versions")
	if err != nil {
		fmt.Println("err reading .tool-versions", err)
		return err
	}
	str := string(b)
	s := strings.Fields(str)
	fmt.Println("Checking Node Version for match with .tool-versions", s[1]+"=="+input+"?")

	if s[1] != input {
		fmt.Println("You need to change your language version before continuing...")
		return errors.New("Cancel!")
	}
	return err
}

// A custom install step if you need your bin someplace other than go/bin
// func Install() error {
// 	mg.Deps(Build)
// 	fmt.Println("Installing...")
// 	return os.Rename("./MyApp", "/usr/bin/MyApp")
// }

// Install packages for client
func InstallDeps() error {
	fmt.Println("Installing Deps for client...")
	os.Chdir("./remake-client")
	defer os.Chdir("..")
	err := sh.Run("npm", "install")
	return err
}

// Install packages for backend
func InstallBEDeps() error {
	fmt.Println("Installing Deps for Backend...")
	os.Chdir("./remake-backend")
	defer os.Chdir("..")
	err := sh.Run("go", "mod", "tidy")
	return err
}

// Start Client Dev Server
func FEDev() error {
	fmt.Println("Starting up client development server!")
	os.Chdir("./remake-client")
	defer os.Chdir("..")
	err := sh.Run("npm", "run", "dev")
	return err
}

// Start Backend Dev Server
func BEDev() error {
	fmt.Println("Starting up backend development server!")
	os.Chdir("./remake-backend/cmd/")
	defer os.Chdir("..")
	err := sh.Run("air")
	return err
}

// Clean up after yourself
// func Clean() {
// 	fmt.Println("Cleaning...")
// 	os.RemoveAll("MyApp")
// }
