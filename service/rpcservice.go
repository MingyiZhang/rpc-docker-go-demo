package service

import "errors"

type DemoService struct {}

type Args struct {
  A, B float64
}

func (DemoService) Div(args Args, result *float64) error {
  if args.B == 0 {
    return errors.New("division by zero")
  }

  *result = args.A / args.B
  return nil
}
