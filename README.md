# Advanced Exercise

## Overview

For this week, we will be expanding on the app you built in the *Introduction* section. 
The initial app is a single binary that only converts **JSON** files to **CSV**. This week
we will add some of the concepts that you will be exposed during the Advanced part of the
training. 

## Concepts

In improving your application you will be using the following concepts:

* GO Routines
* Packages
* Project Structure
* Building multiple binaries
* Core Packages
* Advanced Flags

## Boilerplate Code

If this code is not in your go workspace (usually located in $HOME/go/src), you will need to create init
a go module inside the working directory. You can do this by running:

```bash
go mod init [YOUR APP NAME GOES HERE]
```

The boiler plate contains the following structure:

| bin
| cmd
-- | csv-to-json
-- -- | main.go
-- | json-to-csv
-- -- | main.go
| internal
-- | convert.go

The bin folder will be reserved for our binaries to be compiled to. You can compile both by running:

```bash
go build -o ./bin ./...
```

You can additionally run either binary without compiling with:

```bash
go run ./cmd/csv-to-json/main.go
```

or 

```bash
go run ./cmd/json-to-csv/main.go
```

The internal folder will be where we keep all of our non-main logic. You can spread this logic
across files in the internal package, or even create new ones. The possibilities are endless and really up
to you how you want to abstract the code.


## Requirements

Your application will be split into two separate binaries:

* **json-to-csv** - Converts a JSON file to CSV
* **csv-to-json** - Converts a CSV file to JSON

Your application should be able to convert multiple files on a single run (batches)

Your binaries should run using the following commands:

`csv-to-json --out-dir /path/to/output file1.csv file2.csv file3.csv`
`json-to-csv --out-dir /path/to/output file1.json file2.json file3.json`

Your application should use GO routines for processing the files. This ensures that the batches
are converted concurrently. There are various different ways to implement this, such as using *channels*, *worker groups*, etc. 
I'll leave it up to you to choose. :) 

### Optional

A nice extra (if you have the time) would be to add some unit tests for your code. Check out the GO: Generate Unit for [Package/File/Function] command in VSCode.
