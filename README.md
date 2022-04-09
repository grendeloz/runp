# runp

This package help capture information about the run environment.
It can also be used to log that information.

```
package main

import (
    "github.com/grendeloz/runp"
)

var (
    Version = 'myapp v0.1.0dev"
)

func init() {
    runp.SetVersion(Version)
    myRp := runp.NewRunParameters()
}

func main() {
    ...
}
```
