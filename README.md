# logger for Cloudbases

[![Build Status](https://travis-ci.org/openpitrix/logger.svg)](https://travis-ci.org/openpitrix/logger)
[![Go Report Card](https://goreportcard.com/badge/openpitrix.io/logger)](https://goreportcard.com/report/openpitrix.io/logger)
[![GoDoc](https://godoc.org/openpitrix.io/logger?status.svg)](https://godoc.org/openpitrix.io/logger)
[![License](http://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/openpitrix/logger/blob/master/LICENSE)

## Install

1. Install Go1.11+
1. set `GO111MODULE=on`
1. `go get cloudbases.io/logger`
1. `go run hello.go`

## Example

```go
package main

import (
	"bytes"
	"context"
	"fmt"

	"cloudbases.io/logger"
	"cloudbases.io/logger/ctxutil"
)

func main() {
	logger.Infof(nil, "hello1 cloudbases")

	logger.HideCallstack()
	logger.Infof(nil, "hello2 cloudbases")

	logger.ShowCallstack()
	logger.Infof(nil, "hello3 cloudbases")

	{
		var buf bytes.Buffer
		logger := logger.New().SetOutput(&buf)

		logger.Infof(nil, "hello4 cloudbases")
		logger.Infof(nil, "hello5 cloudbases\nlogger")
		logger.Infof(nil, "")

		fmt.Print(buf.String())
	}

	ctx := context.Background()
	ctx = ctxutil.SetRequestId(ctx, "req-id-001")

	logger.Infof(ctx, "hello context1")

	ctx = ctxutil.SetMessageId(ctx, "msg-001", "msg-002")
	logger.Infof(ctx, "hello context2")

	ctx = ctxutil.ClearRequestId(ctx)
	logger.Infof(ctx, "hello context3")
}
```

output:

```
2018-10-20 07:43:31.70066 -INFO- hello1 cloudbases (hello.go:19)
2018-10-20 07:43:31.70085 -INFO- hello2 cloudbases
2018-10-20 07:43:31.70086 -INFO- hello3 cloudbases (hello.go:25)
2018-10-20 07:43:31.70089 -INFO- hello4 cloudbases (hello.go:31)
2018-10-20 07:43:31.70091 -INFO- hello5 cloudbases\nlogger (hello.go:32)
2018-10-20 07:43:31.70093 -INFO-  (hello.go:33)
2018-10-20 07:43:31.70096 -INFO- hello context1 (hello.go:41){@req-id-001}
2018-10-20 07:43:31.70099 -INFO- hello context2 (hello.go:44){msg-001|msg-002@req-id-001}
2018-10-20 07:43:31.70101 -INFO- hello context3 (hello.go:47){msg-001|msg-002@}
```
