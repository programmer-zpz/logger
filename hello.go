// Copyright 2021 The Cloudbases Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

//go:build ignore
// +build ignore

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

	ctx = ctxutil.SetRequestId(ctx, "")
	logger.Infof(ctx, "hello context3")
}
