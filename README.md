# Go package: [`github.com/strongo/dalgo2buntdb`](https://github.com/strongo/dalgo2buntdb)

[![Lint, Vet, Build, Test](https://github.com/strongo/dalgo2buntdb/actions/workflows/ci.yml/badge.svg)](https://github.com/strongo/dalgo2buntdb/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/strongo/dalgo2buntdb)](https://goreportcard.com/report/github.com/strongo/dalgo2buntdb)
[![GoDoc](https://godoc.org/github.com/strongo/dalgo2buntdb?status.svg)](https://godoc.org/github.com/strongo/dalgo2buntdb)

Bridge to [BuntDB](https://github.com/tidwall/buntdb) API
for [`github.com/strongo/dalgo`](https://github.com/strongo/dalgo) interface.

## What is BuntDB?

BuntDB is a low-level, in-memory, key/value store in pure Go. It persists to disk, is ACID compliant, and uses locking
for multiple readers and a single writer. It supports custom indexes and geospatial data. It's ideal for projects that
need a dependable database and favor speed over data size.

## What is `DALgo`?

[`DALgo`](https://github.com/strongo/dalgo) is a Database Abstraction Layer (in) Go (language).

## End-to-End testing

This package is covered by end-to-end tests
from [`github.com/strongo/dalgo-end2end-tests`](https://github.com/strongo/dalgo-end2end-tests).

## License

Open source under [MIT License](LICENSE) & free to use.
