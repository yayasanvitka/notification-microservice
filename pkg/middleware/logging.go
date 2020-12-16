package middleware

import (
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	"net/http"
)

//func LogRequest(next httprouter.Handle) httprouter.Handle {
//	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
//		logger.Info.Printf(
//			"%s - %s %s %s",
//			request.RemoteAddr,
//			request.Proto,
//			request.Method,
//			request.URL.RequestURI(),
//		)
//
//		next(writer, request, params)
//	}
//}

func LogRequest(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		zap.S().Info("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next(w, r, p)
	}
}