package util

import (
	"fmt"
	"time"
)

// FunctionHandlerOpts opts for FunctionHandler
type FunctionHandlerOpts struct {
	Name   string
	Log    func(format string, v ...interface{})
	LogErr func(format string, v ...interface{})
}

// FunctionHandler safely executes cloud function
func FunctionHandler(opts FunctionHandlerOpts, g func() error) error {
	var panicErr error
	err := functionHandlerNotifyPanic(
		opts,
		func(newPanicErr error) {
			panicErr = newPanicErr
		},
		g,
	)
	if err != nil {
		return err
	}
	if panicErr != nil {
		return panicErr
	}
	return nil
}

func functionHandlerNotifyPanic(opts FunctionHandlerOpts, notifyPanic func(err error), g func() error) error {
	sTime := time.Now()
	defer func() {
		if x := recover(); x != nil {
			opts.LogErr("Function %s finished in %v with panic: %v", opts.Name, getDuration(sTime), x)
			notifyPanic(fmt.Errorf("Panic %v", x))
		}
	}()
	err := g()
	if err != nil {
		opts.LogErr("Function %s finished in %v with error: %v", opts.Name, getDuration(sTime), err)
		return err
	}
	opts.Log("Function %s finished in %v with success", opts.Name, getDuration(sTime))
	return nil
}

func getDuration(sTime time.Time) time.Duration {
	return time.Now().Sub(sTime)
}
