package hbgo

import "runtime"

func Trace() {
	if r := recover(); r != nil {
		buf := make([]byte, 1024)
		l := runtime.Stack(buf, false)
		LogInfo("%s, %s", r, string(buf[:l]))
	}
}
