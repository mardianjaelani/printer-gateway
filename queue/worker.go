package queue

import (
	"log"
	"print-gateway/printer"
)

func (q *Queue) StartWorker() {

	go func() {

		for {

			job := q.Pop()

			log.Println("Printing :", job.ID)

			err := printer.RawPrint(
				job.Printer,
				job.Data,
				job.Copies,
			)

			if err != nil {

				log.Println(err)

			}

		}

	}()

}
