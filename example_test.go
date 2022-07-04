// Copyright 2021 The Cloudbases Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package logger_test

import (
	"context"

	"cloudbases.io/logger"
	"cloudbases.io/logger/ctxutil"
)

func Example() {
	logger.Infof(nil, "hello cloudbases")
}

func Example_withContext() {
	ctx := context.Background()
	ctx = ctxutil.SetRequestId(ctx, "req-id-001")
	ctx = ctxutil.SetMessageId(ctx, "msg-001", "msg-002")

	logger.Infof(ctx, "hello cloudbases")
}
