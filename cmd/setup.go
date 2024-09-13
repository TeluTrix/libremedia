/*
Copyright © 2024 Luca Albrecht <luca@albright.one>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Initial installation",
	Long:  `Initially used to set up and install libremedia.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Implement a variety of initial setup steps
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
