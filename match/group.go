package match


//in match, several matcher compose the group

type Group interface {
	//add matcher in group
	Add(Player) 	error
	//current score in group,inherit Interface Scorer
	Score() 	int

	//unique id as group key
	UniqueId()	string

	//return all players in group
	GetAllPlayers()	[]Player
}