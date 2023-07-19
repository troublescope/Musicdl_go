package bot

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Song information
type SongInfo struct {
	gorm.Model
	MusicID      int
	SongName     string
	SongArtists  string
	SongAlbum    string
	FileExt      string
	MusicSize    int
	PicSize      int
	EmbPicSize   int
	BitRate      int
	Duration     int
	FileID       string
	ThumbFileID  string
	FromUserID   int64
	FromUserName string
	FromChatID   int64
	FromChatName string
}

func initDB(config map[string]string) (err error) {
	database := "cache.db"
	if config["Database"] != "" {
		database = config["Database"]
	}
	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{
		Logger:      NewLogger(logger.Silent),
		PrepareStmt: true,
	})
	if err != nil {
		return err
	}
	err = db.Table("song_infos").AutoMigrate(&SongInfo{})
	if err != nil {
		return err
	}
	MusicDB = db.Table("song_infos")
	return err
}
