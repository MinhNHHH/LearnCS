package exercises

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type Database map[string]dollars

func Ex711() {
	db := Database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.List)
	http.HandleFunc("/price", db.Price)
	http.HandleFunc("/update", db.Update)
	http.HandleFunc("/delete", db.Delete)
	http.HandleFunc("/add", db.Add)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (db Database) List(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
func (db Database) Price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
func (db *Database) Update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")
	priceNum, err := strconv.Atoi(priceStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %q\n", priceStr)
		return
	}

	_, ok := (*db)[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	(*db)[item] = dollars(priceNum)
	fmt.Fprintf(w, "%s\n", (*db)[item])
}

func (db *Database) Delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	_, ok := (*db)[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	delete(*db, item)
	fmt.Fprintf(w, "Delete %s success\n", item)
}

func (db *Database) Add(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")
	priceNum, err := strconv.Atoi(priceStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %q\n", priceStr)
		return
	}

	_, ok := (*db)[item]
	if !ok {
		(*db)[item] = dollars(priceNum)
		fmt.Fprintf(w, "%s\n", (*db)[item])
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item existed: %q\n", item)
		return
	}
}
