# About

Full-Stack Web-App тти AnimeLayer.ru

[Hosted on serv00](https://dmji.serv00.net/animelayer)

[Telegram Bot with MiniApp](https://t.me/MyMediaNotifyBot/web)

#### Stack:

* go
* a-h/templ + htmx + htmx websocket + tailwind
* postgresql + sqlc
* go-i18n

#### Tools:

* air
* tailwind cli
* goose
* taskfiles
* [go-stringer](https://github.com/dmji/go-stringer "stringer fork for enums")

#### DevOps

* Host on free serv00 with auto-deploy via github actions
* Chron (scheduler configuration serv00 web panel) auto-update with parsing AnimeLayer each hour

### Work-In-Progress TODO

* [X] Add presentation with htmx based on v0.dev generation
* [X] Implement AnimeLayer parser in separate [repository](https://github.com/dmji/go-animelayer-parser)
* [X] Initialize postgresql with goose migrations
* [X] Make endless scroll pages with filtering via query parameters
* [X] Add string russian translation with goi18n
* [X] Add ([created fork](https://github.com/dmji/go-stringer)) tool to generate enums with stringer, parser and consts as id for goi18n
* [X] Improve goi18n to correctly lookup consts in other files ([fork](https://github.com/dmji/go-i18n/tree/main), [pull request](https://github.com/nicksnyder/go-i18n/pull/355))
* [X] Host on free serv00 with auto-deploy via github actions
* [X] Add settings page with lang and theme selection
* [ ] Add colors for dark theme
  * [X] Init dark theme detectable from the light
  * [ ] Align all colors for dark-contrasts
* [X] Implement auto-update from AnimeLayer vith chron
* [ ] Add notification with tg bot api
* [ ] Improve filtering to use some key words from notes (genre, year etc)
* [ ] Add connection to MyAnimeList
* [ ] Profile page
* [ ] OAuth with Telegram

# Project Structure

`assets` - anything that might requie in runtime (CSS, images etc)

`langlang/translations` - folder to store localized strings as i18n toml files

`build` - should contain dockerfiles and other files requied in build-stage

`cmd` - inheret from clean architure folder for executable applications

`internal` - inheret from clean architure folder

`pkg` - inheret from clean architure folder

`components` - folder with a-h/templ template files

`handlers` - folder with front-end handlers routing

`migrations` - sequence migration files to initialize postgres with goose

# Building

### For AIR developing without docker

```bash
task update-apps
task dev
```

### For AIR developing with docker

```
TODO
```

### For serv00 deployment

```bash
task update-apps
task tailwind:prod
task pre-build-prod
task build-freebsd-amd64
```

# Auto-Deployment to serv00

Reuied secrets for git hub action

* FTP_SERVER
* FTP_USER
* FTP_PASSWORD
* SSH_HOST
* SSH_USER
* SSH_KEY      - raw private ssh-key (password not work for me)
* SSH_PORT     - by default it`s 22

serv00 using freebsd on free servers so docker is not an option.

So I crafted restart.bash in root of ssh (or ftp):

```bash
pkill -f ./server
rm -f repo/gosudarevlist/server
mv server repo/gosudarevlist/server
cd repo/gosudarevlist
chmod 777 server
```

It trigger from GitHub Actions on push after upload builded for freebsd amd64 server.
