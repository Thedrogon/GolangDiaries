package exercise

type Item struct {
	name string
	Type string
}

type Player struct {
	name string
	inventory []Item
}


func(p* Player) PickupItem(name string , item string) bool{
	
	p.inventory = append(p.inventory, Item{name , item})
	return true
}

func(p* Player)Updateinfo(name string) bool{
	p.name = name
	return true
}