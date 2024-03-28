//go:build windows
// +build windows

package main

import (
	"traversal-share/internal/ipv4"
)

func main() {
	ipv4.HolePunch()
}
