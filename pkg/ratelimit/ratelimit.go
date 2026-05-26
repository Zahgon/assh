package ratelimit

// Package ratelimit based on http://hustcat.github.io/rate-limit-example-in-go/

import (
	"io"

	"golang.org/x/time/rate"
)

type reader struct {
	r       io.Reader
	limiter *rate.Limiter
}

// NewReader returns a reader that is rate limited by
// the given token bucket. Each token in the bucket
// represents one byte.
func NewReader(r io.Reader, l *rate.Limiter) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

func (r *reader) Read(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// fmt.Printf("Read %d bytes, delay %d\n", n, delay)

type writer struct {
	w       io.Writer
	limiter *rate.Limiter
}

// NewWriter returns a writer that is rate limited by
// the given token bucket. Each token in the bucket
// represents one byte.
func NewWriter(w io.Writer, l *rate.Limiter) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

func (w *writer) Write(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// fmt.Printf("Write %d bytes, delay %d\n", n, delay)
