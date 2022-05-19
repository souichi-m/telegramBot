package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const telegramApi string = "https://api.telegram.org/bot"

func main() {
	type chat struct {
		Id int
	}

	type photo struct {
		FileId string `json:"file_id"`
	}

	type sticker struct {
		FileId string `json:"file_id"`
	}

	type message struct {
		Chat    chat
		Text    string
		Photo   []photo
		Sticker sticker
	}

	type update struct {
		Message  message
		UpdateId int `json:"update_id"`
	}

	type myStruct struct {
		Result []update
	}

	var offset int

	for {
		resp, err := http.Get(telegramApi + "5291797564:AAHIClHrMd8AUyrfAb-FAddWUCCeki709xQ/getUpdates?offset=" + fmt.Sprint(offset))
		if err != nil {
			// handle error
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)

		var s myStruct

		if err := json.Unmarshal(body, &s); err != nil {
			fmt.Println(err)
			return
		}

		log.Printf("%v", s) // {John Connor 35 true [15 11 37]}

		for _, v := range s.Result {
			if v.Message.Text != "" {
				http.Get(telegramApi +
					"5291797564:AAHIClHrMd8AUyrfAb-FAddWUCCeki709xQ/sendMessage?chat_id=1079276889&text=" +
					v.Message.Text)
			} else if v.Message.Photo != nil {
				http.Get(telegramApi +
					"5291797564:AAHIClHrMd8AUyrfAb-FAddWUCCeki709xQ/sendPhoto?chat_id=1079276889&photo=" +
					v.Message.Photo[len(v.Message.Photo)-1].FileId)
			} else if v.Message.Sticker.FileId != "" {
				http.Get(telegramApi +
					"5291797564:AAHIClHrMd8AUyrfAb-FAddWUCCeki709xQ/sendSticker?chat_id=1079276889&sticker=" +
					v.Message.Sticker.FileId)
			} else {
				log.Printf("%s", "Unknown type message")
			}
			offset = v.UpdateId + 1
		}
		time.Sleep(5 * time.Second)
	}
}
