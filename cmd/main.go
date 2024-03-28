//go:build windows
// +build windows

package main

import (
	"traversal-share/internal/ipv4"
	"traversal-share/internal/ipv6"
)

func main() {
	ipv4.HolePunch()
	ipv6.HoleConnect()
}
