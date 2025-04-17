geneveev
========

A command-line tool to generate object schemas from Go structs.
`It is useful for generating code for validation for front-end development 
and for creating DTOs (Data Transfer Objects) interacting with APIs built with Go

Currently it supports generating JavaScript and Dart from structs containing [go-playground/validator](https://github.com/go-playground/validator) tags. 

 - [Yup](https://yup-docs.vercel.app/docs/Api/yup) schemas
 - [Zod](https://zod.dev) schemas
 - [Dart](https://dart.dev) classes via the embedded go-to-dart functionality


## Installation

```sh
go install github.com/golang-malawi/geneveev@latest
```

## Usage


**Generating Yup Schemas**

```sh
$ geneveev generate yup -d /path/to/package/with/validated/structs/ --output-dir ./yup-schemas
```

**Generating Zod Schemas**

```sh
$ geneveev generate zod -d /path/to/package/with/validated/structs/ --output-dir ./zod-schemas
```

**Generating Dart classes**

```sh
$ geneveev generate dart -d /path/to/package/with/validated/structs/ --output-dir ./dart-classes
```

## Example

Generate from this

![[]](./struct.png)

To this:

![[]](./yup.png)

## Features

### Yup Schema

- [x] Generates basic yup object schemas from basic Go structs

### Zod Schema

- [x] Generates basic zod object schemas from basic Go structs

### go-to-dart

The Go-to-dart implementation helps you convert Go structs to Dart classes that can be used with [json_serializable](https://pub.dev/packages/json_serializable).

- Supports only structs in the same package (no generics or embedded structs yet)
- Supports primitives, slices, maps, and pointers
- Support some other arbitrary types such as `time.Time` and `mo.Option` (easy to extend!)
- Support for `json` tags


## Contributors

- [Zikani Nyirenda Mwase](https://github.com/zikani03) - Maintainer

## LICENSE

MIT LICENSE

