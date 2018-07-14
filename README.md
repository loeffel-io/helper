# helper

[![Go Report Card](https://goreportcard.com/badge/github.com/loeffel-io/helper)](https://goreportcard.com/report/github.com/loeffel-io/helper)

Simple helper package 

### Sentry support

This package supports [sentry.io](https://sentry.io) real time error reporting.
More informations: [Sentry golang docs](https://docs.sentry.io/clients/go)

```go
package main

import "github.com/getsentry/raven-go"

func init() {
    // Setup sentry
    raven.SetDSN("https://<key>:<secret>@sentry.io/<project>")
    
    // Setup cronjobs ...
}
```