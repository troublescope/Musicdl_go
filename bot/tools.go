package bot

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/XiaoMengXinX/Music163Api-Go/api"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/sirupsen/logrus"
)

// Check if a string is in an array
func in(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

// Parse composer information
func parseArtist(songDetail types.SongDetailData) string {
	var artists string
	for i, ar := range songDetail.Ar {
		if i == 0 {
			artists = ar.Name
		} else {
			artists = fmt.Sprintf("%s/%s", artists, ar.Name)
		}
	}
	return artists
}

// Check if a directory exists or create it
func dirExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			logrus.Errorf("mkdir %v failed: %v\n", path, err)
		}
		return false
	}
	logrus.Errorf("Error: %v\n", err)
	return false
}

// Verify MD5 checksum
func verifyMD5(filePath string, md5str string) (bool, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer f.Close()
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		return false, err
	}
	if hex.EncodeToString(md5hash.Sum(nil)) != md5str {
		return false, fmt.Errorf(md5VerFailed)
	}
	return true, nil
}

// Parse MusicID from text
func parseMusicID(text string) int {
	var replacer = strings.NewReplacer("\n", "", " ", "")
	messageText := replacer.Replace(text)
	musicUrl := regUrl.FindStringSubmatch(messageText)
	if len(musicUrl) != 0 {
		if strings.Contains(musicUrl[0], "song") {
			ur, _ := url.Parse(musicUrl[0])
			id := ur.Query().Get("id")
			if musicid, _ := strconv.Atoi(id); musicid != 0 {
				return musicid
			}
		}
	}
	musicid, _ := strconv.Atoi(linkTestMusic(messageText))
	return musicid
}

// Parse ProgramID from text
func parseProgramID(text string) int {
	var replacer = strings.NewReplacer("\n", "", " ", "")
	messageText := replacer.Replace(text)
	programid, _ := strconv.Atoi(linkTestProgram(messageText))
	return programid
}

// Extract integers from text
func extractInt(text string) string {
	matchArr := regInt.FindStringSubmatch(text)
	if len(matchArr) == 0 {
		return ""
	}
	return matchArr[0]
}

// Parse sharing link for MusicID
func linkTestMusic(text string) string {
	return extractInt(reg5.ReplaceAllString(reg4.ReplaceAllString(reg3.ReplaceAllString(reg2.ReplaceAllString(reg1.ReplaceAllString(text, ""), ""), ""), ""), ""))
}

func linkTestProgram(text string) string {
	return extractInt(reg5.ReplaceAllString(reg4.ReplaceAllString(reg3.ReplaceAllString(regP4.ReplaceAllString(regP3.ReplaceAllString(regP2.ReplaceAllString(regP1.ReplaceAllString(text, ""), ""), ""), ""), ""), ""), ""))
}

// Check if an error is a timeout error
func isTimeout(err error) bool {
	if strings.Contains(fmt.Sprintf("%v", err), "context deadline exceeded") {
		return true
	}
	return false
}

// Get the MusicID for a radio program
func getProgramRealID(programID int) int {
	programDetail, err := api.GetProgramDetail(data, programID)
	if err != nil {
		return 0
	}
	if programDetail.Program.MainSong.ID != 0 {
		return programDetail.Program.MainSong.ID
	}
	return 0
}
