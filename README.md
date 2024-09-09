# Stack

A simple, generic `Stack` implementation in Go, fully tested and ready to use.

## Installation

To install the `Stack` package, use the `go get` command:

```bash
go get github.com/andreimerlescu/stack
```

Then, import it into your Go project:

```go
import "github.com/andreimerlescu/stack"
```

## Usage

The `Stack` is a simple LIFO (Last In, First Out) data structure. It supports generic types, so you can create a stack of any type (e.g., `int`, `string`, or custom types).

### Example: Stack of Integers

```go
package main

import (
	"fmt"
	"github.com/andreimerlescu/stack"
)

func main() {
	// Create a stack of integers
	intStack := stack.Stack[int]{}

	// Push elements onto the stack
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)

	// Peek at the top element
	top, _ := intStack.Peek()
	fmt.Println("Top element:", top) // Output: Top element: 30

	// Pop elements off the stack
	val, _ := intStack.Pop()
	fmt.Println("Popped:", val) // Output: Popped: 30

	// Check if the stack is empty
	fmt.Println("Is stack empty?", intStack.IsEmpty()) // Output: Is stack empty? false
}
```

### Example: Stack of Strings

```go
package main

import (
	"fmt"
	"github.com/andreimerlescu/stack"
)

func main() {
	// Create a stack of strings
	stringStack := stack.Stack[string]{}

	// Push elements onto the stack
	stringStack.Push("Go")
	stringStack.Push("Rust")
	stringStack.Push("Python")

	// Pop elements off the stack
	val, _ := stringStack.Pop()
	fmt.Println("Popped:", val) // Output: Popped: Python

	// Peek at the top element
	top, _ := stringStack.Peek()
	fmt.Println("Top element:", top) // Output: Top element: Rust
}
```

### Methods

- `Push(item T)`: Pushes an item onto the stack.
- `Pop() (T, error)`: Pops the top item from the stack. Returns an error if the stack is empty.
- `Peek() (T, error)`: Returns the top item without removing it. Returns an error if the stack is empty.
- `IsEmpty() bool`: Returns `true` if the stack is empty, `false` otherwise.
- `Size() int`: Returns the number of elements in the stack.

## Running Tests

To run the unit tests and benchmarks for the `Stack` implementation, use the following commands:

```bash
go test ./...         # Run all unit tests
go test -bench=. ./... # Run benchmarks
```

## License

```txt
Copyright 2024 Andrei Merlescu (github.com/andreimerlescu)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

