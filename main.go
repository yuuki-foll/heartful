package main

import (
	"fmt"
	"log"
	"net/http"
	"math/rand"
	"context"

	"github.com/gin-gonic/gin"

	firebase "firebase.google.com/go"
	// "firebase.google.com/go/auth"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type Praises struct {
	comment string
}

func setup_firebase() (context.Context, *firestore.Client) {
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
	return ctx, client
}

func creat_praises(comment string, career string, ctx context.Context, client *firestore.Client) {
	_, _, err := client.Collection(career).Add(ctx, map[string]interface{}{
		"comment": comment,
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func random_pick(array []string) []string {
	praises := make([]string, 0)
	idx_slice := make([]int, 0)
	for {
		idx := rand.Intn(len(array))
		add := true
		for _, v := range idx_slice {
			if idx == v {
				add = false
			}
		}
		if add {
			idx_slice = append(idx_slice, idx)
		}

		if len(idx_slice) == 10 {
			break
		}
	}

	for _, v := range idx_slice {
		praises = append(praises, array[v])
	}
	return praises
}

// 誉め言葉をfirestoreから取得する
func read_praises(career string, ctx context.Context, client *firestore.Client) []string {
	iter := client.Collection(career).Documents(ctx)
	praises := make([]string, 0)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		p_map := doc.Data()
		comment := p_map["comment"].(string)
		praises = append(praises, comment)
		if len(praises) > 10 {
			praises = random_pick(praises)
		}
		// fmt.Println(doc.Data())
	}
	fmt.Println(praises)
	return praises
}

func main() {
	c, client := setup_firebase()

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

	// 褒められるページに遷移
	server.GET("/user_info", func(ctx *gin.Context) {
		ctx.Redirect(302, "/templates/praise_get.tmpl")
	})

	server.GET("/happy", func(ctx *gin.Context) {
		ctx.Redirect(302, "/templates/happy.tmpl")
	})
	server.POST("/register_praises", func(ctx *gin.Context) {
		comment := ctx.PostForm("praise_comment")
		career := ctx.PostForm("career")
		fmt.Println(comment)
		fmt.Println(career)
		creat_praises(comment, career, c, client)
		ctx.Redirect(302, "/templates/enter.tmpl")
	})

	// 誉め言葉をjsonで送る
	server.POST("/get_praises", func(ctx *gin.Context) {
		career := ctx.PostForm("career")
		praises := read_praises(career, c, client)
		ctx.JSON(200, gin.H{
			"comment": praises,
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
	defer client.Close()
}
