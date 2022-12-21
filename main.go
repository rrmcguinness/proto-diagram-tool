package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rrmcguinness/proto-diagram/pkg/proto"
)

var directory *string
var recursive *bool
var debug *bool

const (
	ProtoSuffix = ".proto"
)

func init() {
	directory = flag.String("d", ".", "The directory to read.")
	recursive = flag.Bool("r", true, "Read recursively.")
	debug = flag.Bool("debug", true, "Enable debugging")
}

var mdTemplate = `
# %s

## Comments
%s

` + "```mermaid\n%s\n```\n\n"

func main() {
	flag.Parse()
	log.Printf("Reading Directory : %s\n", *directory)
	log.Printf("Recursively: %v\n", *recursive)

	packages := make([]*proto.Package, 0)

	err := filepath.Walk(*directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ProtoSuffix) {

			pkg := proto.NewPackage(path)
			err := pkg.Read(*debug)
			if err != nil {
				fmt.Printf("error while reading package %s with value: %v", path, err)
			}
			packages = append(packages, pkg)
		}

		return nil
	})

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	for _, pkg := range packages {
		err := enc.Encode(pkg)
		if err != nil {
			fmt.Printf("Error encoding package %v", err)
		}
	}

	if err != nil {
		log.Fatalf("failed to process directory: %s with error: %v", *directory, err)
	}
}
