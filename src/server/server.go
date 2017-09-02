package main

import (
  "fmt"
  "net"
  "strconv"
  "encoding/gob"
  "github.com/orcaman/concurrent-map"
  "time"
)
//"sync"

// TODO incorporate this https://medium.com/@lhartikk/a-blockchain-in-200-lines-of-code-963cc1cc0e54

type my_packet struct {
  Key int32
  Bet float32
  Res int
}

var (
  bet_map = cmap.New()
)

func main() {

  fmt.Println("start");
  ln, err := net.Listen("tcp", ":8081")
  // check_err(err, "Listening!")
  if err != nil {
    fmt.Println("Error at listen")
    return
  }

  total_connections := 0

  cont := true

  timer2 := time.NewTimer(time.Second * 60)
  go func() {
      <-timer2.C
      cont = false
      send_close_packet()
  }()

  for {
    if cont == false {
      break
    }

    conn, err := ln.Accept()
    if err != nil {
        fmt.Println("This connection needs a tissue, skipping!")
    }
    go func() {
      listen_packet(conn)
    }()

    total_connections++
    fmt.Println("Connection: ", total_connections)
  }

  // https://github.com/orcaman/concurrent-map
  // https://github.com/orcaman/concurrent-map/blob/master/concurrent_map_test.go
  for item := range bet_map.IterBuffered() {
      val := item.Val
      fmt.Println(val)
  }
}

func hw() {
  fmt.Println("hello world!")
}

func listen_packet(conn net.Conn) {

  dec := gob.NewDecoder(conn)
  p := &my_packet{}
  err := dec.Decode(p)

  if err != nil { fmt.Println("Tell me about it") }

  key := strconv.Itoa(int(p.Key))

  bet_map.Set(key, p)

  conn.Write([]byte("liftoff"))

  conn.Close()

  if dec != nil {
    fmt.Printf("Client disconnected.\n")
    return
  }
}

func send_close_packet() {
  packet := my_packet{0, 0.0, 0}
  conn, _ := net.Dial("tcp", "127.0.0.1:8081")
  encoder := gob.NewEncoder(conn)
  _ = encoder.Encode(&packet)
  conn.Close()
}
