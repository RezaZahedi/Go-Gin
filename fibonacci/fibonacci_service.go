package fibonacci

type FibonacciService struct {
	FibonacciCalculator *func(int) (string, error)
}

func ProvideFibonacciService(fibonacciCalculator *func(int) (string, error)) FibonacciService {
	return FibonacciService{FibonacciCalculator: fibonacciCalculator}
}