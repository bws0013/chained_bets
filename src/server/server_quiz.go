package main

import (
  "net"
  "fmt"
  "encoding/gob"
  "bufio"
  "sync"
  "time"
)

// Use this as a template (see client.go)
// func main() {
//   // Generate random number to send as id
//   rand.Seed(time.Now().UTC().UnixNano())
//
//   k := rand.Int31() + 1 // x>0
//   b := float32(10)
//   r := rand.Intn(2)
//
//   packet := bet_packet{k, b, r}
//
//   dial_server_packet(packet)
// }

var wg sync.WaitGroup

// Pass in the method we are testing for
func run_any_test(test func()) {

  wg.Add(1)
  restart_map()
  go func() {
    defer wg.Done()
    instance()
    time.Sleep(time.Second * 1)
  }()
  time.Sleep(time.Second * 1)

  test()

  wg.Wait()
  fmt.Println("====================")
}

// test where all parties win
func all_win_test() {
  k1 := int32(1) // x>0
  b1 := float32(10)
  r1 := 0

  k2 := int32(2) // x>0
  b2 := float32(10)
  r2 := 0

  packet1 := bet_packet{k1, b1, r1}
  packet2 := bet_packet{k2, b2, r2}

  dial_server_packet(packet1)
  dial_server_packet(packet2)
}

// test where all parties lose
func all_lost_test() {
  k1 := int32(1) // x>0
  b1 := float32(10)
  r1 := 1

  k2 := int32(2) // x>0
  b2 := float32(10)
  r2 := 1

  packet1 := bet_packet{k1, b1, r1}
  packet2 := bet_packet{k2, b2, r2}

  dial_server_packet(packet1)
  dial_server_packet(packet2)
}

// test where someone wins and someone does not
func expected_test() {

}

// test where no packets are sent
func no_sent_test() {

}

// test where a series of packets are sent serially
func lost_of_packets_serial_test() {

}

// test where a series of packets are sent in parallel
func lost_of_packets_parallel_test() {

}

// test where incorrect packet type is sent
func incorrect_packet_type_test() {

}

// test where we keep sending packets even after the timer is done
func send_packet_per_second_test() {

}

// test where we test a weird distribution
func weird_distribution_test() {

}

// Send packet
func dial_server_packet(packet bet_packet) {
  conn, err := net.Dial("tcp", "127.0.0.1:8081")

  if err != nil {
    fmt.Println("Unable to send!")
    return
  }
  encoder := gob.NewEncoder(conn)
  err = encoder.Encode(&packet)

  message, _ := bufio.NewReader(conn).ReadString('\n')
  fmt.Println("Message from server: " + message)
  check_err(err, "everything is fine")
  conn.Close()

}

// Print out if there is an error
func check_err(err error, message string) {
    if err != nil {
      panic(err)
    }
    if len(message) != 0 {
      fmt.Printf("%s\n", message)
    }
}
