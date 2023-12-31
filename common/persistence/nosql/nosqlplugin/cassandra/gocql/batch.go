// Copyright (c) 2017-2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package gocql

import (
	"context"
	"fmt"

	"github.com/gocql/gocql"
)

var _ Batch = (*batch)(nil)

type (
	batch struct {
		*gocql.Batch
	}
)

// Definition of all BatchTypes
const (
	LoggedBatch BatchType = iota
	UnloggedBatch
	CounterBatch
)

func newBatch(
	gocqlBatch *gocql.Batch,
) *batch {
	return &batch{
		Batch: gocqlBatch,
	}
}

func (b *batch) WithContext(ctx context.Context) Batch {
	b2 := b.Batch.WithContext(ctx)
	if b2 == nil {
		return nil
	}
	return newBatch(b2)
}

func (b *batch) WithTimestamp(timestamp int64) Batch {
	b.Batch.WithTimestamp(timestamp)
	return b
}

func (b *batch) Consistency(c Consistency) Batch {
	b.Batch.SetConsistency(mustConvertConsistency(c))
	return b
}

func mustConvertBatchType(batchType BatchType) gocql.BatchType {
	switch batchType {
	case LoggedBatch:
		return gocql.LoggedBatch
	case UnloggedBatch:
		return gocql.UnloggedBatch
	case CounterBatch:
		return gocql.CounterBatch
	default:
		panic(fmt.Sprintf("Unknown gocql BatchType: %v", batchType))
	}
}
