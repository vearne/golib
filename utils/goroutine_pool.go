package utils

import (
	"fmt"
	"sync"
)

const (
	SIZE int = 50
)

// 一个简易的协程池实现
type JobFunc func(str string) bool

type GPool struct {
	sync.Mutex

	// 任务队列
	JobChan chan string
	// 结果队列
	ResultChan chan bool
	// 协程池的大小
	Size int
	// 已经完成的任务量
	FinishCount int
	// 目标任务量
	TargetCount int
}

func NewGPool(size int) *GPool {
	pool := GPool{}
	pool.JobChan = make(chan string, SIZE)
	pool.ResultChan = make(chan bool, SIZE)
	pool.Size = size
	return &pool
}

func (p *GPool) ApplyAsync(f JobFunc, slice []string) <-chan bool {

	p.TargetCount = len(slice)
	// Producer
	go p.Produce(slice)
	// consumer
	for i := 0; i < p.Size; i++ {
		go p.Consume(f)
	}

	return p.ResultChan
}

func (p *GPool) Produce(slice []string) {
	for _, key := range slice {
		p.JobChan <- key
	}
	close(p.JobChan)
}

func (p *GPool) Consume(f JobFunc) {
	for job := range p.JobChan {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Errorf("execute job error, %v", err)
				p.ResultChan <- false
				p.FinishOne()
			}
		}()
		p.ResultChan <- f(job)
		p.FinishOne()
	}
	p.TryClose()
}

// 记录完成了一个任务
func (p *GPool) FinishOne() {
	p.Lock()
	p.FinishCount++
	p.Unlock()
}

// 关闭结果Channel
func (p *GPool) TryClose() {
	p.Lock()
	if p.FinishCount == p.TargetCount {
		close(p.ResultChan)
	}
	p.Unlock()
}
