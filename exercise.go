package exercise


type Item struct {
	name string
	Type string
}

type Playerinfo struct {
	name string
	inventory []Item
}


func(p* Playerinfo)Addplayer(name string,){

}


func(p* Playerinfo) PickupItem(name string , item string) bool{
	
	p.inventory = append(p.inventory, Item{name , item})
	return true
}

func(p* Playerinfo)Updateinfo(name string) bool{
	p.name = name
	return true
}

