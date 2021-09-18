/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "list out any tasks completed today",
	RunE: myConn.completed,
}

func (c *Conn) completed(cmd *cobra.Command, args []string) error {
	return c.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todos"))
		c := b.Cursor()

		fmt.Println("Congratulations, you finished the following tasks today:")

		// Iterate through your database entries and print out tasks completed in thelast 24 hours
		for k, v := c.First(); k != nil; k, v = c.Next() {
			vals := strings.Split(string(v), separator)
			date, err := time.Parse(time.RFC1123, vals[1])
			if err != nil {
				return fmt.Errorf("error parsing date: %v", err)
			}
			if time.Since(date) < 24*time.Hour && vals[0] == "true" {
				fmt.Printf("- %s\n", k)
			}
		}
		return nil
	})
}

func init() {
	rootCmd.AddCommand(completedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
