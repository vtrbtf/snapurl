package main

import (
  "os"
  "net/http"
  "github.com/julienschmidt/httprouter" 
  "fmt"
  "io/ioutil"
  "encoding/base64"
)


func main() {
  r := httprouter.New()

  r.Handler("GET", "/", http.FileServer(http.Dir("public")))

  r.POST("/goshort", GoShort)
  r.GET("/r/:id", Redirect)

  http.ListenAndServe(":" + GetPort(), r)
}

func GetPort() string {  
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }
  return port
}

func Redirect( rw http.ResponseWriter, r *http.Request, p httprouter.Params )  {
  dat, err := ioutil.ReadFile("_r/" + p.ByName("id"))

  if err != nil {
    panic(err)
  }

  http.Redirect(rw, r, string(dat[:]), 301)
}


func HashUrl( url string ) string {  
  return base64.StdEncoding.EncodeToString([]byte(url))
}

func GoShort(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  
  url := r.FormValue("inputUrl")
  urlBytes := []byte(r.FormValue("inputUrl"))

  hash := HashUrl( url )
  dir := "_r"
  path := dir + "/" + hash 

  err := ioutil.WriteFile( path , urlBytes, 0644)

  if err != nil  {
    panic( err ) 
  }

  fmt.Fprintln( rw , path)
}

