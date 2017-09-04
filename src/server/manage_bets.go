package main

import (
  "fmt"
)

func pr() {
  for item := range bet_map.IterBuffered() {
    val := item.Val
    fmt.Println(val)
  }
}

// collect the bets and find the amount of bets for each state
// start with 2 states 0/1
func bets_per_state() {

  // map[Res]Bet, map[int]float32
  state_map := make(map[int]float32)

  bets := organize_bets()

  total := float32(0)

  for _, bet := range bets {
    if val, ok := state_map[bet.Res]; ok {
      current_amount := val
      current_amount += bet.Bet
      state_map[bet.Res] = current_amount
    } else {
      state_map[bet.Res] = bet.Bet
    }
    total += bet.Bet
  }

  for k, v := range state_map {
    fmt.Println(k, "->", v)
  }
  fmt.Println("Total:", total)

}
