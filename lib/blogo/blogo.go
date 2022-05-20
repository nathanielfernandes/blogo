package blogo

import (
	"fmt"
	"time"

	"github.com/bep/debounce"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	m "github.com/nathanielfernandes/blogo/lib/mongo"
	rl "github.com/nathanielfernandes/rl"
)

type Blogo struct {
	db    *m.BlogoMongo
	cache map[string]m.Post

	get_rlm    *rl.RatelimitManager
	view_rlm   *rl.RatelimitManager
	cookie_rlm *rl.RatelimitManager

	debounce func(f func())
}

func NewBlogo() Blogo {
	return Blogo{
		db:    m.NewBlogoMongo(),
		cache: map[string]m.Post{},

		get_rlm:    rl.NewRatelimitManager(5, 1000),
		view_rlm:   rl.NewRatelimitManager(1, 86400000),
		cookie_rlm: rl.NewRatelimitManager(20, 1000),

		debounce: debounce.New(10 * time.Second),
	}
}

func (b *Blogo) NewPost(id, title, hook, content string, tags []string) error {
	fmt.Println("SUBMITTING NEW POST")
	err := b.db.NewPost(id, title, hook, content, tags)
	if err != nil {
		fmt.Println("FAILED TO POST: EXISTS")
		return err
	}

	post, err := b.db.GetPost(id)
	if err != nil {
		fmt.Println("FAILED TO POST: CAN'T GET")
		return err
	}

	fmt.Println("POST SUBMITTED")
	post.Content = b.ToHTML(post.Content)
	b.cache[id] = post
	return nil
}

func (b *Blogo) GetPost(id string) (m.Post, bool) {
	// if post, ok := b.cache[id]; ok {
	// 	return post, true
	// }

	// fmt.Println("FETCHING POST: " + id)
	// post, err := b.db.GetPost(id)
	// if err != nil {
	// 	return m.Post{}, err
	// }

	// post.Content = b.ToHTML(post.Content)

	// b.cache[id] = post
	// return post, err

	post, ok := b.cache[id]
	return post, ok
}

func (b *Blogo) ReFetchPost(id string) error {
	post, err := b.db.GetPost(id)
	if err != nil {
		return err
	}

	post.Content = b.ToHTML(post.Content)
	b.cache[post.ID] = post

	return nil
}

func (b *Blogo) FillCache() {
	fmt.Println("FILLING CACHE")
	posts, err := b.db.GetAllPosts()
	if err != nil {
		fmt.Println("FAILED TO GET ALL POSTS")
		return
	}

	for _, post := range posts {
		post.Content = b.ToHTML(post.Content)
		b.cache[post.ID] = post
	}
}

func (b *Blogo) ToHTML(content string) string {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)
	return string(markdown.ToHTML([]byte(content), p, nil))
}

// func (b *Blogo) ToPage(content string) []byte {
// 	htmlFlags := html.CommonFlags | html.HrefTargetBlank
// 	opts := html.RendererOptions{Flags: htmlFlags}
// 	renderer := html.NewRenderer(opts)

// 	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
// 	p := parser.NewWithExtensions(extensions)

// 	return markdown.ToHTML([]byte(content), p, renderer)
// }
