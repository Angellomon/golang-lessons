package main

import (
	"fmt"
)

type SpecialPosition struct {
	Position
	z float64
}

func (sp *SpecialPosition) MoveSpecial(x, y float64) {
	sp.x += x * x
	sp.y += y * y
}

type Position struct {
	x float64
	y float64
}

func (p *Position) Move(x, y float64) {
	p.x += x
	p.y += y
}

func (p *Position) Teleport(x, y float64) {
	p.x = x
	p.y = y
}

type Player struct {
	*Position
}

func CreateDefaultPlayer() *Player {
	return &Player{
		Position: &Position{},
	}
}

/*func (p* Player) Teleport(x, y float64) {
	p.posX = x
	p.posY = y
}*/

type Enemy struct {
	*SpecialPosition
}

func CreateDefaultEnemy() *Enemy {
	return &Enemy{
		&SpecialPosition{},
	}
}

/*func (e* Enemy) Move(x, y float64) {
	e.posX = x
	e.posY = y
}*/

func main() {
	player1 := CreateDefaultPlayer()

	fmt.Println(player1.Position)

	player1.Move(1.1, 2.1)

	fmt.Println(player1.Position)

	enemy1 := CreateDefaultEnemy()

	enemy1.Move(2, 3)
	fmt.Println("normal move: ", enemy1.Position)

	enemy1.MoveSpecial(2.3, 4.4)
	fmt.Println("enemy special move: ", enemy1.Position)
	fmt.Println("enemy z move: ", enemy1.SpecialPosition)
}
