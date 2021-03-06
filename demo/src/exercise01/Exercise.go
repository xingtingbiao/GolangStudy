package main

import (
	"fmt"
	"time"
)

func main() {
	//sum(1, 2, 5, 7)
	// 数组或者切片类型直接作为可变参数需要加...
	// sum([]int{1, 3, 5, 7, 9}...)
	//testIntSeq()
	//fmt.Println(fact(7))
	//testWorker()
	//testPingPang()
	//testSelect()
	//testTimeOut()
	//testSelect2()
	//testCloseChan()
	//testChanRange()
	//testTimer()
	//testTicker()
	//testChanRange2()
}

// 可变参数的函数
func sum(nubs ...int) {
	total := 0
	for _, n := range nubs {
		total += n
	}
	fmt.Println(total)
}

// 闭包(匿名函数) 闭包就是能够读取其他函数内部变量的函数。
// 例如在javascript中，只有函数内部的子函数才能读取局部变量，所以闭包可以理解成“定义在一个函数内部的函数“。在本质上，闭包是将函数内部和函数外部连接起来的桥梁
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func testIntSeq() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInts := intSeq()
	fmt.Println(nextInts())
	fmt.Println(nextInts())
}

// 阶乘的递归练习
func fact(n int) int {
	if n == 2 {
		return 2
	}
	return n * fact(n-1)
}

// chan的同步实例
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func testWorker() {
	//这里设不设置长度为1对效果不影响, 不设置两边都阻塞, 设置的话, 下面阻塞, 上面不阻塞
	done := make(chan bool, 1)
	// done := make(chan bool)
	go worker(done)
	<-done
}

// chan的路线实例, 具有传导性
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	pongs <- <-pings
}

func testPingPang() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message.")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

// 每个通道将在一段时间后开始接收值，以模拟阻塞在并发goroutines中执行的RPC操作。
// 我们将使用select同时等待这两个值，在每个值到达时打印它们。
func testSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(time.Second)
		ch1 <- "one"
		// close(ch1)
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "two"
		// close(ch2)
	}()

	// 注意这里的ok是判断通道是否关闭掉了
	for i := 0; i < 3; i++ {
		select {
		case msg1, ok := <-ch1:
			fmt.Println(msg1, ok)
		case msg2, ok := <-ch2:
			fmt.Println(msg2, ok)
		}
	}
}

// 超时练习
func testTimeOut() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("timeout 1")
	}

	go func() {
		time.Sleep(time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}

// default实现非阻塞通道
func testSelect2() {
	messages := make(chan string)
	signals := make(chan bool)
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

// 关闭通道练习
func testCloseChan() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	// 这里的ok只有在通道关闭的时候才会为FALSE
	go func() {
		for true {
			j, ok := <-jobs
			if ok {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("sent job", i)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}

// 通道范围实例
func testChanRange() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for e := range queue {
		fmt.Println(e)
	}
	ints := make(chan int)
	close(ints)
	// 注意这里依然可以执行, 但是返回空值
	fmt.Println(<-ints)
}

// 定时器实例
func testTimer() {
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop := timer2.Stop()
	if stop {
		fmt.Println("Timer 2 stopped")
	}
}

// 断续器实例 tickers是用于定期做一些事情 周期性执行直到停止
func testTicker() {
	ticker := time.NewTicker(time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(time.Second * 4)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

// 测试通过range 一个通道进行循环, 是否依然阻塞循环
func testChanRange2() {
	c1 := make(chan int, 1)
	done := make(chan bool)
	go func() {
		for i := 0; i < 4; i++ {
			time.Sleep(time.Second * 2)
			c1 <- i
		}
		close(c1)
		time.Sleep(time.Millisecond)
		done <- true
	}()

	go func() {
		for v := range c1 {
			fmt.Println("current is", v)
		}
	}()
	<-done
}
