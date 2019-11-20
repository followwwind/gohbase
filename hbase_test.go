package gohbase

import (
	"context"
	"fmt"
	"github.com/followwwind/gohbase/hrpc"
	"testing"
	"time"
)

var zClient = NewClient("localhost")

func TestPutRow(t *testing.T) {
	values := make(map[string]string)
	values["name"] = "wind"
	val := make(map[string]map[string]string)
	val["info"] = values
	value := make(map[string]map[string][]byte)
	for k, v := range val {
		cols := make(map[string][]byte)
		for ck, cv := range v {
			cols[ck] = []byte(cv)
		}
		value[k] = cols
	}
	putRequest, err := hrpc.NewPutStr(context.Background(), "test", "333", value)
	if err != nil {
		fmt.Printf("put col error is %s", err)
	}

	rsp, err := zClient.Put(putRequest)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	fmt.Println(rsp)
}

func TestContext(t *testing.T) {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	c := context.WithValue(ctx, "hello", "world")
	ch := make(chan int)
	go print(c, ch)
	i := <-ch
	fmt.Println(i)
	if i == 1 {
		close(ch)
		return
	}

	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err()) // prints "context deadline exceeded"
			fmt.Println(ctx.Deadline())
			return
		}
	}
}

func print(ctx context.Context, ch chan int) {
	for {
		//fmt.Println(time.Now())
		fmt.Println(ctx.Value("hello"))
		select {
		case <-ctx.Done():
			fmt.Println("hello world")
			//fmt.Println(ctx.Err()) // prints "context deadline exceeded"
			//fmt.Println(ctx.Deadline())
			ch <- 1
			return
		case <-time.After(time.Second):
		}
	}
}
