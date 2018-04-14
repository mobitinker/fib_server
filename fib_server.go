package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
    "strconv"
    "encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func TellAFib(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    digits, err := strconv.Atoi(ps.ByName("digits"))
    if err != nil {
        // handle error
        fmt.Println(err)
    }
    f := fib()
    result := make([]int, 0, digits)
    for i := 0; i <= digits; i++ {
        result = append(result, f())
    }
    sResult, _ := json.Marshal(result)
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, "%s", string(sResult))
}

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/fib/:digits", TellAFib)
    log.Fatal(http.ListenAndServe(":8080", router))
}

/*
// TODO:
Output json error instead of 404 page not found
Validate digits value
Add error code and message to result (0 = ok)
Add unit tests
*/
