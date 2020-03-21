// +build mage

package main

import (
	"fmt"
	// mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/mg"
	// sh contains helpers for running shell-like commands
	"github.com/magefile/mage/sh"
)

// set up environment variables
var env = map[string]string{
	"MAGEFILE_VERBOSE": "1",
	"GO111MODULE":      "on",
}

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = Test

// Dep target fetches project dependencies
func Dep() error {
	fmt.Println("Installing Deps...")
	return sh.RunWith(env, "go", "mod", "tidy")
}

// Vet target makes vetting
func Vet() error {
	fmt.Println("Vetting...")
	return sh.RunWith(env, "go", "vet", "./...")
}

// Fmt target makes formatting
func Fmt() error {
	fmt.Println("Formatting...")
	return sh.RunWith(env, "go", "fmt", "./...")
}

// Clean up
func Clean() error {
	fmt.Println("Cleaning...")
	return sh.RunWith(env, "go", "clean")
}

// Test target executes project tests
func Test() error {
	mg.SerialDeps(Clean, Fmt, Vet, Dep)
	fmt.Println("Testing...")
	return sh.RunWith(env, "go", "test", "-v", "-cover", "./...")
}
