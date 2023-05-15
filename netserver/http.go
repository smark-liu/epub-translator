package netserver

import (
	"fmt"
	"github.com/smark-d/epub-translator/parser"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func HttpServer() {
	http.HandleFunc("/api/upload", uploadHandler)
	http.HandleFunc("/api/translate", translateHandler)
	http.HandleFunc("/api/export", exportHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func uploadHandler(writer http.ResponseWriter, request *http.Request) {
	file, _, err := request.FormFile("file")
	filename := request.FormValue("filename")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	os.MkdirAll(filepath.Join("./temp", "file"), fs.ModePerm)
	filePath := filepath.Join("./temp", "file", filename)
	newFile, err := os.Create(filePath)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, file)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(writer, "{\"filePath\": \"%v\"}", filePath)
}

func translateHandler(writer http.ResponseWriter, request *http.Request) {

	filePath := request.FormValue("filePath")
	sourceLang := request.FormValue("sourceLang")
	targetLang := request.FormValue("targetLang")
	translator := request.FormValue("translator")

	// 每10s给客户端发送一次心跳
	endSignal := make(chan bool)
	defer func() {
		endSignal <- true
		close(endSignal)
	}()
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-ticker.C:
				fmt.Fprintf(writer, "The file is being translated")
				writer.(http.Flusher).Flush()
			case <-endSignal:
				ticker.Stop()
				return
			}
		}
	}()

	outPath, err := parser.GetParser("epub", filePath, sourceLang, targetLang, translator).Parse()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(writer, "{\"outFilePath\": \"%v\"}", outPath)
}

func exportHandler(writer http.ResponseWriter, request *http.Request) {
	filePath := request.FormValue("filePath")
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	writer.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
	writer.Header().Set("Content-Type", request.Header.Get("Content-Type"))
	writer.Header().Set("Content-Length", request.Header.Get("Content-Length"))
	_, err = io.Copy(writer, file)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

}
