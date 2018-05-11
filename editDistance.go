package main

import (
	"strings"
	"math"
	"fmt"
)

func main(){
	dp := [10][10]float64{}
	exec := strings.Split(" _e_x_e_c_u_t_i_o_n", "_")
	inte := strings.Split(" _i_n_t_e_n_t_i_o_n", "_")

	for y := 0; y<len(dp); y++{
		for x := 0; x<len(dp[0]); x++{
			if y == 0 {
				dp[y][x] = float64(x)
			}else if x == 0{
				dp[y][x] = float64(y)
			}else{
				choice1 := math.Min(dp[y-1][x], dp[y][x-1])
				choice := math.Min(choice1, dp[y-1][x-1])
				dp[y][x] = choice
				if exec[y] != inte[x]{
					dp[y][x]++
				}
			}
		}
	}

	fmt.Println(dp[9][9])

}