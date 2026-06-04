package main

import "fmt"

type Player struct {
	firstName string
	lastName  string
	goals     int
	assists   int
}

func insertionSort(players []Player) {
	n := len(players)
	for i := 1; i < n; i++ {
		key := players[i]
		j := i - 1
		for j >= 0 && (players[j].goals < key.goals ||
			(players[j].goals == key.goals && players[j].assists < key.assists)) {
			players[j+1] = players[j]
			j--
		}
		players[j+1] = key
	}
}

func main() {
	fmt.Print("Masukkan Data Input :\n")

	var n int
	fmt.Scan(&n)

	players := make([]Player, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&players[i].firstName, &players[i].lastName,
			&players[i].goals, &players[i].assists)
	}

	insertionSort(players)

	fmt.Println("\nHasil Sorting :")
	for _, p := range players {
		fmt.Printf("%s %s %d %d\n", p.firstName, p.lastName, p.goals, p.assists)
	}
}
