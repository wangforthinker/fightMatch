package match
//
//method: Player(); return weight value in match

type Player interface {
	//score for match,inherit Interface Scorer
	Score() 	int
	//unique id
	UniqueId()	string
}