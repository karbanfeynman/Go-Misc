package main

import (
//	"database/sql"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"runtime"
)

type Users struct {
	Id   int
	Name string
}

type Image struct {
	Name string
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func SetupDB() *sql.DB {
	db, err := sql.Open("sqlite3", "/home/karban/SQLite/user.db")
	PanicIf(err)

	return db
}

func handler_Root(r render.Render) {
	rows, err := db.Query("SELECT ID, URL FROM href")
	PanicIf(err)
	defer rows.Close()

	users := []Users{}
	for rows.Next() {
		u := Users{}
		err := rows.Scan(&u.Id, &u.Name)
		PanicIf(err)
		users = append(users, u)
	}
	r.HTML(200, "index", nil)

}

func handler(r render.Render, params martini.Params) {
	switch params["name"] {
	case "index":
		images := []Image{}
		files, _ := ioutil.ReadDir("/home/karban/Go/Go/templates/html5/IMAGE")
		for _, f := range files {
			i := Image{}
			i.Name = f.Name()
			images = append(images, i)
			fmt.Printf("[Karban]name: %s\n", i.Name)
		}
		r.HTML(200, params["name"], images)
	default:
		r.HTML(200, params["name"], nil)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	m := martini.Classic()
	m.Map(SetupDB())

	/* set host and port */
	m.Use(render.Renderer(render.Options{
		Directory: "templates/html5",
		Layout:    "layout",
	}))
	m.Use(martini.Static("templates/html5/"))

	m.Get("/", handler_Root)
	m.Get("/:name", handler)

	m.Run()
}
