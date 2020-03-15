package base

//座位
type Seat struct {
	Id         int
	PlayerId   int64
	NextSeatId int
	PrevSeatId int
}

func (seat *Seat) Seat() {
	seat.Id = 0
	seat.PlayerId = 0
	seat.NextSeatId = 0
	seat.PrevSeatId = 0
}

//座位上是否有人
func (seat *Seat) HasPlayer() bool {
	return seat.Id != 0
}
