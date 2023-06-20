package main

import "context"

type ABProcessor struct {
	outA chan AOut
	outB chan BOut
	errs chan error
}

func NewABProcessor() *ABProcessor {
	return &ABProcessor{
		outA: make(chan AOut, 1),
		outB: make(chan BOut, 1),
		errs: make(chan error, 2),
	}
}

func (p *ABProcessor) start(ctx context.Context, data Input) {
	go func() {
		aOut, err := getResultA(ctx, data.A)
		if err != nil {
			p.errs <- err
			return
		}
		p.outA <- aOut
	}()
	go func() {
		bOut, err := getResultB(ctx, data.B)
		if err != nil {
			p.errs <- err
			return
		}
		p.outB <- bOut
	}()
}

func (p *ABProcessor) wait(ctx context.Context) (CIn, error) {
	var inputC CIn
	for count := 0; count < 2; count++ {
		select {
		case a := <-p.outA:
			inputC.A = a
		case b := <-p.outB:
			inputC.B = b
		case err := <-p.errs:
			return CIn{}, err
		case <-ctx.Done():
			return CIn{}, ctx.Err()
		}
	}
	return inputC, nil
}

type AOut struct {
}

type BOut struct {
}

type CIn struct {
	A AOut
	B BOut
}

func getResultA(ctx context.Context, in string) (AOut, error) {
	return AOut{}, nil
}

func getResultB(ctx context.Context, in string) (BOut, error) {
	return BOut{}, nil
}
