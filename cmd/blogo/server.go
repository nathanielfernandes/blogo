package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	b "github.com/nathanielfernandes/blogo/lib/blogo"
)

func main() {
	blogo := b.NewBlogo()
	blogo.FillCache()

	router := httprouter.New()
	router.GET("/post/json/:id", blogo.Post)
	router.GET("/post/content/:id", blogo.PostContent)
	router.GET("/posts", blogo.Posts)

	router.GET("/refresh/:auth_token", blogo.Refresh)
	router.GET("/refreshpost/:auth_token/:id", blogo.RefreshPost)
	// router.GET("/page/:auth_token/:id", blogo.PostPage)

	// router.POST("/publish/:id/:auth_token", blogo.Publish)
	// router.GET("/ws")

	fmt.Printf("Blogo\nListening on port 80\n")
	if err := http.ListenAndServe("0.0.0.0:80", router); err != nil {
		log.Fatal(err)
	}
}
