// logrusctx provides convenience functions for passing a logrus FieldLogger through context values.
package logrusctx

import (
	"context"

	"github.com/sirupsen/logrus"
)

type contextKey int

const key contextKey = 0

// Get returns a logrus FieldLogger from the context provided, or defaults to StandardLogger if not provided.
//
// Example:
//  func foo(ctx context.Context, f int) {
//    logger := logrusctx.Get(ctx)
//    logger.WithField("f", f).Info("foo")
//  }
func Get(ctx context.Context) logrus.FieldLogger {
	value := ctx.Value(key)
	if logger, _ := value.(*logrus.Entry); logger != nil {
		return logger
	}

	return logrus.StandardLogger()
}

// WithLogger creates a new context with a logger value attached which can be retrieved using Get.
//
// Example:
//  func main() {
//    logger := &logrus.Logger{...}
//    foo(logrusctx.WithLogger(context.Background(), logger), 1)
//  }
func WithLogger(ctx context.Context, logger logrus.FieldLogger) context.Context {
	// explicit copy to prevent receiver from editing caller's Entry
	entry := logger.WithFields(logrus.Fields{})

	return context.WithValue(ctx, key, entry)
}

// WithField is a convenience method for use cases where a field is added to the context without the caller needing the logger itself.
//
// Example:
//
//  func bar(ctx context.Context) {
//    foo(logrusctx.WithField("bar", "2"))
//  }
func WithField(ctx context.Context, key string, value interface{}) context.Context {
	return WithLogger(ctx, Get(ctx).WithField(key, value))
}

// WithField is a convenience method for use cases where fields are added to the context without the caller needing the logger itself.
//
// Example:
//
//  func bar(ctx context.Context) {
//    foo(logrusctx.WithFields(logrus.Fields{"a": 1, "b": 2}))
//  }
func WithFields(ctx context.Context, fields logrus.Fields) context.Context {
	return WithLogger(ctx, Get(ctx).WithFields(fields))
}
