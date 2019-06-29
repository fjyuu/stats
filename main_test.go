package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("skip empty line", func(t *testing.T) {
		stdin := strings.NewReader("1\n\n2\n")
		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}

		if err := Run(stdin, stdout, stderr); err != nil {
			t.Fatal(err)
		}
		if !resultHasCount(stdout, 2) {
			t.Fatalf("count is wrong")
		}
	})

	t.Run("skip invalid value", func(t *testing.T) {
		stdin := strings.NewReader("1\nA\n2\n")
		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}

		if err := Run(stdin, stdout, stderr); err != nil {
			t.Fatal(err)
		}
		if !resultHasCount(stdout, 2) {
			t.Fatalf("count is wrong")
		}
	})

	t.Run("trim spaces", func(t *testing.T) {
		stdin := strings.NewReader("  1  \n")
		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}

		if err := Run(stdin, stdout, stderr); err != nil {
			t.Fatal(err)
		}
		if !resultHasCount(stdout, 1) {
			t.Fatalf("count is wrong")
		}
	})

	t.Run("failed when read empty string", func(t *testing.T) {
		stdin := strings.NewReader("")
		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}

		if err := Run(stdin, stdout, stderr); err == nil {
			t.Fatal("succeeded when read empty string")
		}
		if stdout.Len() != 0 {
			t.Fatal("output result")
		}
	})
}

func resultHasCount(stdout io.Reader, count uint) bool {
	output, _ := ioutil.ReadAll(stdout)
	return strings.Contains(string(output), fmt.Sprintf("count\t%d", count))
}
