package main

import (
	"fmt"
	"log"
	"strings"

	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   appName,
	Short: "Turn an application skeleton into your boilerplate code, using template params, or files",
	Long: `Generate application skeleton/scafold using supplied parameters (command line or YAML/JSON file).
Parameters are saved (add them to source control) so that the skeleton can be updated/regenerated when changed.
Useful for microservice architectures/SOA, enforcing common functions, and in an environemnt
where new common functionality is regularly added.`,
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// checkFiles examines a list of files to check that they all exist
func checkFiles(toCheck []string) error {
	issues := []string{}
	for _, f := range toCheck {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			issues = append(issues, f)
		}
	}
	if len(issues) != 0 {
		return fmt.Errorf("Failed to find files: %s", strings.Join(issues, ","))
	}
	return nil
}

var quiet bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "quiet output mode")
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
