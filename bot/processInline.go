package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/XiaoMengXinX/Music163Api-Go/api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func processInlineMusic(musicid int, query tgbotapi.InlineQuery, bot *tgbotapi.BotAPI) (err error) {
	var songInfo SongInfo
	db := MusicDB.Session(&gorm.Session{})
	err = db.Where("music_id = ?", musicid).First(&songInfo).Error // Check if there is cached data
	if err == nil {                                                // Respond to inlineQuery with cached data
		if songInfo.FileID != "" && songInfo.SongName != "" {
			var numericKeyboard tgbotapi.InlineKeyboardMarkup
			numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonURL(fmt.Sprintf("%s - %s", songInfo.SongName, songInfo.SongArtists), fmt.Sprintf("https://music.163.com/song?id=%d", songInfo.MusicID)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonSwitch(sendMeTo, fmt.Sprintf("https://music.163.com/song?id=%d", songInfo.MusicID)),
				),
			)

			newAudio := tgbotapi.NewInlineQueryResultCachedDocument(query.ID, songInfo.FileID, fmt.Sprintf("%s - %s", songInfo.SongArtists, songInfo.SongName))
			newAudio.Caption = fmt.Sprintf(musicInfo, songInfo.SongName, songInfo.SongArtists, songInfo.SongAlbum, songInfo.FileExt, float64(songInfo.MusicSize+songInfo.EmbPicSize)/1024/1024, float64(songInfo.BitRate)/1000, botName)
			newAudio.ReplyMarkup = &numericKeyboard
			newAudio.Description = songInfo.SongAlbum

			inlineConf := tgbotapi.InlineConfig{
				InlineQueryID: query.ID,
				Results:       []interface{}{newAudio},
				IsPersonal:    false,
				CacheTime:     3600,
			}

			_, err := bot.Request(inlineConf)
			if err != nil {
				return err
			}
		}
	} else {
		inlineMsg := tgbotapi.NewInlineQueryResultArticle(query.ID, noCache, query.Query)
		inlineMsg.Description = tapToDownload

		inlineConf := tgbotapi.InlineConfig{
			InlineQueryID:     query.ID,
			IsPersonal:        false,
			Results:           []interface{}{inlineMsg},
			CacheTime:         60,
			SwitchPMText:      tapMeToDown,
			SwitchPMParameter: fmt.Sprintf("%d", musicid),
		}

		_, err := bot.Request(inlineConf)
		if err != nil {
			return err
		}
	}
	return nil
}

func processEmptyInline(message tgbotapi.InlineQuery, bot *tgbotapi.BotAPI) (err error) {
	inlineMsg := tgbotapi.NewInlineQueryResultArticle(message.ID, "Enter help to get assistance", "Music163bot-Go v2")
	inlineConf := tgbotapi.InlineConfig{
		InlineQueryID: message.ID,
		IsPersonal:    false,
		Results:       []interface{}{inlineMsg},
		CacheTime:     3600,
	}
	_, err = bot.Request(inlineConf)
	if err != nil {
		return err
	}
	return err
}

func processInlineHelp(message tgbotapi.InlineQuery, bot *tgbotapi.BotAPI) (err error) {
	randomID := time.Now().UnixMicro()
	inlineMsg1 := tgbotapi.NewInlineQueryResultArticle(fmt.Sprintf("%d", randomID), "1. Paste music sharing URL or enter MusicID", "Music163bot-Go v2")
	inlineMsg2 := tgbotapi.NewInlineQueryResultArticle(fmt.Sprintf("%d", randomID+1), "2. Enter search+keyword to search for songs", "Music163bot-Go v2")
	inlineConf := tgbotapi.InlineConfig{
		InlineQueryID: message.ID,
		IsPersonal:    false,
		Results:       []interface{}{inlineMsg1, inlineMsg2},
		CacheTime:     3600,
	}
	_, err = bot.Request(inlineConf)
	if err != nil {
		return err
	}
	return err
}

func processInlineSearch(message tgbotapi.InlineQuery, bot *tgbotapi.BotAPI) (err error) {
	randomID := time.Now().UnixMicro()
	keyWord := strings.Replace(message.Query, "search", "", 1)
	if keyWord == "" {
		inlineMsg := tgbotapi.NewInlineQueryResultArticle(fmt.Sprintf("%d", randomID), "Please enter a keyword", "Music163bot-Go v2")
		inlineConf := tgbotapi.InlineConfig{
			InlineQueryID: message.ID,
			IsPersonal:    false,
			Results:       []interface{}{inlineMsg},
			CacheTime:     3600,
		}
		_, err = bot.Request(inlineConf)
		return err
	}
	result, err := api.SearchSong(data, api.SearchSongConfig{
		Keyword: keyWord,
	})
	if err != nil {
		return err
	}
	searchResult := result
	if len(searchResult.Result.Songs) == 0 {
		inlineMsg := tgbotapi.NewInlineQueryResultArticle(fmt.Sprintf("%d", randomID), noResults, noResults)
		inlineConf := tgbotapi.InlineConfig{
			InlineQueryID: message.ID,
			IsPersonal:    false,
			Results:       []interface{}{inlineMsg},
			CacheTime:     3600,
		}
		_, err = bot.Request(inlineConf)
		return err
	}
	var inlineMsgs []interface{}
	for i := 0; i < len(searchResult.Result.Songs) && i < 10; i++ {
		var songArtists string
		for i, artist := range searchResult.Result.Songs[i].Artists {
			if i == 0 {
				songArtists = artist.Name
			} else {
				songArtists = fmt.Sprintf("%s/%s", songArtists, artist.Name)
			}
		}
		inlineMsg := tgbotapi.NewInlineQueryResultArticle(fmt.Sprintf("%d", randomID+int64(i)), searchResult.Result.Songs[i].Name, fmt.Sprintf("/netease %d", searchResult.Result.Songs[i].Id))
		inlineMsg.Description = songArtists
		inlineMsgs = append(inlineMsgs, inlineMsg)
	}
	inlineConf := tgbotapi.InlineConfig{
		InlineQueryID: message.ID,
		IsPersonal:    false,
		Results:       inlineMsgs,
		CacheTime:     3600,
	}
	_, err = bot.Request(inlineConf)
	return err
}
