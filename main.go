package main

import (
	"log"
	"net/http"

	"context"

	"github.com/gin-gonic/gin"

	firebase "firebase.google.com/go"
	// "firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

func setup_firebase() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("path/to/heartful-89dec-firebase-adminsdk-m23dq-cf25613a28.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
}

func main() {

	//サーバを準備
	server := gin.Default()

	// 静的ファイル(CSS等)の場所
	server.Static("/static", "./static")

	// 自動的にファイル返すよう設定
	server.StaticFS("/templates", http.Dir("templates"))

	// ルートなら /template/enter.html にリダイレクト
	server.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(302, "/templates/enter.tmpl")
	})

	// フォームの内容を受け取って挨拶する
	server.GET("/hello", func(ctx *gin.Context) {
		name := ctx.Query("name")
		server.LoadHTMLGlob("templates/main.tmpl")
		ctx.HTML(http.StatusOK, "main.tmpl", gin.H{
			"name": name,
		})
	})

	/* 	// Use the application default credentials
		ctx := context.Background()
		sa := option.WithCredentialsFile("path/to/serviceAccount.json") */

	// サーバーを起動
	err := server.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal("サーバー起動に失敗", err)
	}
}
