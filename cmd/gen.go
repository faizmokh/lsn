/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/faizmokh/lsn/gen"
	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate open source license",
	Long: `Generate an open source license by providing the name of the license and the author's name. For example:

	$~ lsn gen apache --author="John Doe" --year="2023"

Will output an Apache open source license with the author name replace as "John Doe" and the year as "2023".
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		license, err := gen.GetLicense(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(license)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
