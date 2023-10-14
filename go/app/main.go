package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

var rdb *redis.Client
var ctx = context.Background()

func main() {
	// 環境変数から設定を読み込む
	goPort := os.Getenv("GO_PORT")
	redisIP := os.Getenv("REDIS_IP")
	redisPort := os.Getenv("REDIS_PORT")

	// Redisの設定
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisIP, redisPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// ハンドラの設定
	http.HandleFunc("/", setKeyValue)

	// サーバの起動
	log.Printf("Starting server on port %s...\n", goPort)
	if err := http.ListenAndServe(":"+goPort, nil); err != nil {
		log.Fatal(err)
	}
}

// setKeyValue はClientからリクエストされたKeyとValueをRedisに保存する
func setKeyValue(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")

	if key == "" || value == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Both key and value parameters are required.")
		return
	}

	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Failed to set key-value in Redis.")
		return
	}

	fmt.Fprintf(w, "Set key=%s with value=%s successfully!", key, value)
}
