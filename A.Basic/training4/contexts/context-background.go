package contexts

import "context"

type Context = context.Context

type backgroundCtx struct{ emptyCtx }

func (backgroundCtx) String() string {
	return "context.Background"
}

func Background() Context {
	return backgroundCtx{}
}
