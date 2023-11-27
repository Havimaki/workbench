/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var (
	urlPath      string
	pingSchedule = time.Second * 2
	client       = http.Client{
		Timeout: pingSchedule,
	}
)

func ping(domain string) (int, error) {
	url := "http://" + domain
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	resp.Body.Close()
	return resp.StatusCode, nil
}

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "This pings a remote url",
	Run: func(cmd *cobra.Command, args []string) {

		// ====== ping command
		if resp, err := ping(urlPath); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {
	// Add flags
	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "The url to ping")

	// Assert required flags
	if err := pingCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)
	}

	// Add commands to Parent NetCmd
	NetCmd.AddCommand(pingCmd)
}
