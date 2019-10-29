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

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register user",
	Long: `register user, example:
register -u [username] -p [password] -e [email] -t [telephone]
agenda register -u ryanwang -p 123456 -e 123456789@qq.com -t 12345678910
	`,
	Run: func(cmd *cobra.Command, args []string) {
		users, _ := entity.ReadUser(userpath)

		//process flag
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		telephone, _ := cmd.Flags().GetString("telephone")

		if username == "" {
			fmt.Println("username can't be empty")
			return
		}
		if password == "" {
			fmt.Println("password can't be empty")
			return
		}
		if email == "" {
			fmt.Println("email can't be empty")
			return
		}
		if telephone == "" {
			fmt.Println("telephone can't be empty")
			return
		}

		user := entity.User{username, password, email, telephone}
		fmt.Println(users)
		fmt.Println(user)

		//check username unique
		for _, ele := range users {
			if ele == user {
				fmt.Println("fail, exist the same username")
				return
			}
		}
		users = append(users, user)
		entity.WriteUser(userpath, users)
		fmt.Println(users)

		fmt.Println("successful register")
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("username", "u", "", "-u your username")
	registerCmd.Flags().StringP("password", "p", "", "-p your password")
	registerCmd.Flags().StringP("email", "e", "", "-e your email")
	registerCmd.Flags().StringP("telephone", "t", "", "-t your telephone")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
