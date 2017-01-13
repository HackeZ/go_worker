# go_worker

## Install

`$ go get github.com/hackez/go_worker`

## Usage

```golang
package main

import (
    promise "github.com/hackez/go_worker"
)

func main() {
    promise := go_promise.NewPromise()
	params := []interface{}{1, 2, 3, 4}

	promise.Add(mul, params[:2])
	promise.Add(add, params[2:])
	promise.Add(wait, params[:2])

	promise.Run()
	
}

func add(nums ...interface{}) {
	var sum int
	fmt.Println("nums:", nums)
	for _, num := range nums {
		switch num.(type) {
		case int:
			sum += num.(int)
		default:
			fmt.Println("unsupport number type")
			continue
		}
	}
	fmt.Println("sum of add: ", sum)
	return
}

func mul(nums ...interface{}) {
	var sum int = 1
	fmt.Println("nums:", nums)
	for _, num := range nums {
		switch num.(type) {
		case int:
			sum *= num.(int)
		default:
			fmt.Println("unsupport number type")
			continue
		}
	}
	fmt.Println("sum of mul: ", sum)
	return
}

func wait(timeSec ...interface{}) {
	for _, t := range timeSec {
		switch t.(type) {
		case int:
			time.Sleep(2 * time.Second)
		default:
			fmt.Println("unsupport number type")
			continue
		}
	}
	return
}
```