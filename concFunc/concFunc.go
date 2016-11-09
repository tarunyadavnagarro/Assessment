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
    //t := time.Since(now)           
    return 0               
  }
  num,_ :=io.Copy(ioutil.Discard, resp1.Body)
  t := time.Since(now)
  return num, t
}

func pfunc(l Lang) {
  res1D :=&Lang{
  Name:  l.Name,
  URL: l.URL ,
  Bytes: l.Bytes,
  Time: l.Time,
}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

}

func crawl(pfunc func(Lang), lang Lang, w *sync.WaitGroup){
  name:=[3]string{"python","ruby","golang"}
  url:= [3]string{"https://www.python.org/", "https://www.ruby-lang.org/en/","https://golang.org/"}
  
  var lngarr [3]Lang
  Tnow := time.Now()
      w.Add(1)
      go func(){
        lngarr[0].Bytes,lngarr[0].Time = retrieve("https://www.python.org/")
        w.Done()
      }()

      w.Add(1)
      go func(){
        lngarr[1].Bytes,lngarr[1].Time = retrieve("https://www.ruby-lang.org/en/")
        w.Done()
      }()

      w.Add(1)
      go func(){
        lngarr[2].Bytes,lngarr[2].Time = retrieve("https://golang.org/")
        w.Done()
      }()

  w.Wait()
  total:= time.Since(Tnow)
  lngarr[0].URL = url[0]
  lngarr[0].Name=name[0]
  lngarr[1].URL = url[1]
  lngarr[1].Name=name[1]
  lngarr[2].URL = url[2]
  lngarr[2].Name=name[2]
  pfunc(lngarr[0])
  pfunc(lngarr[1])
  pfunc(lngarr[2])
  fmt.Println("\nTOTAL TIME con.: ",total)
}

func main() {
  wg:= sync.WaitGroup{}
  var l Lang
  crawl(pfunc,l,&wg) 
  wg.Wait() 
  fmt.Println("DONE ")
               
}
