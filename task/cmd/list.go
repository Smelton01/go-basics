/*
Copyright © 2021 Simon Juba scimail09@gmail.com

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

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your incomplete tasks",
	RunE: myConn.list,
}

func (c *Conn) list(cmd *cobra.Command, args []string) error {
	
	return c.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todos"))
		c := b.Cursor()

		counter := 1
		fmt.Println("You have the following tasks:")

		// Iterate through your database entries and print out pending tasks
		for k, v := c.First(); k != nil; k, v = c.Next() {
			vals := strings.Split(string(v), separator)
			if vals[0] == "true"{
				continue
			}
			fmt.Printf("%v. %s\n", counter, k)
			counter++
		}
		return nil
	})}

func init() {
	rootCmd.AddCommand(listCmd)
}
