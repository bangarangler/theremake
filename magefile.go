// +build mage

package main

import (
	"fmt"
	"os"

	// "os/exec"

	// "github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

// A build step that requires additional params, or platform specific steps for example
// func Build() error {
// 	mg.Deps(InstallDeps)
// 	fmt.Println("Building...")
// 	cmd := exec.Command("go", "build", "-o", "MyApp", ".")
// 	return cmd.Run()
// }

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
