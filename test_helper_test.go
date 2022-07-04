// Copyright 2021 The Cloudbases Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package logger

import (
	"bytes"
	"context"
	"io/ioutil"
	"testing"

	"cloudbases.io/logger/ctxutil"
)

type tLogger struct {
	t   *testing.T
	buf bytes.Buffer
	*Logger
}

func tNewLogger(t *testing.T) *tLogger {
	p := &tLogger{Logger: New()}
	p.Logger.SetOutput(&p.buf)
	return p
}

func (p *tLogger) ReadAll() string {
	data, err := ioutil.ReadAll(&p.buf)
	if err != nil {
		p.t.Fatal(err)
	}
	return string(data)
}
func (p *tLogger) ReadAllMessage() []logMessage {
	return readLogs(p.ReadAll())
}

func tNewCtxWithMessageId(reqId string, messageId ...string) context.Context {
	ctx := context.Background()
	ctx = ctxutil.SetRequestId(ctx, reqId)
	ctx = ctxutil.SetMessageId(ctx, messageId...)
	return ctx
}
