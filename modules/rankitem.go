package modules

type RankItem struct {
	Id     int64
	Score  int64
	Time   int64
	Player *playerInfo
}
