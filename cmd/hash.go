package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"davFileBrowser/users"
)

func init() {
	rootCmd.AddCommand(hashCmd)
}

var hashCmd = &cobra.Command{
	Use:   "hash <password>",
	Short: "Hashes a password",
	Long:  `Hashes a password using bcrypt algorithm.`,
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		pwd, err := users.HashPwd(args[0])
		checkErr(err)
		fmt.Println(pwd)
	},
}
