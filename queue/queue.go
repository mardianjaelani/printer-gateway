package queue

import "print-gateway/models"

type Queue struct {
	jobs chan models.PrintJob
}

func New(size int) *Queue {

	return &Queue{
		jobs: make(chan models.PrintJob, size),
	}

}

func (q *Queue) Push(job models.PrintJob) {

	q.jobs <- job

}

func (q *Queue) Pop() models.PrintJob {

	return <-q.jobs

}
