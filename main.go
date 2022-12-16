package main

import (
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

const (
	ProtoSuffix = ".proto"
)

func init() {
	directory = flag.String("d", ".", "The directory to read.")
	recursive = flag.Bool("r", true, "Read recursively.")
}

var mdTemplate = `
# %s

## Comments
%s

` + "```plantuml\n@startuml\n%s\n```\n@enduml\n"

func main() {
	flag.Parse()
	log.Printf("Reading Directory : %s\n", *directory)
	log.Printf("Recursively: %v\n", *recursive)

	tree := make([]*proto.Package, 0)

	err := filepath.Walk(*directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ProtoSuffix) {
			reader := proto.NewReadState()
			pkg := reader.Read(path)
			tree = append(tree, pkg)

			content := fmt.Sprintf(mdTemplate, pkg.Name, pkg.Comments, pkg.PlantUML())
			d := path[0:strings.LastIndex(path, string(filepath.Separator))+1] + pkg.Name + ".md"
			err := os.WriteFile(d, []byte(content), 0644)

			if err != nil {
				fmt.Printf("Failed to write file: %s", d)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("failed to process directory: %s with error: %v", *directory, err)
	}

	for _, p := range tree {
		fmt.Printf("%s\n\n", p.PlantUML())
	}
}
