# Golang CLI Module

###### By - Agung

I have some problems in learning to create applications with golang, especially in the CLI (Command Line Interface) part, and I try to create a library that helps me shorten the coding.

```go
// return int
num := cligz.IntegerResult("Insert Number msg", "it's not number")

// return string
abc := cligz.StringResult("Message String")


// return int
abc := cligz.Menu("Enter your choice", []string {
    "Customer","Service", "Result", "Exit",
})

// return slice
abc := cligz.Form([]string {
    "Nama","Email", "Umur",
})
```
