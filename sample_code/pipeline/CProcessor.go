package main

import "context"

type CProcessor struct {
	outC chan COut
	errs chan error
}

func NewCProcessor() *CProcessor {
	return &CProcessor{
		outC: make(chan COut, 1),
		errs: make(chan error, 1),
	}
}

func (p *CProcessor) start(ctx context.Context, inputC CIn) {
	go func() {
		cOut, err := getResultC(ctx, inputC)
		if err != nil {
			p.errs <- err
			return
		}
		p.outC <- cOut
	}()
}

func getResultC(ctx context.Context, c CIn) (COut, error) {
	return COut{}, nil
}

func (p *CProcessor) wait(ctx context.Context) (COut, error) {
	select {
	case out := <-p.outC:
		return out, nil
	case err := <-p.errs:
		return COut{}, err
	case <-ctx.Done():
		return COut{}, ctx.Err()
	}
}
