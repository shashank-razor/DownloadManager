package downloadManagerServer

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type Download struct{
	DownloadId string `json:"id"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
	Status DownloadStatus `json:"status"`
	Mode ModeOfDownload `json:"mode"`
	Files []string `json:"files"`
}

type Repository interface {
	CreateDownload(ctx context.Context, download Download) error
}

type repo struct {
	db   *sql.DB
}

func GetDBconn() *sql.DB {
   db, err := sql.Open("mysql",  "root:mysql-razor@tcp(127.0.0.1:3306)/downloadManager")
   if err != nil {
   	panic(err.Error())
   }
   return db

}

//Creates and returns an instance
func NewRepo(db *sql.DB) (repo, error) {
	return repo{
		db: db,
	}, nil
}

func (repo *repo) CreateDownload(ctx context.Context, download Download) error{
	var query = "INSERT INTO Download(id,start_date,end_date,status,mode,files) VALUES (?, ?, ?, ?, ?, ?"
	_, err := repo.db.ExecContext(
		ctx,
		query,
		download.DownloadId,
		download.StartTime,
		download.EndTime,
		download.Status,
		download.Mode,
		download.Files)
	if err != nil {
		log.Printf("Error occurred inside CreateDownload in repo")
		return err
	} else {
		log.Printf("Download item added:", download.DownloadId)
	}
	return nil
}