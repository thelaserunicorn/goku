# goku

CLI tool to convert between JSON and YAML.

## Requirements

- Go 1.25+

## Build

```bash
go build -o goku ./
```

## Usage

```bash
./goku -i <input-file> -o <json|yaml>
```

### Examples

Convert JSON to YAML:

```bash
./goku -i testdata/config.json -o yaml
```

Convert YAML to JSON:

```bash
./goku -i testdata/config.yaml -o json
```

### Output behavior

- The output file is written in the same folder as the input.
- The output file uses the same base name with the new extension.
- The converted content is printed to stdout.

Example: testdata/config.json -> testdata/config.yaml

### Notes

- The input and output formats cannot be the same.
- Supported extensions: .json, .yaml, .yml.
