# go-to-dart

Go-to-dart helps you convert Go structs to Dart classes that can be used with [json_serializable](https://pub.dev/packages/json_serializable).

## Features

- Supports only structs in the same package (no generics or embedded structs yet)
- Supports primitives, slices, maps, and pointers
- Support some other arbitrary types such as `time.Time` and `mo.Option` (easy to extend!)
- Support for `json` tags

Need something more? Please open an issue or even better, a PR!

## Installation

```bash
go install github.com/11wizards/go-to-dart@latest
```

The above command will install go-to-dart in your `$GOPATH/bin` directory. Make sure that directory is in your `$PATH`.

## Usage

For plain JSON serialization, you can use the `json` mode.
```bash
go-to-dart -i ./examples/user -o ./examples/user -m json
```

For serialization that works with Firestore, you can use the `firestore` mode.
```bash
go-to-dart -i ./examples/user -o ./examples/user -m firestore
```

## Example

Running the command above would take the package `./examples/user` below and generate a file `./examples/user/user.dart`.

```go
package user

import (
	"time"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Options   map[string]string
	Tags      []string
}
```

Contents of `./examples/user/user.dart`:
```dart
import 'package:json_annotation/json_annotation.dart';

part 'user.go.g.dart';

@JsonSerializable(explicitToJson: true)
class User {
	@JsonKey(name: "ID")final int id;
	@JsonKey(name: "Name")final String name;
	@JsonKey(name: "Email")final String email;
	@JsonKey(name: "Password")final String password;
	@JsonKey(name: "CreatedAt")final DateTime createdAt;
	@JsonKey(name: "UpdatedAt")final DateTime updatedAt;
	@JsonKey(name: "DeletedAt")final DateTime? deletedAt;
	@JsonKey(name: "Options", defaultValue: <String, String>{})final Map<String, String> options;
	@JsonKey(name: "Tags", defaultValue: <List<String>>[])final List<String> tags;
	
	User({
		required this.id,
		required this.name,
		required this.email,
		required this.password,
		required this.createdAt,
		required this.updatedAt,
		this.deletedAt,
		required this.options,
		required this.tags,
	});
	
	Map<String, dynamic> toJson() => _$UserToJson(this);
	
	factory User.fromJson(Map<String, dynamic> json) => _$UserFromJson(json);
}
```
