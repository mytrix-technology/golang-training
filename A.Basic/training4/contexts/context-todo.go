package contexts

type todoCtx struct{ emptyCtx }

func (todoCtx) String() string {
	return "context.TODO"
}

func TODO() Context {
	return todoCtx{}
}
