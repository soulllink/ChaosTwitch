package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type keyvalue struct {
	Name    string `bson:"Name"`
	Value   string `bson:"Value"`
	Counter int    `bson:"Counter"`
}

type userbase struct {
	Name  string `bson:"Name"`
	Coins int    `bson:"Coins"`
}

type quebase struct {
	Vid   int    `bson:"Vid"`
	Uname string `bson:"Uname"`
	Url   string `bson:"Url"`
}

var currentsavevid int

//querry
//TODO UPDATE COUNTER
func dataq(name string) keyvalue {
	var out keyvalue
	i := name
	filter := bson.M{"Name": i}
	db := Client.Database("Nothing").Collection("databaseC")
	db.FindOne(context.Background(), filter).Decode(&out)
	if out.Name == "" || out.Name == " " {
		return out
	} else {
		update := bson.D{{"$inc", bson.M{"Counter": 1}}}
		db.UpdateOne(context.TODO(), filter, update)
		db.FindOne(context.Background(), filter).Decode(&out)
		return out
	}
}

//save

func datas(name string, value string) {

	kv := &keyvalue{name, value, 0}
	db := Client.Database("Nothing").Collection("databaseC")
	db.InsertOne(context.Background(), kv)
	return

}

//Userbase
func Userq(name string) userbase {
	var out userbase
	filter := bson.M{"Name": name}
	db := Client.Database("Nothing").Collection("dataUser")
	db.FindOne(context.Background(), filter).Decode(&out)

	return out
}

//save
func UserSave(name string) {
	i := Userq(name)
	if i.Name == "" || i.Name == " " {

		ud := &userbase{name, 10}
		db := Client.Database("Nothing").Collection("dataUser")
		db.InsertOne(context.Background(), ud)
		return

	} else {
		filter := bson.M{"Name": name}
		update := bson.D{{"$inc", bson.M{"Coins": 1}}}
		db := Client.Database("Nothing").Collection("dataUser")
		db.UpdateOne(context.TODO(), filter, update)
		return
	}
}

//Que
func qhelp(id int) quebase {
	var out quebase
	filter := bson.M{"Vid": id}
	db := Client.Database("Nothing").Collection("Video")
	db.FindOne(context.Background(), filter).Decode(&out)
	// log.Println(out)
	return out
}

func quebasehelp() int {
	if currentvideoid == 0 {
		return 1
	} else {
		currentsavevid++
		return currentsavevid
	}

}

func quesave(Uname string, Url string) {
	Vid := quebasehelp()
	qb := &quebase{Vid, Uname, Url}
	db := Client.Database("Nothing").Collection("Video")
	db.InsertOne(context.Background(), qb)
}
