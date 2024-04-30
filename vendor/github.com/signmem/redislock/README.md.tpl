# redislock

[![Test](https://github.com/signmem/redislock/actions/workflows/test.yml/badge.svg)](https://github.com/signmem/redislock/actions/workflows/test.yml)
[![GoDoc](https://godoc.org/github.com/signmem/redislock?status.png)](http://godoc.org/github.com/signmem/redislock)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Simplified distributed locking implementation using [Redis](http://redis.io/topics/distlock).
For more information, please see examples.

## Examples

```go
import (
  "context"
  "fmt"
  "log"
  "time"

  "github.com/signmem/redislock"
  "github.com/redis/go-redis/v9"
)

func main() {{ "Example" | code }}
```

## Documentation

Full documentation is available on [GoDoc](http://godoc.org/github.com/signmem/redislock)
