package go_promise

import (
	"sync"
	"fmt"
	"time"
)

type Promise struct {
    handleFunc []func(...interface{})
	funcParams [][]interface{}
}

func NewPromise() (p Promise) {
	p = Promise{}
	return
}

func (p *Promise) Add(fn func(...interface{}), params []interface{}) {
	p.handleFunc = append(p.handleFunc, fn)
	p.funcParams = append(p.funcParams, params)

}

func (p *Promise) Run() {
	var wg sync.WaitGroup

	wg.Add(len(p.funcParams))
	start := time.Now()

	for index, fn := range p.handleFunc {
		go func(fn func(...interface{}), i int) {
			fn(p.funcParams[i]...)
			wg.Done()
		}(fn, index)
	}

	wg.Wait()
	end := time.Now()
	fmt.Println("total:", end.Sub(start).Nanoseconds(), "ns")
}