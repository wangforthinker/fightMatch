package match

import (
	"sync"
	"errors"
	"fmt"
	"github.com/Sirupsen/logrus"
)

//base score
type Scorer interface {
	Score()		int
}

type Matcher interface {
	Run() error
	AddPlayer(Player) error

	AddPlayers([]Player) error
	RemovePlayers([]Player) error
	RemoveGroups([]Group) error
}

var(
	ErrorPlayerHasAlready = errors.New("player has been already")

)


//default matcher
type MatcherExecutor struct {
	currentGroups		[]Group
	//pending players, not belong to group
	pendingPlayer		[]Player
	_playerMapper		map[string]Player
	_groupMapper		map[string]Group

	//differ
	_differ			Difference

	pendingPlayerMutex 	sync.Mutex
	curGroupsMutex		sync.Mutex
}

func (this *MatcherExecutor) AddPlayer(player Player) error {
	logrus.Debug(fmt.Sprintf("player %s adding", player.UniqueId()))

	return this.addPlayersWithLock([]Player{player})
}

func (this *MatcherExecutor) Run() {

}

func (this *MatcherExecutor) addPlayersWithLock(players []Player) error {
	this.pendingPlayerMutex.Lock()
	defer this.pendingPlayerMutex.Unlock()

	for _,p := range players{
		err := this.addPlayerWithoutLock(p)
		if(err != nil){
			return fmt.Errorf("player %s add error:%s",p.UniqueId(), err.Error())
		}
	}

	return nil
}

func (this *MatcherExecutor) addPlayerWithoutLock(player Player) error {
	_,ok := this._playerMapper[player.UniqueId()]

	if(ok){
		return ErrorPlayerHasAlready
	}

	this._playerMapper[player.UniqueId()] = player
	this.pendingPlayer = append(this.pendingPlayer, player)

	return nil
}