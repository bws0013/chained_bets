package main

import (
  "fmt"
  "net"
  "strconv"
  "encoding/gob"
  "github.com/orcaman/concurrent-map"
  "sync"
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

  var wg sync.WaitGroup
  cont := true

  timer2 := time.NewTimer(time.Second * 15)
  go func() {
      <-timer2.C
      cont = false
  }()

  for {
    conn, err := ln.Accept() // this blocks until connection or error
    if err != nil {
        fmt.Println("This connection needs a tissue, skipping!")
        continue
    }

    wg.Add(1)
    go func() {
      defer wg.Done()
      listen_packet(conn)
    }()

    if cont == false {
      break
    }

    total_connections++
    fmt.Println("Connection: ", total_connections)

    // if total_connections >= 2 {
    //   wg.Wait()
    //   break
    // }

  }

  // https://github.com/orcaman/concurrent-map
  // https://github.com/orcaman/concurrent-map/blob/master/concurrent_map_test.go
  for item := range bet_map.IterBuffered() {
      val := item.Val
      fmt.Println(val)
  }


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
