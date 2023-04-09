package main

import (
	"fmt"
	"github.com/looplab/fsm"
)

type State struct {
	Data struct{}
	FSM  *fsm.FSM
}

func NewState(s struct{}) (state *State) {
	state = &State{
		Data: s,
		FSM: fsm.NewFSM(
			"pending",
			fsm.Events{
				{Name: "confirm", Src: []string{"pending"}, Dst: "confirm"},
				{Name: "remove", Src: []string{"pending", "confirm"}, Dst: "remove"},
				{Name: "white", Src: []string{"pending", "remove"}, Dst: "white"},
			},
			fsm.Callbacks{
				"enter_confirm": state.enterConfirm,
				"enter_remove":  state.enterRemove,
				"enter_white":   state.enterWhite,
			},
		),
	}
	return
}

func (s *State) enterConfirm(event *fsm.Event) {
	fmt.Printf("it's confirmed, before: %s, after: %s\n", event.Src, event.Dst)
}

func (s *State) enterRemove(event *fsm.Event) {
	fmt.Printf("it's removed, before: %s, after: %s\n", event.Src, event.Dst)
}

func (s *State) enterWhite(event *fsm.Event) {
	fmt.Printf("it's white, before: %s, after: %s\n", event.Src, event.Dst)
}

func main() {
	state := NewState(struct{}{})

	err := state.FSM.Event("confirm")
	if err != nil {
		fmt.Println(err)
	}

	err = state.FSM.Event("remove")
	if err != nil {
		fmt.Println(err)
	}
}
