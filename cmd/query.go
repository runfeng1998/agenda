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
	"fmt"
    "io/ioutil"
    "agenda/entity"
	"github.com/spf13/cobra"
)

const userpath="user.txt"
const statepath="curUser.txt"

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "query",
    Long: `query user information, should login, example:
agenda query
    `,
	Run: func(cmd *cobra.Command, args []string) {
        state,_:=ioutil.ReadFile(statepath);
        if string(state)=="none" {
            fmt.Println("no login")
            return
        }
        fmt.Println("user information table:") 
        users,_:=entity.ReadUser(userpath)
        for _,user:= range users {
            fmt.Println(user.Username, user.Email, user.Telephone)
        }

		fmt.Println("query success")
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
