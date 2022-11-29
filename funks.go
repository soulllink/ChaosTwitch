package main

import (
	"math/rand"
	"strconv"
	"strings"
	"unicode"
)

var currentvideoid int
var mediafl []string

func random() bool {
	return rand.Intn(2) == 0
}

func cycle(ctx string) string {
	s := ctx
	s = strings.Replace(s, "!cycle", "", -1)
	f := func(r rune) rune {
		if random() {
			return unicode.ToUpper(r)
		} else {
			return unicode.ToLower(r)
		}
	}
	out := strings.Map(f, s)
	return out
}

//video
func videocmd(name string, ctx string) string {
	i := ctx
	i = strings.Replace(i, "!video", "", -1)
	//NEEED OT FIX THIS
	if i == "" || i == " " {
		if currentvideoid == 0 {
			v := qhelp(1)
			out := "Видео отправил: " + v.Uname + " ссылка: " + v.Url
			return out
		} else {
			v := qhelp(currentvideoid)
			out := "Видео отправил: " + v.Uname + " ссылка: " + v.Url
			return out
		}

	} else {
		quesave(name, i)
		out := "Добавил в очередь"
		return out
	}
}

func nextvideo(name string) string {
	if name == "teoreez" {
		i := currentvideoid + 1
		v := qhelp(i)
		if v.Uname == "" || v.Uname == " " {
			return "Больше ничего нет."
		} else {
			currentvideoid++
			ctx := ""
			out := videocmd(name, ctx)
			return out
		}
	} else {
		return "Менять видео может только NAME"
	}
}

func cmd(Uname string, ctx string) string {
	s := strings.Replace(ctx, "!cmd", "", -1)
	a := strings.Split(s, ":")
	name, value := a[0], a[1]
	//fixes _
	name = strings.Replace(name, " ", "", 1)
	if Uname == "NAME" {
		datas(name, value)
		return "Готово PogChamp"
	} else {
		return "Это может делать только NAME NotLikeThis"
	}
}

func cmdtest(ctx string) bool {
	i := dataq(ctx).Value
	if i != "" {
		return true
	} else {
		return false
	}
}

func custcmd(ctx string) string {
	querry := dataq(ctx)
	i := querry.Value
	cc := strconv.Itoa(querry.Counter)
	return strings.Replace(i, "!C", cc, -1)
}
