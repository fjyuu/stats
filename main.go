package main

import (
	"bufio"
	"fmt"
	"io"
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
		return Run(os.Stdin, os.Stdout, os.Stderr)
	},
}

// Run reads numbers from stdin and output statistics.
func Run(stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	calculator := lib.NewCalculator()
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			fmt.Fprintln(stderr, "[Warning] skip empty line")
			continue
		}
		value, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Fprintf(stderr, "[Warning] failed to parse '%s'\n", text)
			continue
		}
		if err := calculator.Input(value); err != nil {
			return fmt.Errorf("failed to input value '%s'", text)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read standard input: %s", err.Error())
	}

	result, err := calculator.GetResult()
	if err != nil {
		return err
	}
	printResult(stdout, result)

	return nil
}

func printResult(stdout io.Writer, result *lib.Result) {
	fmt.Fprintf(stdout, "count\t%d\n", result.Count)
	fmt.Fprintf(stdout, "mean\t%f\n", result.Mean)
	fmt.Fprintf(stdout, "std\t%f\n", result.Std)
	fmt.Fprintf(stdout, "min\t%f\n", result.Min)
	fmt.Fprintf(stdout, "max\t%f\n", result.Max)
	fmt.Fprintf(stdout, "sum\t%f\n", result.Sum)
}

func main() {
	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "Show version")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
