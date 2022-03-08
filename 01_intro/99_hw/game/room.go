package main

func (r *Room) WhereToGo() []Way {
	links := make([]Way, 0, 3)
	for _, way := range r.Game.Ways {
		if way.RoomFrom.Name == r.Name {
			links = append(links, way)
		}
	}
	return links
}

func (r *Room) ConnectedTo(r2 *Room) bool {
	for _, way := range r.WhereToGo() {
		if way.RoomFrom.Name == r.Name && way.RoomTo.Name == r2.Name {
			return true
		}
	}
	return false
}
