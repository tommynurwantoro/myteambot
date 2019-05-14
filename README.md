# Act BL Bot

Telegram bot for some activities in ACT (my team)

## Owner

@tommynurwantoro

## Onboarding and Development Guide

### Prerequisite
- Git
- Go 1.11 or later
- Go Dep 0.5 or later
- MySQL 5.7
- Redis 4.0.11 or later

### Installation

- Install Git  
  See [Git Installation](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

- Install Go (Golang)  
  See [Golang Installation](https://golang.org/doc/install)

- Install Go Dep  
  See [Dep Installation](https://golang.github.io/dep/docs/installation.html)

- Install MySQL  
  See [MySQL Installation](https://www.mysql.com/downloads/)
  
- Install Redis  
  See [Redis Installation](https://redis.io/topics/quickstart)

- Clone this repo in your local at `$GOPATH/src/github.com/bot`
  If you have not set your GOPATH, set it using [this](https://golang.org/doc/code.html#GOPATH) guide.
  If you don't have directory `src`, `github.com`, or `bot` in your GOPATH, please make them.

  ```sh
  git clone git@github.com:tommynurwantoro/act-bl-bot.git
  ```

- Go to act_bl_bot directory, then sync the vendor file

  ```sh
  cd $GOPATH/src/github.com/bot/act-bl-bot
  make dep
  ```

- Copy env.sample and if necessary, modify the env value(s)

  ```sh
  cp env.sample .env
  ```

- Copy db/config.yml.sample and if necessary, modify the config file

  ```sh
  cp db/config.yml.sample db/config.yml
  ```

- Install Bundler
  
  ```sh
  gem install bundler
  ```

- Prepare database

  ```sh
  bundle install
  rake db:create db:migrate
  ```

- Run Bot

  ```sh
  make run
  ```
