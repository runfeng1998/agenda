/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"agenda/entity"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "log in",
	Long: `
		log in, example:
login -u [username] -p [password]
agenda login -u ryanwang -p 123456
	`,
	Run: func(cmd *cobra.Command, args []string) {
		users, _ := entity.ReadUser(userpath)

		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		for _, ele := range users {
			if ele.Username == username {
				if ele.Password != password {
					fmt.Println("username or password not correct")
				} else if ele.Password == password {
					ioutil.WriteFile(statepath, []byte(username), os.ModeAppend)
					fmt.Println("login successful")
				}
				return
			}
		}

		fmt.Println("username or password not correct")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	loginCmd.Flags().StringP("username", "u", "", "-u your username")
	loginCmd.Flags().StringP("password", "p", "", "-p your password")
}
