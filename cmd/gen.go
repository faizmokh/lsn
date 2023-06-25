/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/faizmokh/lsn/gen"
	"github.com/spf13/cobra"
)

var (
	author string
	year   int16
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
	RunE: func(cmd *cobra.Command, args []string) error {
		data := gen.LicenseData{Author: author, Year: year}
		license, err := gen.GetLicense(args[0], &data)
		if err != nil {
			return err
		}

		fmt.Println(license)
		return nil
	},
}

func init() {
	genCmd.PersistentFlags().StringVarP(&author, "author", "a", "", "Authors's name")

	defaultYear := int16(time.Now().Year())
	genCmd.PersistentFlags().Int16VarP(&year, "year", "y", defaultYear, "License's year")

	rootCmd.AddCommand(genCmd)
}
