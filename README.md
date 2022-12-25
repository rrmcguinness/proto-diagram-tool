# Proto Diagram Utility

This utility package is a compiled Go program that reads a protobuf
source directory and generates Mermaid Diagrams in PROTO_DIAGRAM.md files
in each directory.

> NOTE: Only Proto 3 syntax is supported.

This utility was created to ease documentation generation of complex
Protobuf Libraries in order to simplify understanding the models and services
described in a Protobuf.

If you find this useful, awesome! If you find a bug, please contribute a patch,
or open a bug. Please follow the [Contributing](CONTRIBUTING.md) guidelines.

Use:

```shell
// Clone
git clone https://github.com/rrmcguinness/proto-diagram-tool

cd proto-diagram-tool

// Build
go build

./proto-diagram -d /path/to/directory -r true
```

## Options
* -d - The directory to scan, this is required.
* -r - Recursively scan all subdirectories (default is true)