package base

type Seat struct {
	Id     int
	PrevId int
	NextId int
}

//座位上是否有人
func (seat *Seat) HasPlayer() bool {
	return seat.Id != 0
}
