# runp

This package helps capture information about the run environment.

```
package main

import (
    "github.com/grendeloz/runp"
)

var MyRp runp.RunParameters

func init() {
    runp.SetTool(`myapp`)
    runp.SetVersion(`v0.1.0dev`)
    MyRp = runp.NewRunParameters()
}

func main() {
    ...
}
```
