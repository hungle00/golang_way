package main

import "fmt"

type ToyRobot struct {
	x         int
	y         int
	direction Direction
	bounds   Bounds
}

type Direction string

const (
	North Direction = "NORTH"
	East  Direction = "EAST"
	South Direction = "SOUTH"
	West  Direction = "WEST"
)

type Bounds	struct {
	MinX, MinY int
	MaxX, MaxY int
}

var DefaultBounds = Bounds{MinX: 0, MaxX: 10, MinY: 0, MaxY: 10}

func NewToyRobot() *ToyRobot {
	return &ToyRobot{
		x:         0,
		y:         0,
		direction: North,
		bounds:    DefaultBounds,
	}
}

// Getters
func (r *ToyRobot) GetX() int { return r.x }
func (r *ToyRobot) GetY() int { return r.y }
func (r *ToyRobot) Direction() Direction { return r.direction }

func (r *ToyRobot) Place(x, y int, direction Direction) {
	r.x = x
	r.y = y
	r.direction = direction
}

func (r *ToyRobot) Report() string {
	return fmt.Sprintf("X: %d, Y: %d, Direction: %s", r.x, r.y, r.direction)
}

func (r *ToyRobot) Move() {
	switch r.direction {
	case North:
		r.y++
	case East:
		r.x++
	case South:
		r.y--
	case West:
		r.x--
	}
}

// Right turns the robot 90 degrees clockwise.
func (r *ToyRobot) Right() {
	switch r.direction {
	case North:
		r.direction = East
	case East:
		r.direction = South
	case South:
		r.direction = West
	case West:
		r.direction = North
	}
}

// Left turns the robot 90 degrees counter-clockwise.
func (r *ToyRobot) Left() {
	switch r.direction {
	case North:
		r.direction = West
	case West:
		r.direction = South
	case South:
		r.direction = East
	case East:
		r.direction = North
	}
}


func main() {
	robot := NewToyRobot()
	robot.Place(0, 0, North)
	fmt.Println(robot.Report()) // Output: X: 0, Y: 0, Direction: NORTH

	robot.Move()
	fmt.Println(robot.Report()) // Output: X: 0, Y: 1, Direction: NORTH

	robot.Right()
	robot.Move()
	fmt.Println(robot.Report()) // Output: X: 1, Y: 1, Direction: EAST
}
