## Todo

> User score BROKEN (+)
> > No collection (+)
> > > No updates (+)
> 
> Video almost done, !nextvideo overshot (+) - done
> 
> cmd din't work - done
> > nil return (+) - done
> > > Auto Count (+) - done
>
> Websockets (+)
> > Working (+)
> > > Problem with reconnect and missing error

## Build
go build main.go funks.go database.go websocketserver.go

***************
https://habr.com/ru/post/306840/
https://github.com/gobwas/ws


## OpCode
https://godoc.org/github.com/gobwas/ws#OpCode

## IoWriter
https://godoc.org/io#Writer

https://github.com/gobwas/ws-examples/blob/master/src/chat/main.go
func (d deadliner) Write(p []byte) (int, error) {
	if err := d.Conn.SetWriteDeadline(time.Now().Add(d.t)); err != nil {
		return 0, err
	}
	return d.Conn.Write(p)
}

https://github.com/gobwas/ws-examples/blob/master/src/chat/chat.go
func (u *User) write(x interface{}) error {
	w := wsutil.NewWriter(u.conn, ws.StateServerSide, ws.OpText)
	encoder := json.NewEncoder(w)

	u.io.Lock()
	defer u.io.Unlock()

	if err := encoder.Encode(x); err != nil {
		return err
	}

	return w.Flush()
}