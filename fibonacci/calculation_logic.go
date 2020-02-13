package fibonacci

import (
	"context"
	"github.com/RezaZahedi/Go-Gin/model/memo"
	"github.com/RezaZahedi/Go-Gin/model/proto/fibo_model"
	"math/big"
)

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

//func Calculate(input int, cache map[int]*big.Int) string {
//	if input < 0 {
//		panic("Calculate: input must be positive")
//	}
//	if input == 0 { return "0" }
//
//	one := big.NewInt(1)
//	var fibo func(a int) *big.Int
//	fibo = func(a int) *big.Int {
//		if a == 1 || a == 2 {
//			return one
//		}
//		b := new(big.Int)
//		if b, ok := cache[a]; ok {
//			return b
//		}
//		b.Add(fibo(a - 2), fibo(a - 1))
//		cache[a] = b
//		return b
//	}
//
//	return fibo(input).String()
//}

var fibo func(a int) *big.Int

func fiboTemp(key int) interface{} {
	ans := fibo(key)
	return ans
}

func calculate(input int, m *memo.Memo) string {
	one := big.NewInt(1)

	//m := New(fiboTemp)
	//defer m.Close()

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

