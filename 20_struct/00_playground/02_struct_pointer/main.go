package main

import (
	"fmt"
)

type player struct {
	playerName string
	level      int
	xp         int
	health     int
}

// Notice that we could use type player as function reciever.
// But the purpose of this example is to use pointers.
func levelUp(p *player) {
	p.level++
	restoreHealth(p)
	gainXP(p)
}

func restoreHealth(p *player) {
	p.health = 100
}

func gainXP(p *player) {
	p.xp += 1000
}

func main() {

	var p1 player
	fmt.Println(p1) // { 0 0 0}
	levelUp(&p1)
	fmt.Println(p1) // { 1 1000 100}

}
