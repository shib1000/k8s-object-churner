package worker

import "math/rand"

var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

type Worker interface {
	Churn()
	Delete()
}

type Context struct {
	values map[string]interface{}
}

func BuildNewContext() *Context {
	return &Context{}
}

func (ctx *Context) SetContextString(key string, value string) {
	ctx.values[key] = value
}

func (ctx *Context) GetContextString(key string) *string {
	val, _ := ctx.values[key].(string)
	return &val
}

func (ctx *Context) SetContext(key string, value interface{}) {
	ctx.values[key] = value
}

func (ctx *Context) GetContext(key string) *interface{} {
	val := ctx.values[key]
	return &val

}

func RandomString(n int) string {
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
