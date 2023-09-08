/*
Copyright © 2023 github.com/m4schini
*/
package cmd

import (
	"bytes"
	"fmt"
	"github.com/c-bata/go-prompt"
	"io"
	"logs_analyzer/counter"
	"math"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "logs_analyzer",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		raw, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}

		limit, err := cmd.Flags().GetInt("limit")
		if err != nil {
			limit = -1
		}
		if limit < 0 {
			limit = math.MaxInt
		}

		var levelCounter = counter.New()
		var contextCounter = counter.New()
		var metaCounter = counter.NewTwoDCounter()

		lines := bytes.Split(raw, []byte("\n"))
		processedEntries := 0
		for _, line := range lines {
			if processedEntries >= limit {
				break
			}
			if reHeader.Match(line) {
				match := reHeader.FindSubmatch(line)
				print := func(match [][]byte, i int, prefix string) {
					fmt.Printf("%v: %v\n", prefix, string(match[i]))
				}

				print(match, 1, "timestamp")
				print(match, 2, "level")
				levelCounter.Inc(string(match[2]))
				print(match, 3, "context")
				contextCounter.Inc(string(match[3]))
			} else if reContext.Match(line) {
				match := reContext.FindSubmatch(line)
				matches := reMeta.FindAllSubmatch(match[1], -1)
				for _, i := range matches {
					fmt.Println(string(i[1]), "===>", string(i[2]))
					metaCounter.Inc(string(i[1]), string(i[2]))
				}
			} else {
				fmt.Println("LOG MESSAGE:", string(line))
				fmt.Println("---")
				processedEntries++
			}
		}

		for s, i := range levelCounter {
			fmt.Println(s, i)
		}
		fmt.Println("---")
		for s, i := range contextCounter {
			fmt.Println(i, s)
		}
		fmt.Println("---")
		for s, m := range metaCounter {
			fmt.Println(s)
			for s2, i := range m {
				fmt.Println("   ", i, s2)
			}
		}

		fmt.Println("Please select table.")
		t := prompt.Input("> ", completer)
		fmt.Println("You selected " + t)

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.logs_analyzer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().IntP("limit", "l", -1, "max log entries limit (-1 means no limit)")
}
