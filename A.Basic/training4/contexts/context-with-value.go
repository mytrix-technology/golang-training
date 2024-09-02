package contexts

type valueCtx struct {
	Context
	key, val any
}

func WithValue(parent Context, key, val any) Context {
	return &valueCtx{parent, key, val}
}

func (c *valueCtx) Value(key any) any {
	// If it this one, just return it
	if c.key == key {
		return c.val
	}
	return value(c.Context, key)
}

func value(c Context, key any) any {
	for {
		switch ctx := c.(type) {
		case *valueCtx:
			// If the parent is a `valueCtx`, check its key
			if key == ctx.key {
				return ctx.val
			}
			c = ctx.Context
		case backgroundCtx, todoCtx:
			// If we have reached the top, ie base context
			// Return as we did not find anything
			return nil
		default:
			// If it is some other context,
			// just calls its `.Value` method
			return c.Value(key)
		}
	}
}
