//go:build windows
// +build windows

package main

import (
	"traversal-share/internal/traversal"
)

func main() {
	traversal.HolePunch()
	traversal.HoleConnect()
}
