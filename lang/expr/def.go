package expr

var symbols = map[string]*interface{}{}

func DefExpr(symbol string, value interface{}) {
	symbols[symbol] = &value
}

func lookupSymbol(symbol string) *interface{} {
	return symbols[symbol]
}
