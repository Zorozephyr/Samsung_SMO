# 📘 Go (Golang) — Complete Revision Notes

---

## Table of Contents

1. [Go CLI Commands](#1--go-cli-commands)
2. [Packages & Project Structure](#2--packages--project-structure)
3. [Variables & Types](#3--variables--types)
4. [Arrays & Slices](#4--arrays--slices)
5. [Custom Types & Receivers](#5--custom-types--receivers)
6. [Functions](#6--functions)
7. [Structs](#7--structs)
8. [Pointers](#8--pointers)
9. [Maps](#9--maps)
10. [Interfaces](#10--interfaces)
11. [The `io` Package — Reader & Writer](#11--the-io-package--reader--writer)
12. [Goroutines](#12--goroutines)
13. [Channels](#13--channels)
14. [Select Statement](#14--select-statement)
15. [Sync Package — Mutex & Defer](#15--sync-package--mutex--defer)
16. [Context](#16--context)
17. [Generics (Go 1.18+)](#17--generics-go-118)
18. [Interview-Style Q&A](#18--interview-style-qa)
19. [Capstone Example — FastestFetch](#19--capstone-example--fastestfetch)

---

## 1. 🛠 Go CLI Commands

| Command        | Description                                          |
| -------------- | ---------------------------------------------------- |
| `go run`       | Compiles and executes one or two files               |
| `go build`     | Compiles and builds an executable binary              |
| `go fmt`       | Formats all the code in the repository               |
| `go install`   | Compiles and installs a package                      |
| `go get`       | Downloads the raw source code of someone else's package |
| `go test`      | Runs any tests associated with the current project   |
| `go mod init`  | Initializes a new Go module (e.g. `go mod init cards`) |

---

## 2. 📦 Packages & Project Structure

- **Package == Project == Workspace**
- Two types of packages:
  - **Executable** → Generates a runnable file. Uses `package main` and **must** have a `func main()`.
  - **Reusable** → Library code used as helpers. Can be any other package name.

### Cross-File Access Within a Package

Files in the **same package** can call each other's functions directly — no imports needed.

**main.go**
```go
package main

func main() {
    printState()
}
```

**state.go**
```go
package main

import "fmt"

func printState() {
    fmt.Println("California")
}
```

> 💡 Both files are `package main`, so `main.go` can call `printState()` from `state.go` without importing it.

---

## 3. 🔤 Variables & Types

### Declaration Styles

```go
// Explicit type declaration
var name string = "Vishnu"

// Short declaration (type inferred) — only inside functions
name := "Vishnu"

// Declare without value (gets zero value)
var age int   // age == 0
```

### Rules to Remember

| Rule | Example |
|------|---------|
| `:=` is for **initial** assignment only | `color := "red"` |
| `=` is for **reassignment** | `color = "blue"` |
| Global variables can be **declared** but not **assigned** outside functions | `var x int` ✅ / `x := 5` ❌ (at package level) |

### Zero Values

| Type    | Zero Value |
| ------- | ---------- |
| `string`| `""`       |
| `int`   | `0`        |
| `float` | `0`        |
| `bool`  | `false`    |

---

## 4. 📋 Arrays & Slices

| Feature | Array | Slice |
|---------|-------|-------|
| Length  | **Fixed** at compile time | **Dynamic** — can grow and shrink |
| Syntax  | `[3]string{"a","b","c"}` | `[]string{"a","b","c"}` |

> ⚠️ Every element in an array/slice must be of the **same type**.

### Slice Operations

```go
// Create a slice
cards := []string{"Ace of Spades", "Two of Hearts"}

// Append (returns a NEW slice — does NOT modify original)
cards = append(cards, "Three of Diamonds")

// Iterate
for index, card := range cards {
    fmt.Println(index, card)
}

// Sub-slicing
cards[0:2]   // from index 0 up to (not including) index 2
cards[:2]    // same as above
cards[1:]    // from index 1 to end
```

---

## 5. 🧩 Custom Types & Receivers

### Creating a Custom Type

```go
type deck []string
```

### Adding a Receiver Function

```go
func (d deck) print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}
```

```go
cards := deck{"Ace of Spades", "Two of Hearts"}
cards.print() // Every variable of type 'deck' can now call print()
```

> 📝 **Convention:** The receiver variable is typically a 1–2 letter abbreviation of the type (`d` for `deck`).

---

## 6. ⚡ Functions

### Multiple Return Values

```go
func deal(d deck, handSize int) (deck, deck) {
    return d[:handSize], d[handSize:]
}

hand, remainingDeck := deal(cards, 5)
```

### Type Conversion

```go
greeting := "Hello"
byteSlice := []byte(greeting)  // Convert string → []byte
```

> `[]byte` (byte slice) is the common way to represent raw string data for file I/O.

### Writing to a File

```go
import "os"

os.WriteFile("my_file.txt", []byte("Hello!"), 0666)
```

### Testing

- Test files must end with `_test.go`
- Run with `go test`

```go
// deck_test.go
func TestNewDeck(t *testing.T) {
    d := newDeck()
    if len(d) != 52 {
        t.Errorf("Expected 52, but got %v", len(d))
    }
}
```

---

## 7. 🏗 Structs

### Declaration

```go
type contactInfo struct {
    email   string
    zipCode int
}

type person struct {
    firstName string
    lastName  string
    contact   contactInfo   // Nested struct
}
```

### Initialization (Multiple Ways)

```go
// 1. Positional (not recommended for readability)
alex := person{"Alex", "Anderson", contactInfo{}}

// 2. Named fields (recommended)
alex := person{
    firstName: "Alex",
    lastName:  "Anderson",
    contact: contactInfo{
        email:   "alex@gmail.com",
        zipCode: 94000,
    },
}

// 3. Zero-value then assign
var alex person
alex.firstName = "Alex"
alex.lastName = "Anderson"
```

### Print Struct with Field Names

```go
fmt.Printf("%+v", alex)
// Output: {firstName:Alex lastName:Anderson contact:{email:alex@gmail.com zipCode:94000}}
```

### Struct Receiver

```go
func (p person) print() {
    fmt.Printf("%+v", p)
}
```

---

## 8. 🔗 Pointers

> 🚨 **Go is a pass-by-value language.** When you pass a struct to a function, a **copy** is made.

### The Problem

```go
func (p person) updateName(newName string) {
    p.firstName = newName  // ❌ Only updates the COPY, not the original!
}
```

### The Solution — Pointer Receivers

```go
func (p *person) updateName(newName string) {
    (*p).firstName = newName  // ✅ Dereferences pointer, updates original
}
```

### Pointer Operators

| Operator | Meaning | Example |
|----------|---------|---------|
| `&variable` | Get the **memory address** of a variable | `ptr := &jim` |
| `*pointer`  | Get the **value** at a memory address (dereference) | `fmt.Println(*ptr)` |
| `*type`     | Describes a **pointer type** in function signatures | `func (p *person)` |

### Go's Shortcut

```go
jim := person{firstName: "Jim"}

// Both of these work — Go automatically converts:
jimPointer := &jim
jimPointer.updateName("Jimmy")

jim.updateName("Jimmy")  // ✅ Go auto-converts jim → &jim for pointer receivers
```

---

## 9. 🗺 Maps

- **All keys** must be the same type
- **All values** must be the same type (can differ from key type)

### Creation

```go
// Literal
colors := map[string]string{
    "red":   "#ff0000",
    "green": "#00ff00",
}

// Using make
colors := make(map[string]string)

// Zero-value declaration
var colors map[string]string
```

### Operations

```go
colors["white"] = "#ffffff"      // Add / Update
delete(colors, "red")            // Delete
```

### Iterating

```go
func printMap(c map[string]string) {
    for key, value := range c {
        fmt.Println(key, ":", value)
    }
}
```

### Maps vs Structs

| Feature       | Map                        | Struct                      |
| ------------- | -------------------------- | --------------------------- |
| Keys          | Same type, dynamic         | Fixed field names at compile time |
| Values        | Same type                  | Can be different types       |
| Iterable      | ✅ Yes                     | ❌ No                       |
| Use Case      | Collection of related items | Represent a "thing" with properties |

---

## 10. 🔌 Interfaces

> Interfaces define a **contract** — if a type has the required methods, it **automatically** satisfies the interface (implicit implementation).

### Defining and Using an Interface

```go
type bot interface {
    getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (eb englishBot) getGreeting() string {
    return "Hello!"
}

func (sb spanishBot) getGreeting() string {
    return "¡Hola!"
}

func printGreeting(b bot) {
    fmt.Println(b.getGreeting())
}

func main() {
    eb := englishBot{}
    sb := spanishBot{}

    printGreeting(eb)  // "Hello!"
    printGreeting(sb)  // "¡Hola!"
}
```

### Key Rules

- Interfaces are **implicit** — no `implements` keyword needed
- Interfaces are **not generic types**
- Interfaces are a **contract** to help manage types
- Interfaces enable **decoupling** — you can mock dependencies for testing without the original author defining an interface

### Interface Internals

An interface variable is actually a **pair**:

| Component | Description |
|-----------|-------------|
| **Type**  | The concrete type (e.g., `*Dog`) |
| **Value** | The actual data (the pointer to `Dog`) |

---

## 11. 📡 The `io` Package — Reader & Writer

### The HTTP Response Struct

```
resp, err := http.Get("https://example.com")

resp.Status     → string
resp.StatusCode → int
resp.Body       → io.ReadCloser
```

### Interface Hierarchy

```
io.ReadCloser (interface)
├── io.Reader  →  Read([]byte) (int, error)
└── io.Closer  →  Close() error
```

### Reading a Response Body

```go
// Manual approach — make a byte slice and read into it
bs := make([]byte, 99999)
resp.Body.Read(bs)
fmt.Println(string(bs))
```

> ⚠️ `Read` does **not** auto-resize the slice. You must pre-allocate enough space.

### The Better Way — `io.Copy`

```go
io.Copy(os.Stdout, resp.Body)
```

```
func Copy(dst Writer, src Reader) (written int64, err error)
```

> `os.Stdout` is of type `*os.File`, which has a `Write` method → satisfies the `io.Writer` interface.

---

## 12. 🚀 Goroutines

> A running Go program starts as a **single goroutine** (the `main` goroutine).

### Launching a Goroutine

```go
for _, link := range links {
    go checkLink(link)  // Launches a new goroutine for each call
}
```

### How the Go Scheduler Works

- Default: uses **one CPU core** (configurable with `GOMAXPROCS`)
- Detects **blocking calls** → pauses that goroutine → runs another
- With multiple cores, goroutines run **in parallel**

> ⚠️ **If the main goroutine exits, ALL child goroutines are killed immediately.**

### Why Goroutines Are Special

| Property | Detail |
|----------|--------|
| **Lightweight** | Start with ~2KB stack that grows/shrinks dynamically |
| **User-Space** | Managed by the Go runtime, not the OS kernel |
| **M:N Scheduling** | Maps M goroutines onto N OS threads |

### The GMP Model

| Entity | Name | Role |
|--------|------|------|
| **G** | Goroutine | Contains stack, instruction pointer, and code to run |
| **M** | Machine | An OS thread |
| **P** | Processor | Local scheduler with its own run queue of goroutines |

- Number of P's = number of CPU cores (set by `GOMAXPROCS`)
- Each **P** has a **Local Run Queue (LRQ)** — no global lock needed → extremely fast
- This design allows Go to **scale linearly** across CPU cores

---

## 13. 📨 Channels

> Channels allow goroutines to **communicate** with each other safely.

### Basics

```go
c := make(chan string)       // Create an unbuffered channel of type string

c <- "Hello"                 // Send a value INTO the channel
msg := <-c                   // Receive a value FROM the channel (BLOCKING)
fmt.Println(<-c)             // Receive and print directly
```

### Unbuffered vs Buffered Channels

```go
// Unbuffered (default) — sender blocks until receiver is ready
ch := make(chan int)

// Buffered — sender blocks only when buffer is full
ch := make(chan int, 5)   // Capacity of 5
```

| Type | Sender Blocks When... | Receiver Blocks When... |
|------|----------------------|------------------------|
| Unbuffered | No receiver is ready | No sender has sent |
| Buffered | Buffer is full | Buffer is empty |

> 💡 Buffered channels create **natural backpressure** — prevents memory overflow from unprocessed tasks.

### Iterating Over a Channel

```go
for link := range c {
    go checkLink(link, c)
}
```

### Function Literals (Anonymous Goroutines)

```go
for _, l := range links {
    go func(link string) {
        time.Sleep(5 * time.Second)
        checkLink(link, c)
    }(l)   // Pass 'l' as argument to avoid closure issues
}
```

---

## 14. 🔀 Select Statement

The `select` statement lets a goroutine wait on **multiple channel operations**.

```go
select {
case msg := <-ch1:
    fmt.Println("Received from ch1:", msg)
case msg := <-ch2:
    fmt.Println("Received from ch2:", msg)
default:
    fmt.Println("No channel ready")
}
```

### Behavior

| Scenario | What Happens |
|----------|-------------|
| One case ready | That case runs |
| Multiple cases ready | One is chosen **at random** (prevents starvation) |
| No case ready + `default` | `default` runs immediately (non-blocking) |
| No case ready + no `default` | `select` **blocks** until a case is ready |

---

## 15. 🔒 Sync Package — Mutex & Defer

### The Problem — Race Conditions

If 100 goroutines do `count++` simultaneously, the result is unpredictable.

### Solution — `sync.Mutex`

```go
var mu sync.Mutex
var count int

func increment() {
    mu.Lock()         // 🔒 Only one goroutine can enter
    count++           // Critical section — safe to modify
    mu.Unlock()       // 🔓 Release the lock
}
```

### Safer with `defer`

> `defer` schedules a function call to run **when the surrounding function exits** — even if it panics.

```go
func increment() {
    mu.Lock()
    defer mu.Unlock()  // ⏰ Guaranteed to run on function exit

    // Even if this code panics, Unlock still happens
    count++
}
```

> 💡 `defer` pushes calls onto a stack. They execute in **LIFO** (last-in, first-out) order.

---

## 16. 🌐 Context

> Context ties **concurrency** and **interfaces** together. It's used to manage deadlines, cancellation, and request-scoped values across goroutines.

### How `ctx.Done()` Works

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
```

| State | `ctx.Done()` Channel |
|-------|--------------------|
| **Alive** | Blocks (no values) |
| **Canceled** | Closed immediately |

### Using Context with Select

```go
select {
case <-ctx.Done():
    // 🛑 Timeout or manual cancel — stop wasting resources
    return ctx.Err()
case result := <-databaseChannel:
    // ✅ Got data before the deadline
    return result
}
```

### The Broadcast Power of `close()`

> **Why `close()` instead of sending a value?**

| Method | Effect on 100 waiting goroutines |
|--------|--------------------------------|
| `ch <- true` | Only **1** goroutine receives it |
| `close(ch)` | **All 100** goroutines get the signal 📡 |

Closing a channel acts like a **broadcast** — every goroutine reading from it immediately receives the zero value.

---

## 17. 🧬 Generics (Go 1.18+)

### The Problem Before Generics

```go
// Had to write separate functions for each type 😩
func SumInts(list []int) int { ... }
func SumFloats(list []float64) float64 { ... }
```

### The Old Hack — `interface{}`

```go
func Add(a, b interface{}) int {
    aVal := a.(int)   // ⚠️ Type assertion — panics if wrong type!
    bVal := b.(int)
    return aVal + bVal
}
```

> `interface{}` requires zero methods → every type satisfies it. But you lose type safety.

### The Modern Solution — Generics

```go
func Add[T int | float64](a, b T) T {
    return a + b
}

// Usage
Add(3, 5)       // T inferred as int → returns 8
Add(1.5, 2.5)   // T inferred as float64 → returns 4.0
```

| Approach | Type Safe | Flexible | Compile-Time Check |
|----------|-----------|----------|--------------------|
| Separate functions | ✅ | ❌ | ✅ |
| `interface{}` | ❌ | ✅ | ❌ |
| **Generics** | ✅ | ✅ | ✅ |

---

## 18. 💬 Interview-Style Q&A

### Q1: What happens when you send on an unbuffered channel with no receiver?

> The **sender blocks** at that exact line. It's like a synchronous phone call — you can't hang up until someone picks up. This guarantees that send and receive happen **simultaneously**.

### Q2: What happens when a buffered channel (capacity 2) gets a 3rd send with no receives?

> The **sender blocks**. This creates **natural backpressure** — if the consumer is slow, the producer is forced to pause, preventing memory exhaustion.

### Q3: How do you protect shared memory from concurrent access?

> Use `sync.Mutex`. Lock before accessing, unlock after. Always use `defer mu.Unlock()` to guarantee the lock is released even on panics.

### Q4: Why does context use `close()` instead of sending a value?

> **Broadcasting.** Sending a value notifies only 1 goroutine. Closing the channel notifies **all** goroutines waiting on it simultaneously.

---

## 19. 🏆 Capstone Example — FastestFetch

> This example ties together **interfaces**, **goroutines**, **channels**, **context**, and **select**.

```go
type Fetcher interface {
    Fetch(ctx context.Context) (string, error)
}

// FastestFetch calls multiple fetchers concurrently and returns
// the first successful result, or an error if the context expires.
func FastestFetch(ctx context.Context, fetchers []Fetcher) (string, error) {
    c := make(chan string, len(fetchers)) // Buffered to avoid goroutine leaks

    for _, fetcher := range fetchers {
        go func(f Fetcher) {
            str, err := f.Fetch(ctx)
            if err != nil {
                return // Skip failed fetches
            }
            c <- str
        }(fetcher)
    }

    select {
    case <-ctx.Done():
        return "", ctx.Err()       // ⏰ Timeout or cancellation
    case result := <-c:
        return result, nil         // 🏁 First successful result wins
    }
}
```

### Concepts Used

| Concept | Where |
|---------|-------|
| **Interface** | `Fetcher` — any type with `Fetch()` qualifies |
| **Goroutines** | Each fetcher runs concurrently |
| **Buffered Channel** | Collects results without blocking senders |
| **Context** | Propagates cancellation/timeout |
| **Select** | Races between context deadline and first result |

---

> 🎯 **Tip:** Re-read each section's code examples and try to write them from memory. That's the fastest path to fluency in Go!
