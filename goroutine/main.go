package main

import (
	"fmt"
	"sync"
	"time"
)

func goroutine(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func syncgoroutine(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutinegrammer() {
	// method の頭に go をつければメソッドが並列で処理される
	// goroutine はあくまで並列処理なので、goroutine 以外の処理が終了した段階でgorutineの処理は終わる
	go goroutine("world")
	normal("hello")
}

func wggrammer() {
	var wg sync.WaitGroup
	// goroutine の処理があることを通知する。上記メソッドのように、goroutineの処理が途中で終了することがないようにすることができる
	wg.Add(1)
	go syncgoroutine("aaa", &wg)
	normal("bbb")
	wg.Wait()
	// wg での待機処理したいをしたい場合は
	// 1. wgのvarを定義
	// 2. wg.Add で何個の待機処理があるか定義
	// 3. goroutine の処理に引数として渡す
	// 4. goroutine 川の処理で wg.Doneを実行。処理が終了したことを通知する
	// 5. wg.Wait でwg に登録していた処理が全て終了するのを待つ
}

func chanelgorutine(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// c（channel）に値を入れる処理
	c <- sum
}

// channel: goroutine と呼び出し元のデータのやり取りをする際に利用できる
func channelgrammer() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go chanelgorutine(s, c)
	// channel int型 -> int型
	// channel int型はあくまでgorutineの中での値のやり取りを行うためのデータ型
	// 本来goroutineの処理は非同期なので勝手に処理が終了するが、以下の処理を入れることで、goroutineの処理が終わるまで待機する
	x := <-c
	fmt.Println(x)
	// channel型のint
	fmt.Printf("%T", c)
	// int型
	fmt.Printf("%T", x)

}

func bufferchannelgrammer() {
	// channel のbufferが2つであることを意味する
	// 2つまでchannelを入れることができるので、3つ目を入れるとエラーが出る
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))
	// channelが開いてる状態で中身を参照しようとするとdeadlockが発生してしまうので
	close(ch)
	for c := range ch {
		fmt.Println(c)
	}
}
func channelgorutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func channelgrammer2() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go channelgorutine2(s, c)
	// channelで値が送信されるたびに受信する
	for i := range c {
		fmt.Println(i)
	}
}

func producer(ch chan int, i int) {
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		func() {
			defer wg.Done()
			fmt.Println("process", i*1000)
		}()
	}
	fmt.Println("############")
}

func producerconsumergrammer() {
	var wg sync.WaitGroup
	ch := make(chan int)

	//producer: データを集める
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	// Consumer: 集めたデータを加工する
	go consumer(ch, &wg)
	wg.Wait()
	// close をしないとconsumer内のchのrangeの参照が終わらず、chに新しい値が入るのを待ち続ける
	close(ch)
	// goroutineの処理は一般処理が終わると閉じるので、本来はそれ以降の処理が走らない。
	// timeSleepを入れることで一般処理が走り続けるので、その間にgoroutineの処理が継続して走り続ける
	time.Sleep(2 * time.Second)
	fmt.Println("DONE")
}

func fofiproducer(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

func fofimulti2(first <-chan int, second chan<- int) {
	defer close(second)
	for i := range first {
		second <- i * 2
	}
}

func fofimulti4(second chan int, third chan int) {
	defer close(third)
	for i := range second {
		third <- i * 4
	}
}

// data flow: main -> producer -> second stage -> third stage -> main
func funoutfuningrammer() {
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go fofiproducer(first)
	go fofimulti2(first, second)
	go fofimulti4(second, third)
	for result := range third {
		fmt.Println(result)
	}
}

func selectgoroutine1(ch chan string) {
	// for loop でネットワークから来るパケットを撮り続ける
	for {
		ch <- "packet from 1"
		time.Sleep(3 * time.Second)
	}
}

func selectgoroutine2(ch chan string) {
	for {
		ch <- "packet from 2"
		time.Sleep(5 * time.Second)
	}
}

// goroutine の処理を同時に受信することができる
func goroutineselectgrammer() {
	c1 := make(chan string)
	c2 := make(chan string)
	go selectgoroutine1(c1)
	go selectgoroutine2(c2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

// for loop で監視中の処理を中断する手段の一つとして、breakがある
func breakselectgrammer() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

OuterLoop2:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			break OuterLoop2
		default:
			fmt.Println("     .")
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Println("##############")
}

func main() {
	breakselectgrammer()
}
