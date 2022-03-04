package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var workers = 100

func main() {
	wg := sync.WaitGroup{}
	for i:=0 ; i< workers; i++{
		wg.Add(1)
		go func(){
			count := 0
			for {
				if count > 20{
					break
				}
				resp, err := http.Get("http://127.0.0.1:8080/dispatch?customer=123")
				if err != nil {
					panic(err.Error())
				}
				defer resp.Body.Close()
				fmt.Println(resp.StatusCode)
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					print(err)
				}
				fmt.Print(string(body))
				count ++
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(25)))
			}
			
			wg.Done()
		}()
	}

	wg.Wait()
}