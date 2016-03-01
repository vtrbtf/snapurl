package main

import (
  "os"
  "time"
  "net/http"
  "github.com/julienschmidt/httprouter" 
  "fmt"
  "io/ioutil"
  "strconv"
  "strings"
  "crypto/sha1"
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
  filePath := "_r/" + p.ByName("id")
  fileContentAsBytes, err := ioutil.ReadFile( filePath )
  fileContent := string(fileContentAsBytes)
  info, statErr := os.Stat(filePath)

  if err != nil || statErr != nil {
    panic(err)
  }

  split := strings.Split(fileContent, "#")
  sec, parseError := strconv.ParseInt(split[0], 16, 32)
  if parseError != nil {
    panic(parseError) 
  }
  URI := split[1]
  
  if (( info.ModTime().Unix() + sec ) - time.Now().Unix()) >= 0 { 
    http.Redirect(rw, r, URI, 302)
  } else {
    fmt.Fprintln(rw, "Timeout :) ")
  } 
}


func HashUrl( url string , unixTime int64 ) string {  
  h := sha1.New()
  h.Write([]byte(url + string(unixTime)))
  hash := h.Sum(nil) 
  return fmt.Sprintf("%x", hash)[:10]

}

func GoShort(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  
  url := r.FormValue("inputUrl")
  secs := r.FormValue("inputTime")
  secHEX, _ := strconv.Atoi(secs)   
  secondsHash := strconv.FormatInt(int64(secHEX), 16)

  urlBytes := []byte( secondsHash + "#" + r.FormValue("inputUrl"))

  unixTime := time.Now().Unix() 
  hash := HashUrl( url , unixTime )
  dir := "_r"
  path := dir + "/" + hash 

  ioerr := ioutil.WriteFile( path , urlBytes, 0644)

  if ioerr != nil  {
    panic( ioerr ) 
  }

  fmt.Fprintln( rw , "http://su.apps.vtrbtf.com/r/" + hash)
}

