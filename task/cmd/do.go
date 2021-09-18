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
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do [task number]",
	Short: "Mark a task on your TODO list as complete",
	RunE: myConn.do,
}

func (c *Conn) do(cmd *cobra.Command, args []string) error{
	// mark task as done
	if len(args) != 1 {
		return fmt.Errorf("too many arguments")
	}

	if _, e := strconv.ParseInt(args[0], 10, 32); e != nil {
		return fmt.Errorf("task index must be number")
	}

	return c.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todos"))
		c := b.Cursor()

		index := 1
		for k, v := c.First(); k != nil; k, v = c.Next() {
			// Update task to done if index matches
			if strconv.Itoa(index) == args[0]{
				fmt.Printf("Task \"%v\" marked as done", k)
				val := strings.Split(string(v), separator)
				val[0] = "true"
				return b.Put(k, []byte(strings.Join(val, separator)))}
			index++
		}
		return fmt.Errorf("invalid task index")
	})
}

func init() {
	rootCmd.AddCommand(doCmd)
}
