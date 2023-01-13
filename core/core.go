package core

type function func(list []interface{})

var identifiers = map[string]function{
	"defn": defn,
}

func defn(list []interface{}) {}
