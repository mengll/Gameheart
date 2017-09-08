package controller

```
Jump This function is for get the GameHeart message
Logs This function is save the request log
Downup This function is delete user when you leave
```
type GameHeart interface {
	Jump()
	Logs()
	Downup()
}


```
 This is config the GameHeart Message type
 Date string Heart jump time
 Ucid The game ucid anfanapi table
 Pid  the game id which game 
 Rid  the channel 
 Imei   android is imei ios is idfa 
```

type HeartMessageType struct {
	Date string "json:'date'",
	Ucid string "json:'ucid'",
	Pid  string "json:'pid'" ,
	Rid  string "json:'rid'" ,
	Imei string "json:'imei'"
}

```
  GameHeart is handle 

```

type  GameHeart struct {
	
}


```
 Get the oheart jump Data
```

func (self *GameHeart) Jump(msg *HeartMessageType){
	
}


```
 The Log handle
```

func (self *GameHeart)Logs(msg *HeartMessageType){
	
}

```
   Downup The player is leave
```
func (self *GameHeart)Downup(msg *HeartMessageType){
	
}




