# Pointer

1. A pointer is simply a variable that holds the location in memory where a value is stored.
2. The **&** is the address operator. It precedes a value type and returns the address of the memory location where the value is stored:
```go
x := "hello"
pointerToX := &x
```
3. The * is the indirection operator. It precedes a variable of pointer type and returns the pointed-to value. This is called dereferencing:
```go
x := 10
pointerToX := &x
fmt.Println(pointerToX) // prints a memory address
fmt.Println(*pointerToX) // prints 10
z := 5 + *pointerToX
fmt.Println(z) // prints 15
```

4. Careful using pointer:
Don’t do this:
```go
func MakeFoo(f *Foo) error {
    f.Field1 = "val"
    f.Field2 = 20

    return nil
}
```

Do this:
```go
func MakeFoo() (Foo, error) {
    f := Foo {
        Field1: "val",
        Field2: 20,
    }

    return f, nil
}
```

5. Pointer passing performance:
- If a struct is large enough, there are performance improvements from using a pointer to the struct as either an input parameter or a return value.
- The time to pass a pointer into a func is constant for all data sizes, roughly one nanosecond.
- Passing a value into a func takes longer as the data gets larger. It takes about a millisecond once the value gets to be around 10MiB of data.
- Example: https://github.com/learning-go-book/pointer_performance
