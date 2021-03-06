package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go-demo/pkg/kindergarten"
)

type A struct {
	ctx    context.Context
	cancle context.CancelFunc
	value  string
}

func (a *A) Start() error {
	for {
		select {
		case <-a.ctx.Done():
			return errors.New("cancle!")
		default:
			fmt.Println(a.value)
		}

		time.Sleep(time.Second * 1)
	}

	return nil
}

func (a *A) Finalize() error {
	a.cancle()

	fmt.Println("-----")

	return nil
}

func NewA(value string) *A {
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)

	return &A{
		ctx:    ctx,
		cancle: cancle,
		value:  value,
	}
}

func main() {
	kind := kindergarten.NewKind()

	kind.Register(NewA("swy"))
	kind.Register(NewA("ll"))

	kind.Start()
}
