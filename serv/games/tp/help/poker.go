package help

import (
	sf "github.com/hanjingo/algorithm/shuffle"
)

/**
一张牌
*/
type Poker struct {
	value byte //前4位花色 后四位点数
}

func NewPoker(v byte) *Poker {
	return &Poker{value: v}
}

//获得点数
func (pk *Poker) Point() byte {
	return pk.value & mask_point
}

//获得花色
func (pk *Poker) Color() byte {
	return pk.value & mask_color
}

/**
一套牌
*/
type PokerSuit struct {
	Cards []*Poker
}

func NewPokerSuit() *PokerSuit {
	return &PokerSuit{}
}

//洗下牌
func (ps *PokerSuit) Shuffle() {
	sf.Shuffle(ps.Cards)
}

func (ps *PokerSuit) Init() {
	//方块
	ps.Cards = append(ps.Cards, NewPoker(diamond_2))
	ps.Cards = append(ps.Cards, NewPoker(diamond_3))
	ps.Cards = append(ps.Cards, NewPoker(diamond_4))
	ps.Cards = append(ps.Cards, NewPoker(diamond_5))
	ps.Cards = append(ps.Cards, NewPoker(diamond_6))
	ps.Cards = append(ps.Cards, NewPoker(diamond_7))
	ps.Cards = append(ps.Cards, NewPoker(diamond_8))
	ps.Cards = append(ps.Cards, NewPoker(diamond_9))
	ps.Cards = append(ps.Cards, NewPoker(diamond_10))
	ps.Cards = append(ps.Cards, NewPoker(diamond_J))
	ps.Cards = append(ps.Cards, NewPoker(diamond_Q))
	ps.Cards = append(ps.Cards, NewPoker(diamond_K))
	ps.Cards = append(ps.Cards, NewPoker(diamond_A))

	//梅花
	ps.Cards = append(ps.Cards, NewPoker(club_2))
	ps.Cards = append(ps.Cards, NewPoker(club_3))
	ps.Cards = append(ps.Cards, NewPoker(club_4))
	ps.Cards = append(ps.Cards, NewPoker(club_5))
	ps.Cards = append(ps.Cards, NewPoker(club_6))
	ps.Cards = append(ps.Cards, NewPoker(club_7))
	ps.Cards = append(ps.Cards, NewPoker(club_8))
	ps.Cards = append(ps.Cards, NewPoker(club_9))
	ps.Cards = append(ps.Cards, NewPoker(club_10))
	ps.Cards = append(ps.Cards, NewPoker(club_J))
	ps.Cards = append(ps.Cards, NewPoker(club_Q))
	ps.Cards = append(ps.Cards, NewPoker(club_K))
	ps.Cards = append(ps.Cards, NewPoker(club_A))

	//红桃
	ps.Cards = append(ps.Cards, NewPoker(heart_2))
	ps.Cards = append(ps.Cards, NewPoker(heart_3))
	ps.Cards = append(ps.Cards, NewPoker(heart_4))
	ps.Cards = append(ps.Cards, NewPoker(heart_5))
	ps.Cards = append(ps.Cards, NewPoker(heart_6))
	ps.Cards = append(ps.Cards, NewPoker(heart_7))
	ps.Cards = append(ps.Cards, NewPoker(heart_8))
	ps.Cards = append(ps.Cards, NewPoker(heart_9))
	ps.Cards = append(ps.Cards, NewPoker(heart_10))
	ps.Cards = append(ps.Cards, NewPoker(heart_J))
	ps.Cards = append(ps.Cards, NewPoker(heart_Q))
	ps.Cards = append(ps.Cards, NewPoker(heart_K))
	ps.Cards = append(ps.Cards, NewPoker(heart_A))

	//黑桃
	ps.Cards = append(ps.Cards, NewPoker(spade_2))
	ps.Cards = append(ps.Cards, NewPoker(spade_3))
	ps.Cards = append(ps.Cards, NewPoker(spade_4))
	ps.Cards = append(ps.Cards, NewPoker(spade_5))
	ps.Cards = append(ps.Cards, NewPoker(spade_6))
	ps.Cards = append(ps.Cards, NewPoker(spade_7))
	ps.Cards = append(ps.Cards, NewPoker(spade_8))
	ps.Cards = append(ps.Cards, NewPoker(spade_9))
	ps.Cards = append(ps.Cards, NewPoker(spade_10))
	ps.Cards = append(ps.Cards, NewPoker(spade_J))
	ps.Cards = append(ps.Cards, NewPoker(spade_Q))
	ps.Cards = append(ps.Cards, NewPoker(spade_K))
	ps.Cards = append(ps.Cards, NewPoker(spade_A))
}

//牌的id
var (
	diamond_2 uint8 = color_diamond | point_2 //方块2 (0000 0001)
	club_2    uint8 = color_club | point_2    //梅花2 (0001 0001)
	heart_2   uint8 = color_heart | point_2   //红桃2 (0010 0001)
	spade_2   uint8 = color_spade | point_2   //黑桃2 (0011 0001)

	diamond_3 uint8 = color_diamond | point_3 //方块3 (0000 0010)
	club_3    uint8 = color_club | point_3    //梅花3 (0001 0010)
	heart_3   uint8 = color_heart | point_3   //红桃3 (0010 0010)
	spade_3   uint8 = color_spade | point_3   //黑桃3 (0011 0010)

	diamond_4 uint8 = color_diamond | point_4 //方块4 (0000 0011)
	club_4    uint8 = color_club | point_4    //梅花4 (0001 0011)
	heart_4   uint8 = color_heart | point_4   //红桃4 (0010 0011)
	spade_4   uint8 = color_spade | point_4   //黑桃4 (0011 0011)

	diamond_5 uint8 = color_diamond | point_5 //方块5 (0000 0100)
	club_5    uint8 = color_club | point_5    //梅花5 (0001 0100)
	heart_5   uint8 = color_heart | point_5   //红桃5 (0010 0100)
	spade_5   uint8 = color_spade | point_5   //黑桃5 (0011 0100)

	diamond_6 uint8 = color_diamond | point_6 //方块6 (0000 0101)
	club_6    uint8 = color_club | point_6    //梅花6 (0001 0101)
	heart_6   uint8 = color_heart | point_6   //红桃6 (0010 0101)
	spade_6   uint8 = color_spade | point_6   //黑桃6 (0011 0101)

	diamond_7 uint8 = color_diamond | point_7 //方块7 (0000 0110)
	club_7    uint8 = color_club | point_7    //梅花7 (0001 0110)
	heart_7   uint8 = color_heart | point_7   //红桃7 (0010 0110)
	spade_7   uint8 = color_spade | point_7   //黑桃7 (0011 0110)

	diamond_8 uint8 = color_diamond | point_8 //方块8 (0000 0111)
	club_8    uint8 = color_club | point_8    //梅花8 (0001 0111)
	heart_8   uint8 = color_heart | point_8   //红桃8 (0010 0111)
	spade_8   uint8 = color_spade | point_8   //黑桃8 (0011 0111)

	diamond_9 uint8 = color_diamond | point_9 //方块9 (0000 1000)
	club_9    uint8 = color_club | point_9    //梅花9 (0001 1000)
	heart_9   uint8 = color_heart | point_9   //红桃9 (0010 1000)
	spade_9   uint8 = color_spade | point_9   //黑桃9 (0011 1000)

	diamond_10 uint8 = color_diamond | point_10 //方块10 (0000 1001)
	club_10    uint8 = color_club | point_10    //梅花10 (0001 1001)
	heart_10   uint8 = color_heart | point_10   //红桃10 (0010 1001)
	spade_10   uint8 = color_spade | point_10   //黑桃10 (0011 1001)

	diamond_J uint8 = color_diamond | point_j //方块J (0000 1010)
	club_J    uint8 = color_club | point_j    //梅花J (0001 1010)
	heart_J   uint8 = color_heart | point_j   //红桃J (0010 1010)
	spade_J   uint8 = color_spade | point_j   //黑桃J (0011 1010)

	diamond_Q uint8 = color_diamond | point_q //方块Q (0000 1011)
	club_Q    uint8 = color_club | point_q    //梅花Q (0001 1011)
	heart_Q   uint8 = color_heart | point_q   //红桃Q (0010 1011)
	spade_Q   uint8 = color_spade | point_q   //黑桃Q (0011 1011)

	diamond_K uint8 = color_diamond | point_k //方块K (0000 1100)
	club_K    uint8 = color_club | point_k    //梅花K (0001 1100)
	heart_K   uint8 = color_heart | point_k   //红桃K (0010 1100)
	spade_K   uint8 = color_spade | point_k   //黑桃K (0011 1100)

	diamond_A uint8 = color_diamond | point_a //方块A (0000 1101)
	club_A    uint8 = color_club | point_a    //梅花A (0001 1101)
	heart_A   uint8 = color_heart | point_a   //红桃A (0010 1101)
	spade_A   uint8 = color_spade | point_a   //黑桃A (0011 1101)

	SMALL_JOKER uint8 = 0x4E //小王 (0100 1110)
	BIG_JOKER   uint8 = 0x4F //大王 (0100 1111)
)

//牌点数 (从小到大)
const (
	point_2          uint8 = 0x01 //2 (0000 0001)
	point_3          uint8 = 0x02 //3 (0000 0010)
	point_4          uint8 = 0x03 //4 (0000 0011)
	point_5          uint8 = 0x04 //5 (0000 0100)
	point_6          uint8 = 0x05 //6 (0000 0101)
	point_7          uint8 = 0x06 //7 (0000 0110)
	point_8          uint8 = 0x07 //8 (0000 0111)
	point_9          uint8 = 0x08 //9 (0000 1000)
	point_10         uint8 = 0x09 //10 (0000 1001)
	point_j          uint8 = 0x0A //J (0000 1010)
	point_q          uint8 = 0x0B //Q (0000 1011)
	point_k          uint8 = 0x0C //K (0000 1100)
	point_a          uint8 = 0x0D //A (0000 1101)
	point_small_king uint8 = 0x0E //小王 (0000 1110)
	point_big_king   uint8 = 0x0F //大王 (0000 1111)
)

//花色
const (
	color_diamond uint8 = 0x00 //方块 (0000 0000)
	color_club    uint8 = 0x10 //梅花 (0001 0000)
	color_heart   uint8 = 0x20 //红桃 (0010 0000)
	color_spade   uint8 = 0x30 //黑桃 (0011 0000)
)

//掩码
const (
	mask_color = 0xF0 //花色掩码 (1111 0000)
	mask_point = 0x0F //数值掩码 (0000 1111)
)
