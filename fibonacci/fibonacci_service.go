package fibonacci

type FibonacciService struct {
	FibonacciCalculator *func(int) int
}

func ProvideFibonacciService(fibonacciCalculator *func(int) int) FibonacciService {
	return FibonacciService{FibonacciCalculator: fibonacciCalculator}
}