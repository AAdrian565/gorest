package web

import (
	"net/http"
	"os"
)

func ErrorCheck(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Serve(fileName string, status ...int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filePath := "./pages/view/" + fileName + ".html"
		file, err := os.Open(filePath)
		if err != nil {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		fi, err := file.Stat()
		if err != nil {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}

		responseStatus := http.StatusOK
		if len(status) > 0 {
			responseStatus = status[0]
		}

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(responseStatus)
		http.ServeContent(w, r, fileName+".html", fi.ModTime(), file)
	}
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Serve("404", http.StatusNotFound)(w, r)
		return
	}
	Serve("index")(w, r)
}
