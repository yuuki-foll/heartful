package main

import (
	"log"
	"net/http"

	"context"

	"github.com/gin-gonic/gin"

	firebase "firebase.google.com/go"
	// "firebase.google.com/go/auth"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type Praises struct {
	comment string
}

func setup_firebase() *firestore.Client{
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
	return client
}

func creat_praises(comment string, career string, client *firestore.Client) {
	_, _, err := client.Collection(career).Add(context.Background(), map[string]interface{}{
		"comment": comment,
	})
	if err != nil {
		log.Fatalln(err)
	}
}


func main() {
	client := setup_firebase()

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

	// 褒めるページに遷移
	server.GET("/praise", func(ctx *gin.Context) {
		ctx.Redirect(302, "/templates/praise.tmpl")
	})

	server.POST("/register-praises", func(ctx *gin.Context) {
		comment := ctx.Query("praise_comment")
		career := ctx.Query("career")
		creat_praises(comment, career, client)
	})
	// 褒めるページに遷移
	server.GET("/comment_list", func(ctx *gin.Context) {
		ctx.Redirect(302, "/templates/comment_list.tmpl")
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
