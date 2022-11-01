/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/scmtble/gomodctl/core"
	"github.com/scmtble/gomodctl/parse"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "Gomodctl",
		Short: "Gomodctl is a cmd tool to manage go.mod",
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}

	versionCmd = &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Gomodctl version: v1.0.0")
		},
		Short: "Show version",
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// PersistentFlags全局flag
	rootCmd.PersistentFlags().StringVarP(&parse.FilePath, "path", "p", ".", "Set Go Mod Directory")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(core.NewLsCmd())
}
