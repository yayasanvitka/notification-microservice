package middleware

import (
	"github.com/julienschmidt/httprouter"
)

type Middleware func(httprouter.Handle) httprouter.Handle

//func Chain (f httprouter.Handle, m ...Middleware) httprouter.Handle {
//	if len(m) == 0 {
//		return f
//	}
//
//	return m[0](Chain(f, m[1:]...))
//}

// Chain - chains all middleware functions right to left
// https://husobee.github.io/golang/http/middleware/2015/12/22/simple-middleware.html
func Chain(f httprouter.Handle, m ...Middleware) httprouter.Handle {
	// if our chain is done, use the original handlerfunc
	if len(m) == 0 {
		return f
	}
	// otherwise run recursively over nested handlers
	return m[0](Chain(f, m[1:]...))
}
