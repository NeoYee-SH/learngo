package main

import (
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/gpmgo/gopm/modules/log"

	"github.com/yihuaiyuan/learngo/errhandling/listserver/filelisting"
)

type userError interface {
	error
	Message() string
}
type appHandle func(writer http.ResponseWriter, request *http.Request) error

func erpWrapper(handler appHandle) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			if r := recover(); r != nil {
				log.Warn("panic : %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)

		if err != nil {
			log.Warn("Error handling request : %s", err.Error())

			//userError
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

//显示服务器的文件内容
func main() {
	http.HandleFunc("/", erpWrapper(filelisting.Handle))
	err := http.ListenAndServe(":7701", nil)

	if err != nil {
		panic(err)
	}

}
