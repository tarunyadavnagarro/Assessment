package main

import (        
  "fmt"
  "net/http"
  "io"      
  "io/ioutil"     
  "time"
  "encoding/json"
)

type Lang struct {
  Name string
  URL string
  Bytes int64
  Time time.Duration
}

func retrieve(uri string) ( int64, time.Duration) {  
  // this functions returs number of bytes load for specified URL and time taken for crawling it 
                           
  now := time.Now() // timer started
  resp1, err := http.Get(uri)
  if err != nil { 
    t := time.Since(now)    // timer ends for failed case       
    return 0, t                
  }
  num,_ :=io.Copy(ioutil.Discard, resp1.Body)
  t := time.Since(now)  // timer stops for successful read
  return num, t
}

func pfunc(l Lang) {
    // printe function
    res1B, _ := json.Marshal(l)
    fmt.Println(string(res1B))
  }

func crawl(pfunc func(Lang), lang Lang) {
  // driver function to assign particulars of specified URl obtained after crawling
  var ttime time.Duration
  url:= [3]string{"https://www.python.org/", "https://www.ruby-lang.org/en/","https://golang.org/"}
  var nameSet = [3]string{"Python","Ruby","Golang"}
  now := time.Now()
  for i,url_val := range url {
    lang.Name = nameSet[i]
    lang.URL = url_val
    lang.Bytes,lang.Time = retrieve(url_val)
    
    pfunc(lang)
  }
  ttime = time.Since(now)
  fmt.Println("total time to crawl 3 sites id %v", ttime) // total time taken to crawl all three sites

}

func main() {
  
  var l Lang

      crawl(pfunc,l)

  }