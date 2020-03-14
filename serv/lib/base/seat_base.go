package lib

type SeatBase struct {
	id     int
	prevId int
	nextId int
}

func (base *SeatBase) GetId() int {
	return base.id
}

func (base *SeatBase) SetId(id int) {
	base.id = id
}

func (base *SeatBase) GetPrevId() int {
	return base.prevId
}

func (base *SeatBase) SetPrevId(id int) {
	base.prevId = id
}

func (base *SeatBase) GetNextId() int {
	return base.nextId
}

func (base *SeatBase) SetNextId(id int) {
	base.nextId = id
}
