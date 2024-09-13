package internal

import (
	"runtime"
	"strings"
)

func FigureOutOS() byte {
	switch {
	// in case of a windows machine
	case strings.Contains(runtime.GOOS, "windows"):
		return 'w'
	// in case of a mac machine
	case strings.Contains(runtime.GOOS, "darwin"):
		return 'm'
		// in case of a unix-based system (mac aside)
	case strings.Contains(runtime.GOOS, "linux") || strings.Contains(runtime.GOOS, "openbsd") || strings.Contains(runtime.GOOS, "freebsd") || strings.Contains(runtime.GOOS, "netbsd"):
		return 'l'
		// in case of an unsupported os
	default:
		return '0'
	}
}
