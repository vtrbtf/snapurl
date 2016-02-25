package main

import (
  "os"
  "time"
  "net/http"
  "github.com/julienschmidt/httprouter" 
  "fmt"
  "io/ioutil"
  "encoding/base64"
  "strconv"

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

  split := strings.Split(dat, "#")
  sec, _ := strconv.ParseInt(split[0], 16, 32)
  URI := split[1]
  
  if  time.Now().Unix()
  http.Redirect(rw, r, string(dat[:]), 301)
}


func HashUrl( url string , unixTime int64 ) string {  
  return base64.StdEncoding.EncodeToString([]byte(url + string(unixTime)))
}

func GoShort(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  
  url := r.FormValue("inputUrl")
  secs := r.FormValue("inputTime")
  secHEX, _ := strconv.Atoi(secs)   
  urlBytes := []byte(strconv.FormatInt(int64(secHEX), 16) + "#" + r.FormValue("inputUrl"))

  unixTime := time.Now().Unix() 
  hash := HashUrl( url , unixTime )
  dir := "_r"
  path := dir + "/" + hash 

  ioerr := ioutil.WriteFile( path , urlBytes, 0644)

  if ioerr != nil  {
    panic( ioerr ) 
  }

  fmt.Fprintln( rw , path)
}

