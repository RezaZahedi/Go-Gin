package fibonacci

import (
	"context"
	"github.com/RezaZahedi/Go-Gin/model/proto/fibo_model"
	"math/big"
)

type FiboGenerator struct{
	cache map[int]*big.Int
}

func NewFiboGenerator() *FiboGenerator {
	return &FiboGenerator{cache: make(map[int]*big.Int)}
}

func (g *FiboGenerator) GenerateNumber(ctx context.Context, req *fibo_model.Request, rsp *fibo_model.Response) error {
	rsp.Output = Calulate(int(req.Input), g.cache)
	return nil
}

func Calulate(input int, cache map[int]*big.Int) string {

	if input == 0 { return "0" }

	one := big.NewInt(1)
	var fibo func(a int) *big.Int
	fibo = func(a int) *big.Int {
		if a == 1 || a == 2 {
			return one
		}
		b := new(big.Int)
		if b, ok := cache[a]; ok {
			return b
		}
		b.Add(fibo(a - 2), fibo(a - 1))
		cache[a] = b
		return b
	}

	return fibo(input).String()
}