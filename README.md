# GoLearn
### Go Lang Basics
1. Compiled
2. No vm required
3. Executables are different for OS
### For what types of apps it can be used?
1. Sys app
2. Web app
3. Cloud based app
#### Is object-oriented and not object-oriented
### Hello World
1. Install go
2. Create directory
3. Create a file `main.go`
4. Go to terminal and `go mod init hello` - this is similar to `npm init` and instead of hello, it can be anything.
   - Here a module `hello` is created
   - A file `go.mod` will be added.
   - This `go.mod` will be almost like `package.json`
   - It contains module name `module hello`
   - And go version `go 1.19`
6. Entry point is `main` package and `main` function
7. ```
   package main
   import "fmt"
   
   func main() {
      fmt.Println("Hello World...!")
   }
   ```
7. To run the `go run main.go` or `go run .`
8. Use `go help` to know more commands
9. `go env GOPATH` shows the folder in which the imports are resolved. Equivalent to `.m2` or `node-modules`
   By default if nothing is set, will point to `users home directory/go`
10. fmt = Package formatted I/O
    - **Package**: Directory containing a group of .go files
    - **Module** : Collection of packages
### Lexer:
1. Analyze code before compile for any problems and cleanup
2. Puts ; internally for each step and removes those ; added by us

### Types
1. Variables should be declared before use.
2. Strict in type 
3. If a variable name starts with Capital letter, it indicates that the variable is public. Otherwise, its package private.
4. From 1.5, there is `internal` which ensures that only files in specific part of package can access those inside `internal
      - eg: package a/b/c/internal/d
      - Here, files inside a/b/c can use those inside d.
      - But files inside a/b/e cannot access those inside d.

```
Types:
Basic
   - String
   - Bool
   - Integer
      - uint8, uint16, uint32, uint64 (unsigned) (each of them can hold integers in the range of 0 - 2 ^ n -1 )
      - int8, int16, int32, int64
      - byte = uint8 (alias and hence for a var of type byte prints `byte` for %T )
      - int is unique and can hold 32 bits and `not equal to uint32`
      - uintptr (unsigned)
   - Float
      - float32
      - float64
   - Complex
Advanced
   - Array
   - Slice
   - Maps
   - Structs
   - Pointers
Functions
Channels
```
### Variables and Constants
1. var x <type> = value
   fmt.printf("%T", x) //prints the type
2. var x= "" // implicit type
3. x := "" // no var style
4. const SampleConst int = 10 // will be public
5. const sampleConst int = 10 // will be within package

### How to accept input from users
1. Use package `bufio`
```
   import (
   "bufio"
   "fmt"
   )
   
   func main(){
      reader := bufio.newReader(os.Stdin)
      // below way of getting return and error is called comma ok | err ok way 
      input, _ := reader.ReadString('\n');
      fmt.println(input);
   }
   
```
### How can we convert string to int etc. Just like Integer.parseInt() and so on
1. Package `strconv` with `strconv.parseFloat(str)` etc.
2. Remove new line etc. use `strings` package. `strings.trimSpace(str)`

### Time package
1. time.now() gives current time
2. time.now().format("01-02-2006 15:04:05 Monday") prints time in format. The string given into the format() function is
   layout string and is defined by language spec

### Build stand-alone apps
1. For any os in which we do dev
   `go build` is enough 
2. To build it for another os,
   `GOOS="linux" go build` // gives linux executable
3. `GOOS="windows" go build` // gives windows executable version of our program
4. `GOOS` is an env variable and can be listed doing `go env`. Mac is treated as GOOS="darwin"
5. Executable will have name of `package`. Eg: timeStudy.exe
6. 
