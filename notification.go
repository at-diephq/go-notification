/**
Author: dung138 (blog.daoanhdung.com)
**/
package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
)

/** Data Structure **/
type RepsonseData struct {
  Response string
  Message  string
}

var (
  rs                         = RepsonseData{Response: "", Message: ""}
  config                     = &Config{}
  response                   string
)

type Config struct {
  SMTP_EMAIL      string
  SMTP_PASSWORD   string
  SMTP_SERVER     string
  SMTP_PORT       int
  SLACK_WEBHOOK   string
  CHATWORK_TOKEN  string
  ROOM_ID         string
  LISTEN          string
}

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}

func loadConfig() *Config {
  conf := Config{}
  content, e := ioutil.ReadFile("./config.json")
  checkErr(e)
  err := json.Unmarshal(content, &conf)
  checkErr(err)
  return &conf
}

func wirteJson(data RepsonseData, w http.ResponseWriter) {
  js, err := json.Marshal(data)
  checkErr(err)
  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func parseBody(resp *http.Response) []byte {
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  checkErr(err)

  return body
}

func doSendChatwork(body string) []byte {
  client := &http.Client{}
  req, requestErr := http.NewRequest("POST", "https://api.chatwork.com/v2/rooms/"+config.ROOM_ID+"/messages?body="+body, nil)
  checkErr(requestErr)
  req.Header.Add("X-ChatWorkToken", config.CHATWORK_TOKEN)
  resp, responseErr := client.Do(req)
  checkErr(responseErr)
  
  return parseBody(resp)
}

func notifyLogin(w http.ResponseWriter, r *http.Request) {
  response := doSendChatwork("Warning%21%20Logged")
  fmt.Println(response)
}

func main() {
  config = loadConfig()
  http.HandleFunc("/notify/login", notifyLogin)
  err := http.ListenAndServe(config.LISTEN, nil)
  checkErr(err)
}
