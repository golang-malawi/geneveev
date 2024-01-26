geneveev
========

A command-line tool to generate [Yup](https://yup-docs.vercel.app/docs/Api/yup) object schemas from Go structs that use [go-playground/validator](https://github.com/go-playground/validator) tags.

Useful for setting up validation stuff for front-end development.

> NOTE: still in very early development

## Installation

```sh
go install github.com/golang-malawi/geneveev@latest
```

## Usage

```sh
$ geneveev -d /path/to/package/with/validated/structs/
```

## Features

- [x] Generates basic yup object schemas from basic Go structs

## TODO

- Support nested structs (map to mixed, objects)
- Support ref fields

## Contributors

- [Zikani Nyirenda Mwase](https://github.com/zikani03) - Maintainer

## LICENSE

MIT LICENSE

