package main

// import (
//   "fmt"
//   "strings"
//   "net/http"
//   "io/ioutil"
// )

// func main() {

//   url := "localhost:8080/items"
//   method := "GET"

//   payload := strings.NewReader(`{`+"
// "+`
//     "title": "item 2",`+"
// "+`
//     "description": "description 2",`+"
// "+`
//     "url": "item2 url"`+"
// "+`
// }`)

//   client := &http.Client {
//   }
//   req, err := http.NewRequest(method, url, payload)

//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   req.Header.Add("Authorization", "Basic VXNlcjoxMjM0NTY3OA==")
//   req.Header.Add("Content-Type", "application/json")

//   res, err := client.Do(req)
//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   defer res.Body.Close()

//   body, err := ioutil.ReadAll(res.Body)
//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   fmt.Println(string(body))
// }
