package downloadManagerServer

import (
	"context"
	"github.com/satori/go.uuid"
	pb "github.com/shashank-razor/DownloadManager/rpc/downloadManager"
	"log"
	"time"
)

type Server struct {
	MysqlRepository repo
}

type ModeOfDownload string

const (
	Serial ModeOfDownload = "serial"
	Concurrent ModeOfDownload = "concurrent"
)

type DownloadStatus int8

const (
	Successful DownloadStatus = 1
	Failed DownloadStatus = 2
	Waiting DownloadStatus = 3
)

func (s *Server) Downloads(ctx context.Context,request *pb.DownloadRequest) (response *pb.DownloadResponse,err error) {
	startTime := time.Now()
	status := Waiting
	if modeOfDownload:= ModeOfDownload(request.Type); modeOfDownload==Serial {
        log.Printf("serial download request")
        status=Successful
	} else if modeOfDownload==Concurrent {
		log.Printf("current download request")
		status=Successful
	} else {
		log.Printf("invalid mode request")
	}
	if status==Waiting{
		status=Failed
	}
	endTime := time.Now()
	uuid,_ := uuid.NewV4()
	downloadItems:= Download{
		DownloadId: uuid.String(),
		StartTime: startTime,
		EndTime: endTime,
		Files: request.Urls,
	}
	dberr:= s.MysqlRepository.CreateDownload(ctx,downloadItems)
	if dberr!=nil {
		log.Printf("cannot post to download item to database")
	}
	return &pb.DownloadResponse{
		Id: uuid.String(),
	},nil
}