package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	foo(logger)
}
func foo(log *zap.Logger) {
	log.Error("foo")
}
