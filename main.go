package main

import (
	"encoding/json"
	"fmt"
	logic2 "github.com/Jason-cqtan/shorturl/logic"
	"net/http"
)

const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`

var store = logic2.NewURLStore()

type user struct {
	name string
	age  int
}

var port = "8099"

func fooHandler(w http.ResponseWriter, r *http.Request) {
	obj, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, string(obj))
}

func main() {
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(":"+port, nil)

}

func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	url := store.Get(key)
	if url == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func Add(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	fmt.Println(url)
	if url == "" {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, AddForm)
		return
	}
	key := store.Put(url)
	fmt.Fprintf(w, "http://localhost:%s/%s", port, key)
}
