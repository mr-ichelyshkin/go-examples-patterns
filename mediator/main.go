package main

import "fmt"

//Mediator defines an interface for communicating with Colleague objects
type Mediator interface {
	Sync(switcher *Switcher)
	Add(switcher *Switcher)
}

// -- >

//Switcher is Colleague
type Switcher struct {
	State     bool
	_mediator Mediator
}

//NewSwitcher creates a new Switcher
func NewSwitcher(mediator Mediator) *Switcher {
	switcher := &Switcher{false, mediator}
	mediator.Add(switcher)
	return switcher
}

//Sync starts the mediator Sync function
func (s Switcher) Sync() {
	s._mediator.Sync(&s)
}

// -- >

//SyncMediator is ConcreteMediator
type SyncMediator struct {
	Switchers []*Switcher
}

//Sync synchronizes the state of all Colleague objects
func (sm *SyncMediator) Sync(switcher *Switcher) {
	for _, curSwitcher := range sm.Switchers {
		curSwitcher.State = switcher.State
	}
}

//Add append Colleague to the Mediator list
func (sm *SyncMediator) Add(switcher *Switcher) {
	sm.Switchers = append(sm.Switchers, switcher)
}

func main() {
	mediator := &SyncMediator{[]*Switcher{}}
	switcher1 := NewSwitcher(mediator)
	switcher2 := NewSwitcher(mediator)
	switcher3 := NewSwitcher(mediator)

	switcher1.State = true
	state2 := switcher2.State
	state3 := switcher3.State

	switcher1.Sync()
	state2 = switcher2.State
	state3 = switcher3.State

	fmt.Println(state2, state3)
}
