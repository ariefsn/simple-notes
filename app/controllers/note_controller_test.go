package controllers

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eaciit/toolkit"
)

var N = NewNoteController()

func TestNoteGetAll(t *testing.T) {
	// var k *knot.WebContext
	// req, err := http.NewRequest("GET", "/api/getall", nil)
	// if err != nil {
	// 	t.Fatalf("Test failed! Error: %s", err.Error())
	// }

	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		// k.Writer = w
		// k.Request = r
		w.Write([]byte("damen"))
		// N.GetAll(k)
	}

	// rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerFunc)

	// handler.ServeHTTP(rr, req)
	ts := httptest.NewServer(handler)
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", greeting)

	t.Errorf("kwkwkwkwk")
	// status := rr.Code
	// if status != http.StatusOK {
	// 	t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	// }
}

func TestExampleResponseRecorder(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" {
			res := struct {
				status  string
				message string
				data    toolkit.M
			}{}
			res.status = "ok"
			res.message = "success"
			res.data = toolkit.M{"name": "ayik"}

			w.Write(toolkit.Jsonify(&res))
			// io.WriteString(w, "<html><body>Hello World! GET Method!</body></html>")
		} else if r.Method == "POST" {
			io.WriteString(w, "<html><body>Hello World! POST Method!</body></html>")
		} else if r.Method == "PUT" {
			io.WriteString(w, "<html><body>Hello World! PUT Method!</body></html>")
		} else {
			io.WriteString(w, "<html><body>Hello World! OTHERS Method!</body></html>")
		}
	}

	req := httptest.NewRequest("", "https://jsonplaceholder.typicode.com/users", nil)
	w := httptest.NewRecorder()
	// handler(w, req)
	// xHandler := http.HandleFunc()
	xhandler := http.HandlerFunc(handler)
	xhandler.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body.String())

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
	t.Errorf("Stop")
	// Output:
	// 200
	// text/html; charset=utf-8
	// <html><body>Hello World!</body></html>
}
