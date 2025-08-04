package workers

import (
	"context"
	"fmt"
	"log"
	"rinha_de_backend_vinisooo_2025/api/services"
	"rinha_de_backend_vinisooo_2025/config"
	"time"
)

func Worker() {
	redisClient := config.NewRedisClient()
	queueService := services.NewQueueService(redisClient)

	log.Println("Worker started, waiting for jobs...")

	go processQueue(queueService, "payments:processing")
	go processQueue(queueService, "summary_queue")

	select {}
}

func processQueue(queueService *services.QueueService, queueName string) {
	for {
		job, err := queueService.GetJob(context.Background(), queueName)

		if err != nil {
			log.Printf("error getting job from %s: %v", queueName, err)
			time.Sleep(5 * time.Second)
			continue
		}

		log.Printf("Processing job from %s: %s", queueName, job.Type)

		switch job.Type {
		case "process_payment":
			services.ProcessPayment(job)
		case "generate_summary":
			// TODO: Call implementation
			fmt.Println("process summary")
		default:
			log.Printf("Unknown job type: %s", job.Type)
		}
	}
}
