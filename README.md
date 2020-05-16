# health-check-slack

## このアプリケーションについて

Slackを利用したサーバーの稼働状況を通知する、対話型システムです。

## 使用技術

- Golang

## 使用方法

### .envを設定

```
$ cp .env.sample .env
```

- SLACK_BOT_TOKEN
  - SLACK_BOTを利用するためのトークンを指定
- ROLE
  - Slack Botの権限を設定(master|node)
  - 1台目のBotはmasterを指定してください
  - 2台目以降はnodeを指定してください

### build

```
$ go build cmd/main.go
```

### Run

```
$ ./main
```

バックグラウンドで実行する場合(Linux用)

```
nohup ./main &
```