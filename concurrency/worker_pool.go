package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

type Job struct {
	ID int
	Data interface{}
}

type Result struct {
    JobID int
    Output interface{}
    Error error
}

type WorkerPool struct {
    workerCount int
    jobQueue    chan Job
    resultQueue chan Result
    quit        chan bool
    wg          sync.WaitGroup
}

func NewWorkerPool(workerCount, jobQueueSize int) *WorkerPool {
    return &WorkerPool{
        workerCount: workerCount,
        jobQueue   : make(chan Job, jobQueueSize),
        resultQueue: make(chan Result, jobQueueSize),
        quit       : make(chan bool),
    }
}

func (wp *WorkerPool) Start() {
    for i := 0; i < wp.workerCount; i++ {
        wp.wg.Add(1)
        go wp.worker(i)
    }
}

func (wp *WorkerPool) worker(id int) {
    defer wp.wg.Done()

    for {
        select {
        case job := <-wp.jobQueue:
            fmt.Printf("Worker %d processing job %d\n", id, job.ID)
            result := wp.processJob(job)
            wp.resultQueue <- result

        case <-wp.quit:
            fmt.Printf("Worker %d stopping\n", id)
            return
        }
    }
}

func (wp *WorkerPool) processJob(job Job) Result {
    // Simulate work
    time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
    
    // Process the job (example: square the number)
    if num, ok := job.Data.(int); ok {
        return Result{
            JobID:  job.ID,
            Output: num * num,
        }
    }
    return Result{
        JobID:  job.ID,
        Error:  fmt.Errorf("invalid job data type"),
    }
}

func (wp *WorkerPool) Submit(job Job) {
    wp.jobQueue <- job
}

func (wp *WorkerPool) Results() <- chan Result {
    return wp.resultQueue
}

func (wp *WorkerPool) Stop() {
    close(wp.quit)
    wp.wg.Wait()
    close(wp.jobQueue)
    close(wp.resultQueue)
}

func main() {
    pool := NewWorkerPool(3, 10)
    pool.Start()
    defer pool.Stop()

    go func() {
        for i := 1; i <= 10; i++ {
            job := Job {
                ID: i,
                Data: i * 10,
            }
            pool.Submit(job)
        }
    }()
    // Collect results
    for i := 0; i < 10; i++ {
        result := <-pool.Results()
        if result.Error != nil {
            fmt.Printf("Job %d failed: %v\n", result.JobID, result.Error)
        } else {
            fmt.Printf("Job %d result: %v\n", result.JobID, result.Output)
        }
    }
}
