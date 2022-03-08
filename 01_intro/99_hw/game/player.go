package main

func (p *Player) Apply(o *Item) bool {
	if o.Exist {
		return false
	} else {
		return true
	}
}

func (p *Player) View() string {
	var obj string
	for _, v := range p.InRoom.Items {
		obj += v.Name + " "
	}
	return "ты находишься: " + p.InRoom.Name + ". " + "предметы: " + obj
}

func (p *Player) Go(nextRoom string) string {
	if p.InRoom.ConnectedTo(g.GetRoom(nextRoom)) {
		p.InRoom = g.GetRoom(nextRoom)
		return "Вы находитесь в " + nextRoom
	} else {
		return "нет пути в " + nextRoom
	}
}
