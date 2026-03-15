package main

import (
	"fmt"
	"os"
)

type BluePrintStatus string

const (
	BluePrintStatusActive   BluePrintStatus = "active"
	BluePrintStatusInactive BluePrintStatus = "inactive"
	BluePrintStatusPending  BluePrintStatus = "pending"
)

type BluePrintCartAbandonmentDetail struct {
	DelaySeconds int    `json:"delaySeconds"`
	Message      string `json:"message"`
}

type BluePrintCartAbandonment struct {
	Status BluePrintStatus                `json:"status"`
	Detail BluePrintCartAbandonmentDetail `json:"detail"`
}

type BluePrintConfirmationOfSafetyDetail struct {
	SafetyLevel int    `json:"safetyLevel"`
	Message     string `json:"message"`
}

type BluePrintConfirmationOfSafety struct {
	Status BluePrintStatus                     `json:"status"`
	Detail BluePrintConfirmationOfSafetyDetail `json:"detail"`
}

func main() {
	err := run()
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	bluePrintCA := BluePrintCartAbandonment{
		Status: BluePrintStatusActive,
		Detail: BluePrintCartAbandonmentDetail{
			DelaySeconds: 60,
			Message:      "Cart abandonment process is active.",
		},
	}

	bluePrintCOS := BluePrintConfirmationOfSafety{
		Status: BluePrintStatusPending,
		Detail: BluePrintConfirmationOfSafetyDetail{
			SafetyLevel: 5,
			Message:     "Confirmation of safety is pending.",
		},
	}

	fmt.Printf("%+v\n", bluePrintCA)
	fmt.Printf("%+v\n", bluePrintCOS)

	return nil
}
