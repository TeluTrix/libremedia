/*
Copyright © 2024 Luca Albrecht <luca@albright.one>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/telutrix/libremedia/cmd/internal/command"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Initial installation",
	Long:  `Initially used to set up and install libremedia.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Implement a variety of initial setup steps
		command.SetupCmd()
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
