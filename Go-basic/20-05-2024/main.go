package main

import (
	"fmt"
	"sync"
	"time"
)

func demoUnbufferedChannel() {
	var ch = make(chan int)

	go func() {
		ch <- 1
		fmt.Println("sent")
	}()

	time.Sleep(time.Second * 2)
	fmt.Printf("receive value = %d \n", <-ch)

	// Neu khong co sleep thi ko in "Sent"?
	time.Sleep(time.Second * 5)

}

func demoBufferedChannel() {
	// khac nhau dong nay, 2 la so luong (co tui du phong)
	var ch = make(chan int, 2)

	// time.Sleep(time.Second * 2)
	// fmt.Printf("receive value sao = %d \n", <-ch)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		fmt.Println("continue")
	}()

	time.Sleep(time.Second * 2)
	fmt.Printf("receive value = %d \n", <-ch)
	// time.Sleep(time.Second * 2)
	fmt.Printf("receive value = %d \n", <-ch)
	// time.Sleep(time.Second * 2)
	fmt.Printf("receive value = %d \n", <-ch)
	time.Sleep(time.Second * 5)

}

func demoSelectChannelvs1() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)

	ch1 <- 10
	ch2 <- 20

	select {
	// ly do fatal error: all goroutines are asleep - deadlock!?
	// ch2 day vao nhung ko lay ra nen deadlock
	case fromch1 := <-ch1:
		fmt.Println(fromch1)
	case fromch2 := <-ch2:
		fmt.Println(fromch2)
	}

}

func demoSelectChannelvs2() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)

	ch1 <- 10
	ch2 <- 20

	// ly do fatal error: all goroutines are asleep - deadlock!?
	// Vi vong lap vo tan nen khong con ch1 hay ch2 nao duoc day vao tro lai
	for {
		select {
		// ly do fatal error: all goroutines are asleep - deadlock!?
		// ch2 day vao nhung ko lay ra nen deadlock
		case fromch1 := <-ch1:
			fmt.Println(fromch1)
		case fromch2 := <-ch2:
			fmt.Println(fromch2)
		}
	}

}

func demoSelectChannelvs3() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	go func() {
		ch1 <- 10
		ch2 <- 20
	}()

	// ly do fatal error: all goroutines are asleep - deadlock!?
	// tai sao?
	for i := range 2 {
		fmt.Print(i, " ")
		select {
		// ly do fatal error: all goroutines are asleep - deadlock!?
		// ch2 day vao nhung ko lay ra nen deadlock
		case fromch1 := <-ch1:
			fmt.Println(fromch1)
		case fromch2 := <-ch2:
			fmt.Println(fromch2)
		}
	}
}

// Dung bufferedChannel
func demoSelectChannelvs4() {
	var ch1 = make(chan int, 2)
	var ch2 = make(chan int, 2)

	ch1 <- 10
	ch2 <- 20
	// ly do fatal error: all goroutines are asleep - deadlock!?
	// tai sao?
	for i := range 2 {
		fmt.Print(i, " ")
		select {
		// ly do fatal error: all goroutines are asleep - deadlock!?
		// ch2 day vao nhung ko lay ra nen deadlock
		case fromch1 := <-ch1:
			fmt.Println(fromch1)
		case fromch2 := <-ch2:
			fmt.Println(fromch2)
		}
	}
}

// dung khi caai ham thuc thi xong moi duoc mutex?
var mutex sync.Mutex

// control cai so luong go routine thực thi xong
var wg sync.WaitGroup

// vi co nhieu tk cung tuong tac, nen moi xay ra truong hop nhu vay
func demoRaceCondition() {
	var counter = 0

	wg.Add(10000)
	for i := 0; i < 9999; i++ {
		go func() {
			mutex.Lock()

			// giai phong resource
			defer mutex.Unlock()
			counter++

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
	// time.Sleep(time.Second * 5)
}
func main() {
	// var ch = make(chan int, 1)

	// var myFunc = func() {
	// 	ch <- 1
	// 	fmt.Println("sent")
	// }

	// go myFunc()

	// fmt.Printf("receive value = %d \n", <-ch)

	// demoUnbufferedChannel()
	// demoBufferedChannel()
	// demoSelectChannel()
	// demoSelectChannelvs2()
	// demoSelectChannelvs3()
	// demoSelectChannelvs4()
	demoRaceCondition()

}
