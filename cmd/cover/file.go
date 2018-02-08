// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements the visitor that computes the (line, column)-(line-column) range for each function.

package main

import (
	"bufio"
	"fmt"
	"os"
	"text/tabwriter"

	"golang.org/x/tools/cover"
)

// funcOutput takes two file names as arguments, a coverage profile to read as input and an output
// file to write ("" means to write to standard output). The function reads the profile and produces
// as output the coverage data broken down by function, like this:
//
//	fmt/format.go:30:	init			100.0%
//	fmt/format.go:57:	clearflags		100.0%
//	...
//	fmt/scan.go:1046:	doScan			100.0%
//	fmt/scan.go:1075:	advance			96.2%
//	fmt/scan.go:1119:	doScanf			96.8%
//	total:		(statements)			91.9%

func fileOutput(profile, outputFile string) error {
	profiles, err := cover.ParseProfiles(profile)
	if err != nil {
		return err
	}

	var out *bufio.Writer
	if outputFile == "" {
		out = bufio.NewWriter(os.Stdout)
	} else {
		fd, err := os.Create(outputFile)
		if err != nil {
			return err
		}
		defer fd.Close()
		out = bufio.NewWriter(fd)
	}
	defer out.Flush()

	tabber := tabwriter.NewWriter(out, 1, 8, 1, '\t', 0)
	defer tabber.Flush()

	for _, profile := range profiles {
		fn := profile.FileName

		fmt.Fprintf(tabber, "%s\t%.1f%%\n", fn, 100.0*float64(c)/float64(t))
		if err != nil {
			return err
		}
	}

	return nil
}

