# go_refresh

**Go 1.14.1**

This contains personal notes for Go along with course code.

1. [courses](./courses) contains Coursera and schoolwork code.
1. [experiments](./experiments) contains some code I used to experiment with and learn about Go concepts involving REST APIs a long time ago.
1. [examples](./examples) contains some recent examples demonstrating how to do basic things including the topics below.

> **Note:** executing the supplied Go code requires navigating to `PROJECT_DIR/src` or setting the `GO_PATH` appropriately.

## Floats

Use decimals when multiplying floats. [Go to](courses/8-accel/accel.go).

## OOD

Go does not have inheritance - polymorphism is achieved through flexible structs where an attribute plays the role of a type. [Go to](courses/9-structsreceivers/main.go).

Encapsulation is achieved using structs, keeping method and type names lowercase (to prevent exporting), and using receivers. [Go to](examples/5-receivertype/main.go)

## Array and Maps Initialization

Use `make()` for Arrays and Maps:

1. [Arrays](courses/4-slice/slice.go)
1. [Maps](courses/10-structsreceivers/main.go)

## Input

Use:

```go
bscan := bufio.NewScanner(os.Stdin)

// Continually reads input - try multiple times
for bscan.Scan() { }
```

or for singular inputs:

```go
fileName := ""
fmt.Println("Enter filename!")
fmt.Scan(&fileName)
```