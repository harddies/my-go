package main

import (
	"fmt"
	"github.com/gookit/event"
	"log"
)

type TaskStateDatabus struct {
	TaskId      string `json:"task_id" form:"task_id" validate:"required"`
	TaskType    int8   `json:"task_type" form:"task_type" validate:"required"`
	Progress    int8   `json:"progress" form:"progress"`
	State       int8   `json:"state" form:"state" validate:"required"`
	Remark      string `json:"remark" form:"remark"`
	BeforeState int8   `json:"before_state" form:"before_state" validate:"required"`
	Operate     string `json:"operate" form:"operate" validate:"required"`
}

func Handle1(e event.Event) (err error) {
	fmt.Printf("Handle1 func %+v\n", e.Get("data").(*TaskStateDatabus))
	return
}

func Handle2(e event.Event) (err error) {
	fmt.Printf("Handle2 func %+v\n", e.Get("data").(*TaskStateDatabus))
	return
}

func main() {
	t := &TaskStateDatabus{
		TaskId:      "task1",
		TaskType:    1,
		Progress:    10,
		State:       7,
		Remark:      "",
		BeforeState: 7,
		Operate:     "=",
	}
	event.On("test", event.ListenerFunc(Handle1), event.Normal)
	event.On("test", event.ListenerFunc(Handle2), event.Max)

	err, e := event.Fire("test", event.M{"data": t})
	if err != nil {
		log.Printf("event.Fire event(%+v) error(%+v)", e, err)
	}
}
