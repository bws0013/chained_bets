package main

import (
  "fmt"
  "net"
  "time"
  "strconv"
  "encoding/gob"
  "github.com/orcaman/concurrent-map"
)

type bet_packet struct {
  Key int32
  Bet float32
  Res int
}

var (
  bet_map = cmap.New() // Global map we store bets in
  timer_time = time.Duration(5) // Amount of seconds betting is open for
)

func main() {
  run_any_test(all_win_test)
  run_any_test(all_lost_test)
}

// Used exclusively for testing
func restart_map() {
  bet_map = cmap.New()
}

func instance() {
  collect_bets()

  print_items()
  print_state_map()
  winning_state := get_result()
  mult := calc_winnings_multiple(winning_state)

  distribute_winnings(winning_state, mult)
}

// Return a list of the bets that were made
func organize_bets() []bet_packet {
  bets := []bet_packet{}
  for item := range bet_map.IterBuffered() {
    val := item.Val
    packet := val.(*bet_packet)

    // We may want to if-statement this to remove any key of 0
    if packet.Key != 0 {
      bets = append(bets, *packet)
    }
  }
  return bets
}

// Open the connecion and wait on people to send in bets
// Close after the timer runs out
func collect_bets() {
  fmt.Println("start");
  ln, err := net.Listen("tcp", ":8081")
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
  for cont == true {
    conn, err := ln.Accept()
    if err != nil { fmt.Println("This connection needs a tissue, skipping!") }
    go func() { listen_packet(conn) }()
  }
}

// Listen for packets
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

// Send a packet to ensure we close at the right time
func send_close_packet() {
  packet := bet_packet{0, 0.0, 0}
  conn, _ := net.Dial("tcp", "127.0.0.1:8081")
  encoder := gob.NewEncoder(conn)
  _ = encoder.Encode(&packet)
  conn.Close()
}
