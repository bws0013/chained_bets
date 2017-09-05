package main

import (
  "net"
  "fmt"
  "encoding/gob"
  "bufio"
  "math/rand"
  "time"
)

type bet_packet struct {
  Key int32
  Bet float32
  Res int
}

func main() {
  // Generate random number to send as id
  rand.Seed(time.Now().UTC().UnixNano())

  k := rand.Int31() + 1 // x>0
  b := float32(10)
  r := rand.Intn(2)

  packet := bet_packet{k, b, r}

  dial_server_packet(packet)
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
    if err != nil { panic(err) }
    if len(message) != 0 { fmt.Printf("%s\n", message) }
}
