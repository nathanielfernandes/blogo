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

func (b *Blogo) RefreshPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Cors(w)

	auth_token := ps.ByName("auth_token")
	if auth_token == AUTHTOKEN {
		err := b.ReFetchPost(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Invalid Post ID", http.StatusNotFound)
			return
		}
	} else {
		http.Error(w, "Invalid Auth Token", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("Refreshed Post"))
}

func (b *Blogo) Refresh(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Cors(w)

	auth_token := ps.ByName("auth_token")
	if auth_token == AUTHTOKEN {
		b.FillCache()
	} else {
		http.Error(w, "Invalid Auth Token", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("Refreshed Posts"))
}

// func (b *Blogo) PostPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	Cors(w)

// 	auth_token := ps.ByName("auth_token")
// 	if auth_token == AUTHTOKEN {
// 		post, ok := b.GetPost(ps.ByName("id"))

// 		if !ok {
// 			http.Error(w, "Invalid Post ID", http.StatusNotFound)
// 			return
// 		}

// 		w.Header().Set("Content-Type", "text/html")
// 		w.Write(b.ToPage(post.Content))
// 	} else {
// 		http.Error(w, "Invalid Auth Token", http.StatusUnauthorized)
// 		return
// 	}
// }

func (b *Blogo) Post(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Cors(w)

	post, ok := b.GetPost(ps.ByName("id"))
	if !ok {
		http.Error(w, "Invalid Post ID", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	data, _ := json.Marshal(post)
	w.Write(data)
}

func (b *Blogo) PostContent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Cors(w)

	post, ok := b.GetPost(ps.ByName("id"))
	if !ok {
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
