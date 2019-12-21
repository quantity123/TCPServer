package main

import (
	"fmt"
	"qtt/gameServer/server"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("qttGameServer Begin")
	defer fmt.Println("qttGameServer End")
	sev := server.NewServer()
	wg.Add(1)
	go sev.Launch(&wg)
	wg.Wait()
}