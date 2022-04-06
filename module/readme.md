# readme

> - [Managing dependencies](https://go.dev/doc/modules/managing-dependencies)
> - [Tutorial: Create a Go module](https://go.dev/doc/tutorial/create-module)

## Tutorial: Create a Go module

> Go code is grouped into packages, and packages are grouped into modules. Your module specifies dependencies needed to run your code, including the Go version and the set of other modules it requires.

```bash
go mod init example.com/greetings
```

```go
package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
```

## Call your code from another module

```
<home>/
 |-- greetings/
 |-- hello/
```

```bash
go mod init example.com/hello
```

```go
package main

import (
    "fmt"

    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
```

```bash
go mod edit -replace example.com/greetings=../greetings
```

> The command specifies that `example.com/greetings` should be replaced with `../greetings` for the purpose of locating the dependency. 

```
module example.com/hello

go 1.17

replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000
```

> From the command prompt in the hello directory, run the [`go mod tidy` command](https://go.dev/ref/mod#go-mod-tidy) to synchronize the `example.com/hello` module's dependencies, adding those required by the code, but not yet tracked in the module.

```bash
go mod tidy
```

## Tutorial: Getting started with multi-module workspaces

> With multi-module workspaces, you can tell the Go command that you’re writing code in multiple modules at the same time and easily build and run code in those modules.

### create workspace and first module

```bash
mkdir workspace
cd workspace

mkdir hello
cd hello
go mod init example.com/hello
#go: creating new go.mod: module example.com/hello

go get golang.org/x/example
```

```go
package main

import (
    "fmt"

    "golang.org/x/example/stringutil"
)

func main() {
    fmt.Println(stringutil.Reverse("Hello"))
}
```

### Create the workspace

```bash
go work init ./hello
```

### Download and modify the `golang.org/x/example` module

```go
git clone https://go.googlesource.com/example
go work use ./example

```



## Naming a module

> When you run go mod init to create a module for tracking dependencies, you specify a module path that serves as the module’s name. 
>
> **The module path becomes the import path prefix for packages in the module.**Be sure to specify a module path that won’t conflict with the module path of other modules.
