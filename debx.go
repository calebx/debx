package debx

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func Debx(s string) string {
	wg := sync.WaitGroup{}
	wg.Add(3000)
	x := ""
	for i := 0; i < 3000; i++ {
		go func() {
			defer wg.Done()

			// generate a long random string
			s := randStringBytes(100000)
			x = fmt.Sprintf("Hello from Go! %d %s", i, s)
			x = fmt.Sprintf("Hello from Go! %s %d", s, i)
			if len(x) > 100000 {
				x = x[:100000]
			}

			if i%1000 == 0 {
				runtime.GC()
				runtime.Gosched()
			}
		}()
	}
	wg.Wait()
	return x[:100]
}
