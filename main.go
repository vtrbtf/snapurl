package main

import (
 // "os"
  "net/http"
  "github.com/julienschmidt/httprouter" 
  "fmt"
  "io/ioutil"

)


func main() {
  r := httprouter.New()

  r.Handler("GET", "/", http.FileServer(http.Dir("public")))

  r.POST("/goshort", GoShort)
  //r.GET("/r/:id", Redirect)

  http.ListenAndServe(":" + GetPort(), r)
}

func GetPort() string {  
 return "8080" 
}


func HashUrl( url string ) string {  
 return "123"
}

func GoShort(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  
  url := r.FormValue("inputUrl")
  urlBytes := []byte(r.FormValue("inputUrl"))
  fmt.Println(url)
  err := ioutil.WriteFile("_r/" + HashUrl( url ), urlBytes, 0644)

  if err != nil  {
    panic( err ) 
  }

  fmt.Fprintln( rw , "Created :) ")
}

