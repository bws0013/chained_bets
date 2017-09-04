package main

import (
  "fmt"
)

func get_result() int {
  // Add result condition stuff later
  return 0
}

func distribute_winnings(winning_state int, win_multiple float32) {
  var winnings_list []string
  if win_multiple == 0 {
    winnings_list = return_all_bets()
  } else {
    winnings_list = calc_winnings_amount(winning_state, win_multiple)
  }

  for _, w := range winnings_list {
    fmt.Println(w)
  }

}


func calc_winnings_amount(winning_state int, win_multiple float32) []string {
  var winnings_list []string
  bets := organize_bets()
  _, total := bets_per_state()

  for _, bet := range bets {
    if bet.Res == winning_state {
      total_winning_est := bet.Bet * win_multiple

      if total_winning_est > total {
        winnings_string := fmt.Sprintf("%f -> %d", total, bet.Key)
        winnings_list = append(winnings_list, winnings_string)
        total = 0
      } else {
        winnings_string := fmt.Sprintf("%f -> %d", total_winning_est, bet.Key)
        winnings_list = append(winnings_list, winnings_string)
        total -= total_winning_est
      }
    }
  }
  return winnings_list
}

func return_all_bets() []string {
  var winnings_list []string
  bets := organize_bets()
  for _, bet := range bets {
    winnings_string := fmt.Sprintf("%f -> %s", bet.Bet, bet.Key)
    winnings_list = append(winnings_list, winnings_string)
  }

  return winnings_list
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
