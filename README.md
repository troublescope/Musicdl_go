<h1 align="center">Musicdl_bot</h1>

<h4 align="center">A Telegram bot for downloading/sharing/searching NetEase Cloud Music songs</h4>

<p align="center">Demo bot：<a href="https://t.me/Music163bot">https://t.me/Music163bot</a></p>

<p align="center">
	<a href="https://goreportcard.com/report/github.com/XiaoMengXinX/Music163bot-Go/v2">
      <img src="https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=flat">
	</a>
	<a href="https://github.com/XiaoMengXinX/Music163bot-Go/releases">
    <img src="https://img.shields.io/github/v/release/XiaoMengXinX/Music163bot-Go?include_prereleases&style=flat-square">
  </a>
</p>

## ✨ Features

- Share link sniffing
- inlinebot
- Song search
- Add 163key to song files
- Quick song sharing
- Download lossless flac audio (requires setting MUSIC_U of NetEase Cloud VIP account)
- Dynamic update (using [traefik/yaegi](https://github.com/traefik/yaegi) as a dynamic expansion framework)

## ⚙️ Build

Please make sure you have `Go 1.17` or higher before building

**Clone the code**

```
git clone https://github.com/XiaoMengXinX/Music163bot-Go
```

**Automatically compile using script (supports windows bash environment, such as git bash)**

```
cd Music163bot-Go
bash build.sh 

# You can also add environment variables for cross-compilation, such as
GOOS=windows GOARCH=amd64 bash build.sh
```

## 🛠️ Deployment

**Modify the configuration file**

Open `config_example.ini` in the project root directory

```
# The following are required
# Your Bot Token
BOT_TOKEN = YOUR_BOT_TOKEN

# The value of MUSIC_U in your NetEase Cloud cookie (used to download lossless songs)
MUSIC_U = YOUR_MUSIC_U


# The following are optional
# Custom telegram bot API address
BotAPI = https://api.telegram.org

# Set bot administrator ID, separated by ","
BotAdmin = 1234,3456

# Whether to turn on the debug function of the bot
BotDebug = false

# Custom sqlite3 database file (default is cache.db)
Database = cache.db

# Set log level [panic|fatal|error|warn|info|debug|trace] (default is info)
LogLevel = info

# Whether to turn on automatic updates (default is on), if set to false, it is equivalent to the -no-update parameter
AutoUpdate = true
 
# Whether to automatically redownload corrupted download files (default is true)
AutoRetry = true

# Maximum number of automatic retries (default is 3)
MaxRetryTimes = 3

# Download timeout duration (in seconds, default is 60)
DownloadTimeout = 60

# Whether to check the md5 of the update file (default is on), if set to false, it is equivalent to the -no-md5-check parameter
CheckMD5 = true

# Custom source code path
SrcPath = ./src

# Custom bot function entry (default is bot.Start)
BotEntry = bot.Start
```

**※ After modifying the configuration, rename `config_example.ini` to `config.ini`**

**Start Music163-bot**

```
$ ./Music163bot-Go
2021/10/30 13:05:40 [INFO] Music163bot-Go v2.0.0(20000) (main.go:122)
2021/10/30 13:05:40 [INFO] Checking for updates (main.go:155)
2021/10/30 13:05:40 [INFO] v2.0.0(20000) is the latest version (main.go:361)
2021/10/30 13:05:40 [INFO] Verifying file MD5 (main.go:164)
2021/10/30 13:05:40 [INFO] MD5 verification succeeded (main.go:169)
2021/10/30 13:05:40 [INFO] Loading version v2.0.0(20000) (main.go:195)
2021/10/30 13:05:41 [INFO] Music163bot verification succeeded (value.go:543)
```

## 🤖 Commands

- `/musicid` or `/netease` + `MusicID` —— Get songs from MusicID
- `/search` + `keyword` —— Search for songs
- `/about` —— About this bot