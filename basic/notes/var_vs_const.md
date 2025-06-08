# Differ var and const keyword

`var` is used to declare variables - can be changed after the declaration

`const` is used to declare constants - can *not* be changed after the declaration

example:
```Golang
package main

import "fmt"

func main() {
    // Using var
    var variable1 int = 5
    variable1 = 10  // Valid, the value of variable1 can be changed

    // Using const
    const constant1 int = 5
    // constant1 = 10  // Invalid, constants cannot be reassigned

    fmt.Println(variable1, constant1)
}
```
