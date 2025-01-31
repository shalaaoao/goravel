package tools

import (
	"fmt"
	"time"
)

func ExecTime(closure func() interface{}, name string) interface{} {
	start := time.Now()
	res := closure()
	duration := time.Since(start)
	fmt.Printf("%s: %vs\n", name, duration.Seconds())
	return res
}
