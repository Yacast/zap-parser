[![Documentation](https://godoc.org/github.com/Yacast/zap-parser?status.svg)](https://godoc.org/github.com/Yacast/zap-parser)

# zap-parser

A golang parser for [Uber's zap](https://github.com/uber-go/zap) logger json output.

## Quick Start

```go
import (
    // ...
    "fmt"

	"github.com/Yacast/zap-parser"
)

p, err := zapparser.FromFile("/path/to/logs"))
if err != nil {
	panic(err)
}

p.OnError(func(err error) {
    fmt.Println(err)
})

p.OnEntry(func(e *zapparser.Entry) {
    fmt.Println(e.Message)
})

p.Start()
fmt.Println("Done parsing...")
```
