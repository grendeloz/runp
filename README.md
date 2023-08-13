# runp

Package runp captures basic information about the run environment.

It captures the execution start time so it should be initialised as
early as possible - ideally in an `init()` or at the very start of
`main()`.

Once you have initialised your `RunParameters` object, all of the work
is done and you can directly access the fields all of which are
public and shown below.

```
type RunParameters struct {
        Tool      string
        Version   string
        StartTime time.Time
        Args      []string
        UserId    int
        UserName  string
        GroupId   int
        GroupName string
        HostName  string
}
```

There are 2 very simple functions - `Log()` and `LogFinish()` which
are intended to log run parameters at the start and finish of an
application's execution as shown in the example below.
These are very basic functions so you may wish to create your own 
more sophisticated functions to start and finish logging.

```
package main

import (
	"github.com/grendeloz/runp"
)

var MyRp runp.RunParameters

func init() {
     runp.SetTool("myapp")
     runp.SetVersion("v0.1.0-dev")
     MyRp = runp.NewRunParameters()
}

func main() {
    MyRp.Log()
    // Do some clever stuff
    MyRp.LogFinish()
}
```

Compiling and running this example outputs:

```
2023/08/13 10:00:46 Tool: myapp v0.1.0-dev
2023/08/13 10:00:46 Cmdline: [./myapp]
2023/08/13 10:00:46 Host: my-m1
2023/08/13 10:00:46 User: 1001 (grendeloz)
2023/08/13 10:00:46 Group: 20 (staff)
2023/08/13 10:00:46 Elapsed time: 139.5µs
```

