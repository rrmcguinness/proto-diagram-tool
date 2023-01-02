/*
 * Copyright 2022 Google, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"encoding/json"
	"flag"
	"github.com/rrmcguinness/proto-diagram-tool/pkg/proto"
	"os"
	"path/filepath"
	"strings"
)

var directory *string
var recursive *bool
var debug *bool
var output *string

const (
	ProtoSuffix = ".proto"
)

func init() {
	directory = flag.String("d", ".", "The directory to read.")
	recursive = flag.Bool("r", true, "Read recursively.")
	debug = flag.Bool("debug", false, "Enable debugging")
	output = flag.String("o", ".", "Specifies the output directory, if not specified, the processor will write markdown in the proto directories.")
}

func main() {
	flag.Parse()

	proto.SetDebug(*debug)
	logger := proto.Log
	logger.Infof("Reading Directory : %s\n", *directory)
	logger.Infof("Recursively: %v\n", *recursive)

	packages := make([]*proto.Package, 0)

	err := filepath.Walk(*directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ProtoSuffix) {

			pkg := proto.NewPackage(path)
			err := pkg.Read(*debug)
			if err != nil {
				logger.Errorf("error while reading package %s with value: %v", path, err)
			}
			packages = append(packages, pkg)
		}

		return nil
	})

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	for _, pkg := range packages {
		dir := filepath.Dir(pkg.Path)
		bName := filepath.Base(pkg.Path)
		out := dir + string(filepath.Separator) + bName + ".md"
		err = os.WriteFile(out, []byte(pkg.ToMarkdownWithDiagram()), 0644)
		if err != nil {
			logger.Errorf("failed to write file %v\n", err)
		}
	}

	if *debug {
		for _, pkg := range packages {
			err := enc.Encode(pkg)
			if err != nil {
				logger.Errorf("Error encoding package %v", err)
			}
		}

		if err != nil {
			logger.Errorf("failed to process directory: %s with error: %v", *directory, err)
		}
	}
}
