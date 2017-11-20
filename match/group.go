package match


//in match, several matcher compose the group
type Group interface {
	//add matcher in group
	Add(Matcher) 	error
	//current weight in group
	Weight() 	int

	//unique id as group key
	UniqueId()	string

	//return all matcher in group
	GetAllMatcher()	[]Matcher
}