package services

import "fmt"

func ProcessPayment(payload *JobPayload) {
	data := payload.Data

	fmt.Println(data)
}
