package fibonacci

import (
	"context"
	"github.com/RezaZahedi/Go-Gin/model/memo"
	"github.com/RezaZahedi/Go-Gin/model/proto/fibo_model"
	"math/big"
)

// FiboGenerator implements the GetFibonacciNumberService interface
type FiboGenerator struct{
	cache *memo.Memo
}

func NewFiboGenerator() *FiboGenerator {
	return &FiboGenerator{cache: memo.New(fiboTemp)}
}

func (g *FiboGenerator) GenerateNumber(ctx context.Context, req *fibo_model.Request, rsp *fibo_model.Response) error {
	rsp.Output = calculate(int(req.Input), g.cache)
	return nil
}

// this function is used to break the dependency loop in recursive function call and
// function memoization
var fibo func(a int) *big.Int

func fiboTemp(key int) interface{} {
	ans := fibo(key)
	return ans
}

func calculate(input int, m *memo.Memo) string {
	one := big.NewInt(1)

	fibo = func(a int) *big.Int {
		if a == 1 || a == 2 {
			return one
		}
		b := new(big.Int)
		b.Add(m.Get(a-2).(*big.Int), m.Get(a-1).(*big.Int))
		return b
	}

	return fibo(input).String()
}

