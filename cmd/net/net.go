/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"github.com/spf13/cobra"
)

var NetCmd = &cobra.Command{
	Use:   "net",
	Short: "Net is a palette that contains network based commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {}
