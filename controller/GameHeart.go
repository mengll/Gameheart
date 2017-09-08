package controller

import (
	"fmt"
	"net/http"
)

///*
//Jump This function is for get the GameHeart message
//Logs This function is save the request log
//Downup This function is delete user when you leave
//```
type GameHeartInterface interface {
	Jump()
	Logs()
	Downup()
}

//```
// This is config the GameHeart Message type
// Date string Heart jump time
// Ucid The game ucid anfanapi table
// Pid  the game id which game
// Rid  the channel
// Imei   android is imei ios is idfa
//```

type HeartMessageType struct {
	Date string `json:"date"`
	Ucid string `json:"ucid"`
	Pid  string `json:"pid"`
	Rid  string `json:"rid"`
	Imei string `json:"imei"`
}

//```
//  GameHeart is handle

type mess struct {
	CloseTime int    `json:"close_time"`
	Content   string `json:"content"`
}

type Dat struct {
	Broadcast []mess `json:"broadcast"`
	OnlineNum int    `json:"online_num"`
	WinRate   int    `json:"win_rate"`
}

type messData struct {
	Code int    `json:"code"`
	Data Dat    `json:"data"`
	Msg  string `json:"msg"`
}

//the back data type

func hear(w http.ResponseWriter, r *http.Request) {

	method := r.Method
	if strings.ToLower(method) != "post" {
		con_m := mess{}
		con_m.CloseTime = 3
		con_m.Content = "colose Dat"
		hj := []mess{con_m}

		metdat := Dat{}
		metdat.OnlineNum = 2
		metdat.WinRate = 2
		metdat.Broadcast = hj
		dat := messData{}
		dat.Code = 1
		dat.Msg = "error"
		dat.Data = metdat

		cont, err := json.Marshal(dat)
		if err != nil {
			return
		}

		w.Write(cont)
		fmt.Println("one")
		return
	}
}

//```

type GameHeart struct {
}

//```
// Get the oheart jump Data
//```

func (self *GameHeart) Jump(msg *HeartMessageType) {

}

//```
// The Log handle
//```

func (self *GameHeart) Logs(msg *HeartMessageType) {

}

//```
//   Downup The player is leave
//```
func (self *GameHeart) Downup(msg *HeartMessageType) {

}

//```
//GameHeart http request
//```
func GameHeartRequest(w http.ResponseWriter, r *http.Request) {
	//create The int game data
	fmt.Println("This is Game heart")
}
