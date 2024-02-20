package fibonacci

import (
	"context"
	"fmt"

	"github.com/jgfranco17/golang-api-template/service/pkg/logging"
)

func fibonacciValue(number int) int {
	if number <= 1 {
		return number
	}

	a, b := 0, 1
	for i := 2; i <= number; i++ {
		a, b = b, a+b
	}
	return b
}

func Fibonacci(limit int) ([]uint64, error) {
	ctx := context.WithValue(context.Background(), "section", "Fibonacci")
	log := logging.GetLogger(ctx)
	result := []uint64{0}
	hardCap := 93
	if limit == 1 {
		result = append(result, 1)
	} else {
		if limit >= hardCap {
			return nil, fmt.Errorf("Maximum limit is %d", hardCap)
		}
		result = []uint64{0, 1}
		for i := 2; i < limit; i++ {
			nextFib := result[i-1] + result[i-2]
			result = append(result, nextFib)
		}
	}

	log.Printf("Generated Fibonacci number sequence for %d", limit)
	return result, nil
}
