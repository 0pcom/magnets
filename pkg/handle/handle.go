// handle.go //

package handle

import (
"fmt"
"net/http"
"encoding/json"
"os"
"math/rand"
"time"
"github.com/bitfield/script"
)

//  haltingstate.net/
func Haltingstate(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w,"<!doctype html><html lang=en><head></head><body><div style='width:50%'>")
     file, err := os.ReadFile("result.json")
     if err != nil {
       fmt.Fprintf(w,"error reading file: %v", err)
       os.Exit(1)
     }

     data := TgHist{}

     err = json.Unmarshal([]byte(file), &data)
     if err != nil {
       fmt.Fprintf(w,"error unmarshalling json: %v", err)
       os.Exit(1)
     }
     rand.Seed(time.Now().UnixNano())
     min := 0
     max := len(data.Messages)

     i := rand.Intn(max - min) + min
     //	for i := 0; i < len(data.Messages); i++ {
     var ok bool
     for !ok {
     if data.Messages[i].From == "Synth" || data.Messages[i].From == "Skycoin" {
       if data.Messages[i].Text != "" {
         ok = true
       }
     }
     if !ok {
       i = rand.Intn(max - min) + min
     }
     }

     shcmd :=`/usr/bin/bash -c`
     cmd := fmt.Sprintf(`%s "_var=\"%s\" ; _var=\"${_var#*\[map\[text\:}\" ; _var=\"${_var/type:link]/\n}\" ; _var=\"${_var%]}\" ; echo ${_var}"`, shcmd, fmt.Sprintf("%s", data.Messages[i].Text))
     res, err := script.Exec(cmd).String()
     if err != nil {
     fmt.Fprintf(w,"an error during script.Exec:\n<br> %v\n<br>", err)
     fmt.Fprintf(w,"try reloading\n<br>")
     fmt.Fprintf(w, "<img src='img/haltingstate.jpg'>")
     } else {
     fmt.Fprintf(w,"<a href=/ title='https://t.me/Skycoin/%d'>https://t.me/Skycoin/%d\n</a><br>", data.Messages[i].ID, data.Messages[i].ID)
     fmt.Fprintf(w,"%s\n<br>", data.Messages[i].Date)
     fmt.Fprintf(w,"\n<br>%s\n<br>", res)
     fmt.Fprintf(w,"-%s\n<br>\n<br>", data.Messages[i].From)
     }
     fmt.Fprintf(w,"</div></body></html>")


}




type TgHist struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	ID       int    `json:"id"`
	Messages []struct {
		ID           int    `json:"id"`
		Type         string `json:"type"`
		Date         string `json:"date"`
		DateUnixtime string `json:"date_unixtime"`
		From         string `json:"from,omitempty"`
		FromID       string `json:"from_id,omitempty"`
		Author       string `json:"author,omitempty"`
		Text         any `json:"text"`
		TextEntities []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"text_entities"`
		Photo            string `json:"photo,omitempty"`
		Width            int    `json:"width,omitempty"`
		Height           int    `json:"height,omitempty"`
		Edited           string `json:"edited,omitempty"`
		EditedUnixtime   string `json:"edited_unixtime,omitempty"`
		ReplyToMessageID int    `json:"reply_to_message_id,omitempty"`
		File             string `json:"file,omitempty"`
		MimeType         string `json:"mime_type,omitempty"`
		ForwardedFrom    string `json:"forwarded_from,omitempty"`
		Thumbnail        string `json:"thumbnail,omitempty"`
		MediaType        string `json:"media_type,omitempty"`
		DurationSeconds  int    `json:"duration_seconds,omitempty"`
		Actor            string `json:"actor,omitempty"`
		ActorID          string `json:"actor_id,omitempty"`
		Action           string `json:"action,omitempty"`
		MessageID        int    `json:"message_id,omitempty"`
	} `json:"messages"`
}
