package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func main() {
	rootCmd.Execute()
}

func init() {
	rootCmd = &cobra.Command{Use: "health"}
	rootCmd.AddCommand(healthcheckCmd)

	healthcheckCmd.Flags().StringP("name", "n", "", "Taro")
	healthcheckCmd.MarkFlagRequired("name")
}

var healthcheckCmd = &cobra.Command{
	Use:     "healthcheck",
	Short:   "healthcheck command",
	Aliases: []string{"hc"},
	Run:     healthcheck,
}

func healthcheck(cmd *cobra.Command, arg []string) {

	f := cmd.Flag("name")
	fmt.Println("name", f.Value.String())
	fmt.Println("arg", arg)
}
