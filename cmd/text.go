// Copyright © 2018 Brian Johnson <brian@brianjohnson.co.za>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"encoding/base64"
	"fmt"
	"log"

	"bitbucket.com/brianj-za/hash/hashers"
	"github.com/spf13/cobra"
)

// textCmd represents the text command
var textCmd = &cobra.Command{
	Use:   "text",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		algs, err := cmd.Flags().GetStringArray("algorithms")
		if err != nil {
			panic(err)
		}
		text, err := cmd.Flags().GetString("text")
		if err != nil {
			panic(err)
		}
		encode, err := cmd.Flags().GetBool("encode")
		if err != nil {
			panic(err)
		}

		hashers, err := hashers.Get(algs)
		if err != nil {
			log.Println(err)
		}

		for _, hasher := range hashers {
			hasher.Write([]byte(text))
			if encode {
				str := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
				fmt.Println(str)
			} else {
				fmt.Println(hasher.Sum(nil))
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(textCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	textCmd.PersistentFlags().String("text", "", "Text to hash")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// textCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
