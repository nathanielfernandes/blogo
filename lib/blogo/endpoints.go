package blogo

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

var /* const */ AUTHTOKEN = os.Getenv("AUTHTOKEN")

func Cors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// func (b *Blogo) Publish(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	id := ps.ByName("id")
// 	auth_token := ps.ByName("auth_token")

// 	if auth_token == AUTHTOKEN {

// 	}
// }

func (b *Blogo) Post(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Cors(w)

	post, err := b.GetPost(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid Post ID", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	data, _ := json.Marshal(post)
	w.Write(data)
}

func (b *Blogo) PostContent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Cors(w)

	post, err := b.GetPost(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid Post ID", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(post.Content))
}

func (b *Blogo) Posts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Cors(w)
	w.Header().Set("Content-Type", "application/json")

	data, _ := json.Marshal(b.cache)
	w.Write(data)
}