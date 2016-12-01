package main

import ( 
"encoding/json"       
  "fmt"
  "net/http"
  "io"      
  "io/ioutil"     
  "time"
  "sync"
)

type Lang struct {
  Name string
  URL string
  Bytes int64
  Time time.Duration

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

func pfunc(ch chan Lang) {

  l := <-ch
      res1D := &Lang{
        Name:  l.Name,
        URL: l.URL ,
        Bytes: l.Bytes,
        Time: l.Time,
}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))
    fmt.Println("\n")

}

func crawl(pfunc func(chan Lang), lang Lang, w *sync.WaitGroup){
  name:=[3]string{"python","ruby","golang"}
  url:= [3]string{"https://www.python.org/", "https://www.ruby-lang.org/en/","https://golang.org/"}
  ch := make(chan Lang,3)
  var Cum_time time.Duration
  var tbytes int64
  Tnow := time.Now()
  for i,url_val := range url{
    w.Add(1)
    go routine(name[i], url_val, w, ch, &Cum_time, &tbytes)
  }

  w.Wait()
  total:= time.Since(Tnow)
  fmt.Println("\nTOTAL TIME con.:",total) 
  fmt.Println("TOTAL cumulative TIME : %v",Cum_time)
  fmt.Println("TOTAL bytes read form all 3 sites %v", tbytes)
}


func routine(name string, url string, w *sync.WaitGroup,ch chan Lang, Cum_time *time.Duration, tbytes *int64 ) { 
      var lang Lang
      lang.URL = url
      lang.Name = name
      lang.Bytes,lang.Time = retrieve(url)
      ch <-lang
      pfunc(ch)
      *Cum_time= *Cum_time+lang.Time
      *tbytes = *tbytes+lang.Bytes
      w.Done()

}

func main() {
  wg:= sync.WaitGroup{}
  var l Lang
  
  crawl(pfunc,l,&wg)
  wg.Wait() 
               
}
