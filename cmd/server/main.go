package main

import (
	downloadServer "github.com/shashank-razor/DownloadManager/internal/downloadManagerServer"
	"github.com/shashank-razor/DownloadManager/rpc/downloadManager"
	"log"
	"net/http"
)

func main() {
	db := downloadServer.GetDBconn()
	repository, dBerr := downloadServer.NewRepo(db)
	if dBerr != nil {
		log.Fatal("Database Connection Error")
	}
	server := &downloadServer.Server{
		MysqlRepository: repository,
	} // implements Haberdasher interface
	twirpHandler := downloadManager.NewDownloaderServer(server)
	err :=http.ListenAndServe(":8080", twirpHandler)
	if err != nil {
		return
	}
}
