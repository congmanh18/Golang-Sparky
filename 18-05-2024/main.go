package main

import (
	"fmt"
	"time"
)

func downloadFile() {
	fmt.Println("manh 1")
}

func basic() {
	go downloadFile()

	go func() {
		fmt.Println("manh 2")
	}()

	var f = func() {
		fmt.Println("manh 3")
	}
	go f()

	// đợi go downloadFile trước mới kết thuc go main
	time.Sleep(time.Second)
}

func basic1() {
	for i := 0; i < 1_000_000; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}

func democomunicationGoRoutine() {
	ch := make(chan int)

	go func() {
		ch <- 1
		fmt.Println("sent")
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("receiving")
	var result = <-ch
	fmt.Println(result)
}

// mặc dù đã tạm ngưng 2 giây nhưng
// đợi đến khi lấy được ch ra thì "sent" mới được in ra

type User struct {
	ID   string
	Name string
}

func demoPassStructToChannel() {
	var ch = make(chan User)
	var getUser = func() {
		fmt.Println("start running long task #1 ...")
		time.Sleep(time.Second * 5)
		ch <- User{
			ID:   "1",
			Name: "Manh",
		}
	}

	go getUser()

	var user = <-ch
	fmt.Println(user)
}

func demoDeadlock() {
	var ch = make(chan int)

	go func() {
		ch <- 1
	}()

	<-ch
	<-ch // yeu don phuong, da co nguoi yeu roi
	fmt.Println("done")
}

func main() {
	demoPassStructToChannel()
	// time.Sleep(time.Second * 2) //khong dung vi khi ch duoc ban ra thi da block va doi o do roi nen ko can

	// democomunicationGoRoutine()
	// basic1()
}
