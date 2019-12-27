package logrusctx

import (
  "context"
  
  "github.com/sirupsen/logrus"
)

type contextKey int

const key contextKey = 0

// Logger returns a logrus FieldLogger from the context provided, or defaults to StandardLogger.
func Logger(ctx context.Context) logrus.FieldLogger {
  value := ctx.Value(key)
  if logger, _ := value.(*logrus.Entry); logger != nil {
    return logger
  }
  
  return logrus.StandardLogger()
}

// WithLogger creates a new context with a logger value attached which can be retrieved using Entry.
func WithLogger(ctx context.Context, logger logrus.FieldLogger) context.Context {
  // explicit copy to prevent receiver from editing caller's Entry
  entry := logger.WithFields(logrus.Fields{})

  return context.WithValue(ctx, key, entry)
}

// WithField is a convenience method for use cases where a field is added to the context without the caller needing the logger itself.
func WithField(ctx context.Context, key string, value interface{}) context.Context {
  return WithLogger(ctx, Logger(ctx).WithField(key, value))
}

// WithField is a convenience method for use cases where fields are added to the context without the caller needing the logger itself.
func WithFields(ctx context.Context, fields logrus.Fields) context.Context {
  return WithLogger(ctx, Logger(ctx).WithFields(fields))
}
