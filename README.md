# goku

CLI tool for JSON/YAML conversion and PostgreSQL resource management.

## Requirements

- Go 1.25+
- Docker & Docker Compose (for database features)

## Build

```bash
go build -o goku ./
```

---

## Phase 1: JSON/YAML Converter

### Usage

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

Example: `testdata/config.json` -> `testdata/config.yaml`

### Notes

- The input and output formats cannot be the same.
- Supported extensions: `.json`, `.yaml`, `.yml`.

---

## Phase 2: Database Operations

### Docker Setup

Start the PostgreSQL database:

```bash
# Start database (runs in background)
docker compose up -d

# Check status
docker compose ps

# View logs
docker compose logs -f
```

Stop the database:

```bash
docker compose down
```

To force recreate the container:

```bash
docker compose down -v && docker compose up -d
```

### Database Configuration

The default configuration uses:

- **Host**: localhost
- **Port**: 5433
- **User**: postgres
- **Password**: postgres
- **Database**: goku
- **SSL**: disabled

Override with environment variables:

```bash
export GOKU_DB_HOST=localhost
export GOKU_DB_PORT=5433
export GOKU_DB_USER=postgres
export GOKU_DB_PASSWORD=postgres
export GOKU_DB_NAME=goku
export GOKU_DB_SSLMODE=disable
```

### Commands

#### Save (Create)

```bash
./goku save -i <input-file>
```

Save a new resource from a JSON or YAML file.

#### Dump (List All)

```bash
./goku dump
```

Retrieve all resources as a JSON array.

#### Get (Retrieve Single)

```bash
./goku get -d <id>
```

Retrieve a single resource by its ID.

#### Update

```bash
./goku update -d <id> -i <input-file>
```

Update an existing resource by ID with data from a file.

#### Delete

```bash
./goku delete -d <id>
```

Delete a resource by its ID.

### Examples

```bash
# Save a new resource
./goku save -i testdata/config.yaml

# Get all resources
./goku dump

# Get specific resource
./goku get -d 1

# Update a resource
./goku update -d 1 -i testdata/config.json

# Delete a resource
./goku delete -d 1
```

### Quick Start

```bash
# 1. Start database
docker compose up -d

# 2. Build goku
go build -o goku ./

# 3. Save a resource
./goku save -i testdata/config.json

# 4. List all resources
./goku dump

# 5. Get specific resource
./goku get -d 1

# 6. Stop database (when done)
docker compose down
```
