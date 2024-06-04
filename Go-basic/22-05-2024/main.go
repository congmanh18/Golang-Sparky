package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func downLoadFile() {
	// time.Sleep(time.Second * 5)
	// println("Download file")

	select {
	case <-time.After(time.Second * 2):
		println("Download file Done...")

	}
}

func demoTimeOutContext() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()
	var selectUserById = func() {
		fmt.Println("select * from users where id = 1")
	}
	go func() {
		selectUserById()
		select {
		case <-time.After(time.Microsecond * 500):
			cancel()
		}
	}()
	// go downLoadFile()
	select {
	case <-ctx.Done():
		fmt.Println("warning: query too slow")
	}
}

var errDivision = errors.New("division by error")

func division(a, b int) (int, error) {

	if b == 0 {
		return 0, errDivision
	}

	return a / b, nil
}

func demoCancelCause() {
	ctx, cancel := context.WithCancelCause(context.Background())

	go func() {
		var result, err = division(1, 0)
		if err != nil {
			cancel(err)
			return
		}
		fmt.Println(result, "context canceled")
	}()

	<-ctx.Done()

}

func main() {
	demoTimeOutContext()
	// demoCancelCause()
}
