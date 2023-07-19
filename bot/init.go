package bot

import (
	"regexp"
	"strings"

	"github.com/XiaoMengXinX/Music163Api-Go/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

// MusicDB Music cache database entry
var MusicDB *gorm.DB

// config Configuration file data
var config map[string]string

// data NetEase Cloud cookie
var data utils.RequestData

var bot *tgbotapi.BotAPI
var botAdmin []int
var botAdminStr []string
var botName string
var cacheDir = "./cache"
var botAPI = "https://api.telegram.org"

// maxRetryTimes Maximum retry times, downloaderTimeout Download timeout
var maxRetryTimes, downloaderTimeout int

var (
	reg1   = regexp.MustCompile(`(.*)song\?id=`)
	reg2   = regexp.MustCompile("(.*)song/")
	regP1  = regexp.MustCompile(`(.*)program\?id=`)
	regP2  = regexp.MustCompile("(.*)program/")
	regP3  = regexp.MustCompile(`(.*)dj\?id=`)
	regP4  = regexp.MustCompile("(.*)dj/")
	reg5   = regexp.MustCompile("/(.*)")
	reg4   = regexp.MustCompile("&(.*)")
	reg3   = regexp.MustCompile(`\?(.*)`)
	regInt = regexp.MustCompile(`\d+`)
	regUrl = regexp.MustCompile("(http|https)://[\\w\\-_]+(\\.[\\w\\-_]+)+([\\w\\-.,@?^=%&:/~+#]*[\\w\\-@?^=%&/~+#])?")
)

var mdV2Replacer = strings.NewReplacer(
	"_", "\\_", "*", "\\*", "[", "\\[", "]", "\\]", "(",
	"\\(", ")", "\\)", "~", "\\~", "`", "\\`", ">", "\\>",
	"#", "\\#", "+", "\\+", "-", "\\-", "=", "\\=", "|",
	"\\|", "{", "\\{", "}", "\\}", ".", "\\.", "!", "\\!",
)

var (
	aboutText = `*Music163bot-Go v2*
Github: https://github.com/XiaoMengXinX/Music163bot-Go

\[Build Environment] %s
\[Build Version] %s
\[Build Hash] %s
\[Build Date] %s
\[Runtime Environment] %s`
	musicInfo = `"%s" - %s
Album: %s
#NetEaseMusic #%s %.2fMB %.2fkbps
via @%s`
	musicInfoMsg = `%s
Album: %s
%s %.2fMB
`
	uploadFailed = `Download/Send failed
%v`
	statusInfo = `*\[Statistics\]*
Total cached songs in the database: %d
Number of cached songs in the current session \[%s\]: %d
Number of cached songs for the current user \[[%d](tg://user?id=%d)\]: %d
`
	rmcacheReport    = `Clear [%s] cache successful`
	inputKeyword     = "Please enter search keywords"
	inputIDorKeyword = "Please enter the song ID or song keywords"
	inputContent     = "Please enter song keywords/song sharing link/song ID"
	searching        = `Searching...`
	noResults        = `No results found`
	noCache          = `Song not cached`
	tapToDownload    = `Click the button above to cache the song`
	tapMeToDown      = `Tap me to cache the song`
	hitCache         = `Cache hit, sending...`
	sendMeTo         = `Send me to...`
	getLrcFailed     = `Failed to get lyrics, the song may not exist or it's instrumental`
	getUrlFailed     = `Failed to get song download link`
	fetchInfo        = `Fetching song information...`
	fetchInfoFailed  = `Failed to fetch song information`
	waitForDown      = `Waiting for download...`
	downloading      = `Downloading...`
	downloadStatus   = " %s\n%.2fMB/%.2fMB %d%%"
	redownloading    = `Download failed, retrying...`
	uploading        = `Download complete, sending...`
	md5VerFailed     = "MD5 verification failed"
	reTrying         = "Retrying (%d/%d)"
	retryLater       = "Please retry later"

	reloading    = "Reloading"
	callbackText = "Success"

	fetchingLyric   = "Fetching lyrics"
	downloadTimeout = `Download timeout`
)
