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
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
	bolt "go.etcd.io/bbolt"
)

var cfgFile string

const (
	path = "todos.db"
	separator = "|"
	)

type Todo struct {
	Index uint64 `json:"index"`
	Details string `json:"details"`
	Status string `json:"status"`
}
var Todos []*Todo

type Conn struct{
	db *bolt.DB
}
var myConn Conn

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task [command] {arguments}",
	Short: "A command line tool to manage tasks",
	Long: `A command line tool developed with the cobra framework and written in Go
	Use it to keep track of todo tasks. 
	
	You can add tasks
	View pending/completed tasks
	Mark tasks as done
	Delete tasks`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Initialize database connection
	var err error
	myConn.db, err = bolt.Open(path, 0666, nil)
	if err != nil {
	  log.Fatalf("cannot connect to database %v: %v", path, err)
	}

	err = myConn.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("todos"))
		if err != nil {
			return fmt.Errorf("cannot create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		log.Fatal("update transaction failed: ", err)
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.task.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".task" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".task")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
