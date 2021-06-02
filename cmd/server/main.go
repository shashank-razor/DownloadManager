package main

import (
	downloadServer "github.com/shashank-razor/DownloadManager/internal/downloadManagerServer"
	"github.com/shashank-razor/DownloadManager/rpc/downloadManager"
	"net/http"
)

func main() {
	server := &downloadServer.Server{} // implements Haberdasher interface
	twirpHandler := downloadManager.NewDownloaderServer(server)
	err := http.ListenAndServe(":8080", twirpHandler)
	if err != nil {
		return
	}
}
