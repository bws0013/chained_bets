package main

import (
  "fmt"
  "net"
  "time"
  "strconv"
  "encoding/gob"
  "github.com/orcaman/concurrent-map"
)
//"sync"

// TODO incorporate this https://medium.com/@lhartikk/a-blockchain-in-200-lines-of-code-963cc1cc0e54

type bet_packet struct {
  Key int32
  Bet float32
  Res int
}

var (
  bet_map = cmap.New()
  timer_time = time.Duration(10)
)

func main() {

  collect_bets()

  //pr()

  // bps := bets_per_state()
  print_items()
  print_state_map()

  winning_state := get_result()
  mult := calc_winnings_multiple(winning_state)

  distribute_winnings(winning_state, mult)

}

func organize_bets() []bet_packet {
  bets := []bet_packet{}
  for item := range bet_map.IterBuffered() {
    val := item.Val
    packet := val.(*bet_packet)

    // We may want to if-statement this to remove any key of 0
    bets = append(bets, *packet)
  }
  return bets
}

func collect_bets() {
  fmt.Println("start");
  ln, err := net.Listen("tcp", ":8081")
  // check_err(err, "Listening!")
  if err != nil {
    fmt.Println("Error at listen")
    return
  }

  cont := true
  timer2 := time.NewTimer(time.Second * timer_time)
  go func() {
      <-timer2.C
      cont = false
      send_close_packet()
  }()
  for {
    if cont == false { break }
    conn, err := ln.Accept()
    if err != nil { fmt.Println("This connection needs a tissue, skipping!") }
    go func() { listen_packet(conn) }()
  }
}

func listen_packet(conn net.Conn) {
  dec := gob.NewDecoder(conn)
  p := &bet_packet{}
  err := dec.Decode(p)
  if err != nil { fmt.Println("Tell me about it") }
  key := strconv.Itoa(int(p.Key))
  bet_map.Set(key, p)
  conn.Write([]byte("liftoff"))
  conn.Close()
  if dec != nil {
    return
  }
}

func send_close_packet() {
  packet := bet_packet{0, 0.0, 0}
  conn, _ := net.Dial("tcp", "127.0.0.1:8081")
  encoder := gob.NewEncoder(conn)
  _ = encoder.Encode(&packet)
  conn.Close()
}
