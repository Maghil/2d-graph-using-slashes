package main

import "fmt"

func main() {
	data := []int{3, 1, 2, 3, 6, 2, 3, 6, 2, 3, 6, 3, 2, 3, 6, 2, 3, 4, 3, 2, 5, 4, 2, 1, 2, 1, 2, 3, 1, 2, 6, 2, 3, 6, 2, 3, 6, 3, 2, 3, 1, 5, 3, 2, 1, 2, 4, 2, 1, 8, 1, 2}
	populate_graph(data)
}

func populate_graph(data []int) {
	y_is_zero := false
	graph := make([][]string, 0)
	max_x := sum(data) //horizontal axis
	max_y := -1        //vertical axis
	var x, y int
	for i, value := range data {
		for j := 0; j < value; j++ { //to iterate over each value of data to know the number of slashes required
			if !y_is_zero { // flag to check if the graph is going to cross over to -y region
				if i%2 == 0 {
					if y > max_y {
						max_y += 1
						temp := make([]string, max_x) //creating a empty slice so it can be filled with slashes as requried
						graph = append(graph, temp)
					}
					graph[y][x] = "/"
					y += 1
				} else {
					if x == 0 {
						print("the graph is getting pushed to -x axis, we currently don't a support for that \n")
						y_is_zero = true
						print("PRINTING GRAPH BEFORE REACHING -X AXIS")
						break
					}
					y -= 1
					graph[y][x] = "\\"
				}
				x += 1
			}
		}
	}
	print_graph(graph, max_x, max_y)
}

func print_graph(graph [][]string, max_x int, max_y int) {
	for i := max_y; i >= 0; i-- { // printing in reverse order since graph was populated from bot to top
		for j := 0; j < max_x; j++ {
			if graph[i][j] == "" {
				print(" ")
			}
			fmt.Printf("%v", graph[i][j])
		}
		fmt.Print("\n")
	}
}

func sum(data []int) int {
	var total int
	for i := range data {
		total += data[i]
	}
	return total
}
