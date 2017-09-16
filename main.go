package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/garyburd/redigo/redis"
)

var red *redis.Pool

func list(w http.ResponseWriter, r *http.Request) {
	re := red.Get()
	defer func() {
		err := re.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	keys, err := redis.Strings(re.Do("KEYS", "*"))
	if err != nil {
		log.Fatal(err)
	}

	res, err := json.Marshal(keys)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(res)
	if err != nil {
		log.Fatal(err)
	}
}

func save(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	email := r.FormValue("email")

	re := red.Get()
	defer func() {
		err := re.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	err := re.Send("SET", email, email)
	if err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello DevOps!")
}

func main() {
	red = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "redis:6379")
			if err != nil {
				log.Fatal(err)
			}
			return c, err
		},
	}

	http.HandleFunc("/", hello)
	http.HandleFunc("/list", list)
	http.HandleFunc("/save", save)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
