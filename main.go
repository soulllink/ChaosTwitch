package main

import (
	"context"
	"log"
	"strings"

	twitch "github.com/gempir/go-twitch-irc"
	"github.com/tkanos/gonfig"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Configuration struct {
	Token   string
	Botname string
	Channel string
	Mongodb string
}

var ctx = context.Background()
var Client, _ = mongo.Connect(
	ctx,
	options.Client().ApplyURI("mongodb+srv://KEY"),
)
var wmsg = ""

func main() {
	//initconfig
	cfg := Configuration{}
	gonfig.GetConf("/private/config.json", &cfg)
	//init listfiles

	//mainbot
	client := twitch.NewClient(cfg.Botname, cfg.Token)

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		//chat output
		//log.Println(message.User.Name, message.Message)
		UserSave(message.User.Name)

		// //rules and replys
		if cmdtest(message.Message) {
			client.Say(message.Channel, custcmd(message.Message))
		}
		if strings.Contains(strings.ToLower(message.Message), "!ping") {
			client.Say(message.Channel, "PONG PogChamp")
		}
		if strings.Contains(message.Message, "!vanish") {
			s := "/timeout " + message.User.Name + " 1"
			client.Say(message.Channel, s)
		}
		if strings.Contains(message.Message, "!cycle") {
			client.Say(message.Channel, cycle(message.Message))
		}
		if strings.Contains(message.Message, "!video") {
			client.Say(message.Channel, videocmd(message.User.Name, message.Message))
		}
		if strings.Contains(message.Message, "!nextvideo") {
			client.Say(message.Channel, nextvideo(message.User.Name))
		}
		if strings.Contains(message.Message, "!cmd") {
			client.Say(message.Channel, cmd(message.User.Name, message.Message))
		}
	})

	client.Join(cfg.Channel)

	errb := client.Connect()
	if errb != nil {
		log.Fatal(errb)
	}
}
