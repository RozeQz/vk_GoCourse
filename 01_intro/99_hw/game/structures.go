package main

type Room struct {
	Game  *Game
	Name  string
	Items map[string]Item
	Ways  []Room
}

type Item struct {
	Name  string
	Exist bool
}

type Object struct {
	Name string
	Lock bool
}

type Backpack struct {
	Lock  bool
	Items map[string]bool
}

type Way struct {
	RoomFrom *Room
	RoomTo   *Room
}

type Player struct {
	InRoom *Room
	Bag    *Backpack
}

type Game struct {
	Rooms   map[string]*Room
	Players []Player
	Ways    []Way
}
