# About

Full-Stack Web-App to Replicate AnimeLayer.ru

[Hosted on serv00](https://dmji.serv00.net/animelayer)
[Telegram Bot with MiniApp](https://t.me/MyMediaNotifyBot/web)

#### Stack:

* go
* a-h/templ + htmx
* postgresql + sqlc
* go-i18n

#### Tools:

* air
* tailwind cli
* goose
* taskfiles

### Work-In-Progress TODO

* [X] Host on free serv00 with auto-deploy via github actions
* [ ] Improve filtering
* [ ] Profile page
* [ ] OAuth with Telegram

# Project Structure

### Assets

Generally anything that might requie in runtime (CSS, images etc)

Contains package to embed files into binary-exec

### Lang

lang/translations - folder to store localized strings as i18n toml files

Contains package to embed files into binary-exec

### Build

Should contain dockerfiles and other files requied in build-stage

### ~Builds

Temp folder as build output target

### Cmd

Inheret from clean architure folder for executable applications

cmd/env util package for updaters apps `TODO: merge all into one and remove that package`

### Internal

Inheret from clean architure folder

### Pkg

Inheret from clean architure folder

### Components

Folder with a-h/templ template files

### Handlers

Folder with front-end handlers routing

### Migrations

Sequence migration files to initialize postgres with goose

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
task pre-build
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
