package fibonacci

import (
	"context"
	"github.com/RezaZahedi/Go-Gin/model/proto/fibo_model"
	"strconv"
)

func Calulate(input int) string {
	return strconv.Itoa(input *10)
}

type FiboGenerator struct{}

func (g *FiboGenerator) GenerateNumber(ctx context.Context, req *fibo_model.Request, rsp *fibo_model.Response) error {
	rsp.Output = strconv.Itoa(int(req.Input) * 2)
	return nil
}

