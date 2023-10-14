# 使用するベースイメージを指定
FROM golang:latest

# 作業ディレクトリの設定
WORKDIR /app

RUN apt-get update -y &&\
    apt-get upgrade -y &&\
    apt-get install redis-tools -y

# main.goをContainer環境へCOPY
COPY app/main.go .

# 必要なモジュールを取得
RUN go mod init app
RUN go get github.com/redis/go-redis/v9
RUN go get golang.org/x/net/context

# アプリケーションをビルド
# RUN go build -o main .
RUN go build -o RegisterToRedis main.go

# アプリケーションの実行
# CMD ["./main"]
CMD ["./RegisterToRedis"]
