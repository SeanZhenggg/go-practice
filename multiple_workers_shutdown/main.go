package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Consumer struct {
	inputChan chan int
	jobsChan  chan int
}

func (c *Consumer) queue(input int) bool {
	select {
	case c.inputChan <- input:
		log.Println("already send input value:", input)
		return true
	default:
		return false
	}
}

func (c *Consumer) process(num, job int) {
	n := getRandomTime()
	log.Printf("Sleeping %d seconds...\n", n)
	time.Sleep(time.Duration(n) * time.Second)
	log.Println("worker:", num, " job value:", job)
}

func (c *Consumer) worker(ctx context.Context, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case job := <-c.jobsChan:
			if ctx.Err() != nil {
				log.Printf("get next job %d and close the worker %d", job, num)
				return
			}
			c.process(num, job)
		case <-ctx.Done():
			log.Printf("close the worker %d", num)
			return
		}
	}
}

func (c *Consumer) startConsumer(ctx context.Context) {
	for {
		select {
		case job := <-c.inputChan:
			if ctx.Err() != nil {
				close(c.jobsChan)
			}
			c.jobsChan <- job
		case <-ctx.Done():
			close(c.jobsChan)
			return
		}
	}
}

const poolSize = 2

func withContextFunc(ctx context.Context, f func()) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(c)

		select {
		case <-ctx.Done():
		case <-c:
			cancel()
			f()
		}
	}()

	return ctx
}

func main() {
	finished := make(chan bool)
	wg := &sync.WaitGroup{}

	ctx := withContextFunc(context.Background(), func() {
		log.Println("cancel from ctrl+c event")
		wg.Wait()
		close(finished)
	})
	// create the consumer
	consumer := Consumer{
		inputChan: make(chan int, 10),
		jobsChan:  make(chan int, poolSize),
	}

	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go consumer.worker(ctx, i, wg)
	}

	go consumer.startConsumer(ctx)

	consumer.queue(1)
	consumer.queue(2)
	consumer.queue(3)
	consumer.queue(4)
	consumer.queue(5)

	<-finished
	//time.Sleep(2 * time.Second)
}

func getRandomTime() int {
	return rand.Intn(10)
}
