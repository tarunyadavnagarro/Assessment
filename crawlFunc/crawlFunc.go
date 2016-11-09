package main

import (        
  "fmt"
  "net/http"
  "io"      
  "io/ioutil"     
  "time"
)

type Lang struct {
  Name string
  URL string
  bytes int64
  time time.Duration

}
func retrieve(uri string) ( int64, time.Duration) {  
                           
  now := time.Now()
  resp1, err := http.Get(uri)
  if err != nil { 
    t := time.Since(now)           
    return 0, t                
  }
  num,_ :=io.Copy(ioutil.Discard, resp1.Body)
  t := time.Since(now)
  return num, t
}

func pfunc(l Lang) {
    fmt.Println("name :   %v",l.Name)
    fmt.Println("URL is  %v:",l.URL)
      fmt.Println("number of bytes loaded are %v",l.bytes)
      fmt.Println("time taken is : %v", l.time)
  }

func crawl(pfunc func(Lang), lang Lang){
  var ttime time.Duration
  url:= [3]string{"https://www.python.org/", "https://www.ruby-lang.org/en/","https://golang.org/"}
  var nameSet = [3]string{"Python","Ruby","Golang"}
  now := time.Now()
  for i,url_val := range url {
    lang.Name = nameSet[i]
    lang.URL = url_val
    lang.bytes,lang.time = retrieve(url_val)
    
    pfunc(lang)

  }
  ttime = time.Since(now)
  fmt.Println("total time to crawl 3 sites id %v", ttime)

}

func main() {
  
  var l Lang

      crawl(pfunc,l)

  }
