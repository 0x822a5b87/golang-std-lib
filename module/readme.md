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

## Tutorial: Getting started with fuzzing

> In this tutorial, you’ll write a fuzz test for a simple function, run the go command, and debug and fix issues in the code.

```go
package main

import "fmt"

func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}
```

```go
package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev := Reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}
```

### Add a fuzz test

> The unit test has limitations, namely that each input must be added to the test by the developer. 
>
> **One benefit of fuzzing is that it comes up with inputs for your code, and may identify edge cases that the test cases you came up with didn’t reach.**

```go
func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc)  // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
```

#### Run test code

```bash
go test -fuzz=Fuzz
```

>fuzz: elapsed: 0s, gathering baseline coverage: 0/17 completed
>failure while testing seed corpus entry: FuzzReverse/7d02a7200a7179081b177bc32614070b667f0b30d47a72630b631308f2fd781a
>fuzz: elapsed: 0s, gathering baseline coverage: 3/17 completed
>--- FAIL: FuzzReverse (0.03s)
>--- FAIL: FuzzReverse (0.00s)
>  reverse_test.go:36: Reverse produced invalid UTF-8 string "\xbc\xc6"
>
>
>
>FAIL
>exit status 1
>FAIL    example/fuzz    0.411s

### Fix the invalid string error

> The current `Reverse` function reverses the string byte-by-byte, and therein lies our problem. In order to preserve the UTF-8-encoded runes of the original string, we must instead reverse the string rune-by-rune.
>
> 
>
> To examine why the input (in this case, the Chinese character `泃`) is causing `Reverse` to produce an invalid string when reversed, you can inspect the number of runes in the reversed string.

```go
func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
```

### Fix the double reverse error



## Naming a module

> When you run go mod init to create a module for tracking dependencies, you specify a module path that serves as the module’s name. 
>
> **The module path becomes the import path prefix for packages in the module.**Be sure to specify a module path that won’t conflict with the module path of other modules.
