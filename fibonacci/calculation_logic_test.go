package fibonacci

import (
	"context"
	"github.com/RezaZahedi/Go-Gin/model/proto/fibo_model"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestNewFiboGeneratorSequential(t *testing.T) {
	fiboGenerator := NewFiboGenerator()
	defer fiboGenerator.Close()
	type args struct {
		ctx context.Context
		req *fibo_model.Request
		res *fibo_model.Response
	}
	ctx := context.TODO()
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{ctx, &fibo_model.Request{Input: 1}, &fibo_model.Response{}}, "1"},
		{"2", args{ctx, &fibo_model.Request{Input: 2}, &fibo_model.Response{}}, "1"},
		{"8", args{ctx, &fibo_model.Request{Input: 8}, &fibo_model.Response{}}, "21"},
		{"70", args{ctx, &fibo_model.Request{Input: 70}, &fibo_model.Response{}}, "190392490709135"},
		{"10", args{ctx, &fibo_model.Request{Input: 10}, &fibo_model.Response{}}, "55"},
		//{"-10", args{ctx, &fibo_model.Request{Input: -10}, &fibo_model.Response{}}, "55"},
		//
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := fiboGenerator.GenerateNumber(test.args.ctx, test.args.req, test.args.res)
			assert.Nil(t, err)
			assert.Equal(t, test.want, test.args.res.Output)
		})
	}
}

func TestNewFiboGeneratorParallel(t *testing.T) {
	t.Parallel()
	fiboGenerator := NewFiboGenerator()

	// to make sure that we release the resources
	runtime.SetFinalizer(fiboGenerator,
		func(f *FiboGenerator) {
			f.Close()
			t.Log("resource released!")
		})

	type args struct {
		ctx context.Context
		req *fibo_model.Request
		res *fibo_model.Response
	}
	ctx := context.TODO()
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{ctx, &fibo_model.Request{Input: 1}, &fibo_model.Response{}}, "1"},
		{"2", args{ctx, &fibo_model.Request{Input: 2}, &fibo_model.Response{}}, "1"},
		{"8", args{ctx, &fibo_model.Request{Input: 8}, &fibo_model.Response{}}, "21"},
		{"70", args{ctx, &fibo_model.Request{Input: 70}, &fibo_model.Response{}}, "190392490709135"},
		{"10", args{ctx, &fibo_model.Request{Input: 10}, &fibo_model.Response{}}, "55"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := fiboGenerator.GenerateNumber(test.args.ctx, test.args.req, test.args.res)
			assert.Nil(t, err)
			assert.Equal(t, test.want, test.args.res.Output)
		})
	}
}

func TestNewFiboGeneratorPanicCase(t *testing.T) {
	fiboGenerator := NewFiboGenerator()
	defer fiboGenerator.Close()
	type args struct {
		ctx context.Context
		req *fibo_model.Request
		res *fibo_model.Response
	}
	ctx := context.TODO()
	tests := []struct {
		name string
		args args
		want string
	}{
		{"-10", args{ctx, &fibo_model.Request{Input: -10}, &fibo_model.Response{}}, "Nothing"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("The code did not panic")
				}
			}()

			err := fiboGenerator.GenerateNumber(test.args.ctx, test.args.req, test.args.res)
			assert.Nil(t, err)
			assert.Equal(t, test.want, test.args.res.Output)
		})
	}
}
