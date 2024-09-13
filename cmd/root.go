/*
Copyright © 2024 Luca Albrecht luca@albright.one
*/
package cmd

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/telutrix/libremedia/cmd/internal"
)

// variables
var OS byte

var rootCmd = &cobra.Command{
	Use:   "libremedia",
	Short: "An app used to manage your media library over the command line",
	Long:  `The app "libremedia" can be used to manage and watch movies and tv-shows over the command line.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Check operating system (u = unrecognized)
	if internal.FigureOutOS() == '0' {
		slog.Error("libremedia doesn't support any operating systems outside of linux, mac and windows")
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
