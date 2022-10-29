# Golang speed

This is a speed golang introduction

## Variables

First file to test is 01_datatypes.go of 01_vars/01_datatypes.go

Can see
```
a := 1
```
This is a short variable declaration, sustitute of normal declaration(To learn more about this i recommend see [this](https://blog.learngoprogramming.com/golang-short-variable-declaration-rules-6df88c881ee))
```
var foo int
foo = 32

var foo int = 32

var foo = 32
```

More lecture of vars
* https://blog.learngoprogramming.com/learn-go-lang-variables-visual-tutorial-and-ebook-9a061d29babe


## Control Structures
Inside of the folder "/02_control_structures" we have everything to work

Go to the file "01_if.go", if is easy to understand on golang, 'if', 'else if' and 'else' \
Now to the file "02_for.go". With for we have all loops, the first for we can see "init"; "conditional"; "loop-expression" \
The second for loop only we can see "conditional", is an while loop... \
The third for loop is an infinite loop, "for(;;)" or "while (true)"




## Concurrency


https://go.dev/tour/concurrency/1


### Goroutines
A goroutine is a lightweight thread managed by th Go runtime

### Channels
A channel gives us a way to "connect" the different concurrent parts of our program
