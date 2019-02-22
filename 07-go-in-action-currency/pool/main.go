package main

import (
	"io"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/xiaoheigou/how-to-code/07-go-in-action-currency/pool/pool"
)

const (
	maxGoroutines   = 25
	pooledResources = 2
)

type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	log.Println("close:connection", dbConn.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("create new connetcion", id)
	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup

	wg.Add(maxGoroutines)

	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}
	wg.Wait()

}

func performQueries(querry int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Print(err)
		return
	}
	defer p.Release(conn)

	time.Sleep(time.Duration(time.Second))
	log.Println("正在查询")

}
