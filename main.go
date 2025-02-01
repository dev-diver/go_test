package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	go func() { fmt.Println("다른 고루틴") }()
	fmt.Println("STOP")
	<-ctx.Done()
	fmt.Println("그리고 시간은 움직이기 시작한다.")
}
