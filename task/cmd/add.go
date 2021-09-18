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
	"time"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task details]",
	Short: "Add a new task to your TODO list",
	RunE: myConn.add,
}

func (c *Conn)add(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("task details not provided")
	}
	defer fmt.Printf("Added task \"%s\" to your todo list\n", strings.Join(args, " "))
	return c.db.Update(func(tx *bolt.Tx) error {
        // Retrieve the todos bucket.
        b := tx.Bucket([]byte("todos"))
		
		status := fmt.Sprintf("%s|%s", "false", time.Now().Format(time.RFC1123))
        // Persist bytes to users bucket.
        return b.Put([]byte(strings.Join(args, " ")), []byte(status))
    })
}


func init() {
	rootCmd.AddCommand(addCmd)
}
