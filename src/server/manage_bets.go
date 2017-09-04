package main

import (
  "fmt"
)

func get_result() int {
  // Add result condition stuff later
  return 0
}

func calc_winnings_amount(winning_state int, win_multiple float32) {

  if win_multiple == 0 {
    return_all_bets()
    return
  }

  // bets := organize_bets()




}

/*
type bet_packet struct {
  Key int32
  Bet float32
  Res int
}
*/

func return_all_bets() {
  bets := organize_bets()
  for _, bet := range bets {
    fmt.Println(bet.Bet, "->", bet.Key)
  }
}

func calc_winnings_multiple(winning_state int) float32 {
  state_map, total := bets_per_state()
  total_winning_bet_amount := state_map[winning_state]
  if total_winning_bet_amount == 0 { return 0.0 }
  win_multiple := total / total_winning_bet_amount
  return win_multiple
}

// collect the bets and find the amount of bets for each state
// start with 2 states 0/1
func bets_per_state() (map[int]float32, float32) {
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
  return state_map, total
}

// *****Print functions*****
func print_items() {
  for item := range bet_map.IterBuffered() {
    val := item.Val
    fmt.Println(val)
  }
}
func print_state_map() {
  state_map, total := bets_per_state()
  for k, v := range state_map {
    fmt.Println(k, "->", v)
  }
  fmt.Println("Total:", total)
}
