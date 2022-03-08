package main

import (
	"fmt"
	"strings"
)

func (g *Game) GetRoom(name string) *Room {
	return g.Rooms[name]
}

var g Game

func main() {
	/*
		в этой функции можно ничего не писать
		но тогда у вас не будет работать через go run main.go
		очень круто будет сделать построчный ввод команд тут, хотя это и не требуется по заданию
	*/
	initGame()
	fmt.Println(g)
	fmt.Println(g.Players[0].View())
	fmt.Println(g.Players[0].Go("коридор"))
	fmt.Println(g.Players[0].View())
	fmt.Println(g.Players[0].Bag)
	g.Players[0].Bag = &Backpack{
		Lock:  true,
		Items: map[string]bool{},
	}
	if g.Players[0].Bag.Lock {
		fmt.Println("На вас надет рюкзак")
	} else {
		fmt.Println("На вас не надет рюкзак")
	}
	fmt.Println(g.Players[0].Go("кухня"))
	fmt.Println(g.Players[0].InRoom.WhereToGo())
	fmt.Println(g.Players[0].Go("улица"))
}

func initGame() {
	/*
		эта функция инициализирует игровой мир - все команты
		если что-то было - оно корректно перезатирается
	*/
	g = Game{
		Rooms:   make(map[string]*Room),
		Players: make([]Player, 0, 1),
	}

	g.Rooms = map[string]*Room{
		"кухня": {
			Game: &g,
			Name: "кухня",
			Items: map[string]Item{
				"чай": {
					Name:  "чай",
					Exist: false,
				},
			},
		},
		"коридор": {
			Game:  &g,
			Name:  "коридор",
			Items: map[string]Item{},
		},
		"комната": {
			Game: &g,
			Name: "комната",
			Items: map[string]Item{
				"ключи": {
					Name:  "ключи",
					Exist: false,
				},
				"рюкзак": {
					Name:  "рюкзак",
					Exist: false,
				},
				"конспекты": {
					Name:  "конспекты",
					Exist: false,
				},
			},
		},
		"улица": {
			Game:  &g,
			Name:  "улица",
			Items: map[string]Item{},
		},
	}

	g.Players = []Player{
		{
			InRoom: g.GetRoom("кухня"),
		},
	}

	g.Ways = []Way{
		{RoomFrom: g.GetRoom("кухня"), RoomTo: g.GetRoom("коридор")},
		{RoomFrom: g.GetRoom("коридор"), RoomTo: g.GetRoom("кухня")},
		{RoomFrom: g.GetRoom("коридор"), RoomTo: g.GetRoom("комната")},
		{RoomFrom: g.GetRoom("комната"), RoomTo: g.GetRoom("коридор")},
		{RoomFrom: g.GetRoom("коридор"), RoomTo: g.GetRoom("улица")},
		{RoomFrom: g.GetRoom("улица"), RoomTo: g.GetRoom("коридор")},
	}
}

func handleCommand(command string) string {
	/*
		данная функция принимает команду от "пользователя"
		и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
	*/
	p := &g.Players[0]
	action := strings.Split(command, " ")
	switch action[0] {
	case "осмотреться":
		return p.View()
	case "идти":
		{
			p.Go(g.GetRoom(action[1]).Name)
			return "Идти"
		}
	case "надеть":
		return "Надеть"
	case "взять":
		return "Взять"
	case "применить":
		return "Применить"
	default:
		return "неизвестная команда"
	}
}
