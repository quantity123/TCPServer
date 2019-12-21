package server

import "sync"

type IServer interface {
	Launch(wg *sync.WaitGroup)
}