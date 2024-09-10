# Golang

## Introduction

- Released on Nov 2009
- By Rob Pike, Ken Thompson (Creators of Unix System) at Google
- Golang is a compiled, memory managed (has GC), statically typed language built for simplicity and perfomance
- Highly opinionated, everything has conventions
- A mix between: Java, C++, Python

- Go utilizes multiple cores (meaning it runs multiple processes)
    - leverages IPC - Inter process communication
- Supports concurrency but is not thread-safe by default
- Go binaries are statically linked
    - No external dependencies
    - Just specify OS, and ARCH
```
export GOOS=linux
export GOARCH=amd64
```

- Why go?
    - Super fast builds
    - Low memory footprint
    - Simple syntax
    - Built-in concurrency model
    - Highly popular

- Pain Points
    - Errors
    - Generics

- Go is a Functional language
    - Building blocks: Types, functions, packages
    - Has 3/4 OOPS concepts: encapsulation, abstraction and polymorphism
    - No inheritance !
        - Use composition instead

## Project Structure

- go.mod - required packages to run the project
- go.sum - locks versions for reproducible builds (should be used)