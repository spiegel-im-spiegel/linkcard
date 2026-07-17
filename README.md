# [linkcard] -- A small CLI for Hugo link card workflows

[![ci status](https://github.com/spiegel-im-spiegel/linkcard/workflows/ci/badge.svg)](https://github.com/spiegel-im-spiegel/linkcard/actions)
[![build status](https://github.com/spiegel-im-spiegel/linkcard/workflows/build/badge.svg)](https://github.com/spiegel-im-spiegel/linkcard/actions)
[![CodeQL status](https://github.com/spiegel-im-spiegel/linkcard/workflows/CodeQL/badge.svg)](https://github.com/spiegel-im-spiegel/linkcard/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/linkcard/main/LICENSE)
[![GitHub release](http://img.shields.io/github/release/spiegel-im-spiegel/linkcard.svg)](https://github.com/spiegel-im-spiegel/linkcard/releases/latest)

This package is required Go 1.26 or later.

## Build and Install

```
$ go install github.com/spiegel-im-spiegel/linkcard@latest
```

## Binaries

See [latest release](https://github.com/spiegel-im-spiegel/linkcard/releases/latest).

## Usage

```
$ linkcard [flags] <url> [<url> ...]
```

## Features

- Fetch page metadata (`title`, `description`, `image_url`) from URL(s)
- Output link card data as JSON to stdout
- Optionally save/merge generated cards into a JSON data file
- Optionally download thumbnail images and store their relative path

## Common Flags

- `-u, --user-agent` : custom User-Agent string for HTTP fetch
- `-d, --data-path` : JSON file path to save/merge link card data
- `-i, --image-dir` : directory path to save downloaded thumbnail images
- `-b, --image-base-path` : base path prefix for `image_path` in output
- `-w, --image-width` : thumbnail width for image download
- `-r, --rating` : rating value (1-5; values above 5 are clamped)
- `-t, --page-title` : override page title in output
- `-c, --comment` : attach a comment string to output
- `--release-date` : attach release date string to output (no transformation)
- `-v, --version` : show version information

## Examples

```sh
$ linkcard https://example.com
```

```sh
$ linkcard -d ./cards.json -c "weekly memo" --release-date "2026-07-17" https://example.com
```

```sh
$ linkcard -i ./static/images -b /images -w 320 https://example.com
```

## Configuration File

If `./linkcard.json` exists, its values are loaded as defaults before parsing CLI flags.
CLI flags override the file values.

## Test

```
$ task test
```

## Modules Requirement Graph

[![dependency.png](./dependency.png)](./dependency.png)

[linkcard]: https://github.com/spiegel-im-spiegel/linkcard "spiegel-im-spiegel/linkcard: A small CLI for Hugo link card workflows"
