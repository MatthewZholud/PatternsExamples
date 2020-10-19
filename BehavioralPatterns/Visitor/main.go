package main

import "fmt"

type shape interface {
	getType() string
	accept(visitor)
}

type square struct {
	side int
}

func (s *square) accept(v visitor)  {
	v.visitForSquare(s)
}

func (s *square) getType() string {
	return "Square"
}

type circle struct {
	radius int
}

func (c *circle) accept(v visitor)  {
	v.visitForCircle(c)
}

func (c *circle) getType() string {
	return "Circle"
}

type rectangle struct {
	l int
	b int
}

func (r *rectangle) accept(v visitor)  {
	v.visitForRectangle(r)
}

func (r *rectangle) getType() string {
	return "Rectangle"
}

type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForRectangle(*rectangle)
}

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square)  {
	fmt.Println("Calculating area for square")
}

func (a *areaCalculator) visitForCircle(s *circle) {
	fmt.Println("Calculating area for circle")
}

func (a *areaCalculator) visitForRectangle(s *rectangle) {
	fmt.Println("Calculating area for rectangle")
}

type middleCoordinates struct {
	x int
	y int
}

func (a *middleCoordinates) visitForSquare(s *square) {
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *middleCoordinates) visitForCircle(c *circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}

func (a *middleCoordinates) visitForRectangle(t *rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}

func main()  {
	sq := &square{side: 2}
	cr := &circle{radius: 3}
	rct := &rectangle{l: 2, b: 3}
	areaCalc := &areaCalculator{}

	sq.accept(areaCalc)
	cr.accept(areaCalc)
	rct.accept(areaCalc)

	fmt.Println()
	middleCoord := &middleCoordinates{}
	sq.accept(middleCoord)
	cr.accept(middleCoord)
	rct.accept(middleCoord)
}


