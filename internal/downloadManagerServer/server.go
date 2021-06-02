package downloadManagerServer

import (
	"context"
	pb "github.com/shashank-razor/DownloadManager/rpc/downloadManager"
	"log"
)

type Server struct {}

func (s *Server) Downloads(ctx context.Context,request *pb.DownloadRequest) (response *pb.DownloadResponse,err error) {
	if modeOfDownload:= request.Type ; modeOfDownload=="serial" {
        log.Printf("serial download request")
	} else if modeOfDownload=="concurrent" {
		log.Printf("current download request")
	} else {
		log.Printf("invalid mode request")
	}
	return &pb.DownloadResponse{
		Id: "12345",
	},nil
}