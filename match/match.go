package match

import "sync"

//base interface
//method: weight(); return weight value in match

type Matcher interface {
	//weight for match
	Weight() 	int
	//unique id
	UniqueId()	string
}


//match
type MatchExecutor struct {
	currentGroups		[]Group
	pendingMatcher		[]Matcher


	pendingMatcherMutex 	sync.Mutex
	curGroupsMutex		sync.Mutex
}

func (this *MatchExecutor) Add([]Matcher) error {
	return nil
}

func (this *MatchExecutor) Run() {

}


