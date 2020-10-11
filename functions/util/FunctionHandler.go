package util

import (
	"time"
)

// FunctionHandlerOpts opts for FunctionHandler
type FunctionHandlerOpts struct {
	Name       string
	LogErrorFn func(format string, v ...interface{})
	LogPanicFn func(format string, v ...interface{})
	LogDoneFn  func(format string, v ...interface{})
}

// FunctionHandler safely executes cloud function
func FunctionHandler(opts FunctionHandlerOpts, g func() error) error {
	sTime := time.Now()
	defer func() {
		if x := recover(); x != nil {
			opts.LogPanicFn("Function %s finished in %v with panic: %v", opts.Name, getDuration(sTime), x)
		}
	}()
	err := g()
	if err != nil {
		opts.LogErrorFn("Function %s finished in %v with error: %v", opts.Name, getDuration(sTime), err)
		return err
	}
	opts.LogDoneFn("Function %s finished in %v with success", opts.Name, getDuration(sTime))
	return nil
}

func getDuration(sTime time.Time) time.Duration {
	return time.Now().Sub(sTime)
}
