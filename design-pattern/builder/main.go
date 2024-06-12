package main

import "fmt"

type HouseBuilder interface {
	buildWalls(walls int) int
	buildDoors(door string) string
	buildFloors(floors int) int
	build() House
}

type House struct {
	walls  int
	door   string
	floors int
}

type HouseDirector struct {
	builder HouseBuilder
}

type ConceretBuilder struct {
	house House
}

func (cb *ConceretBuilder) build() House {
	return cb.house
}

func (cb *ConceretBuilder) buildWalls(walls int) int {
	cb.house.walls = walls
	return cb.house.walls
}

func (cb *ConceretBuilder) buildDoors(door string) string {
	cb.house.door = door
	return cb.house.door
}

func (cb *ConceretBuilder) buildFloors(floors int) int {
	cb.house.floors = floors
	return cb.house.floors
}

func (hd *HouseDirector) Construct4FloorDoorStone() House {
	hd.builder.buildDoors("Stone")
	hd.builder.buildFloors(4)
	hd.builder.buildWalls(4)
	return hd.builder.build()
}

func main() {
	builder := &ConceretBuilder{}
	director := &HouseDirector{builder: builder}

	house1 := director.Construct4FloorDoorStone()
	fmt.Println(house1)

}
