package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/fjyuu/stats/lib"
)

var (
	version  string
	revision string
)

var versionFlag bool

var rootCmd = &cobra.Command{
	Use:   "stats",
	Short: "stats calculates statistics such as count, mean, standard deviation, min, and max.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if versionFlag {
			fmt.Printf("stats %s (%s)\n", version, revision)
			return nil
		}

		calculator := lib.NewCalculator()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := strings.TrimSpace(scanner.Text())
			if text == "" {
				fmt.Fprintln(os.Stderr, "[Warning] skip empty line")
				continue
			}
			value, err := strconv.ParseFloat(text, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "[Warning] failed to parse '%s'\n", text)
				continue
			}
			calculator.Input(value)
		}
		if err := scanner.Err(); err != nil {
			return fmt.Errorf("failed to read standard input: %s", err.Error())
		}

		result, err := calculator.GetResult()
		if err != nil {
			return err
		}
		printResult(result)

		return nil
	},
}

func printResult(result *lib.Result) {
	fmt.Printf("count\t%d\n", result.Count)
	fmt.Printf("mean\t%f\n", result.Mean)
	fmt.Printf("std\t%f\n", result.Std)
	fmt.Printf("min\t%f\n", result.Min)
	fmt.Printf("max\t%f\n", result.Max)
	fmt.Printf("sum\t%f\n", result.Sum)
}

func main() {
	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "Show version")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
