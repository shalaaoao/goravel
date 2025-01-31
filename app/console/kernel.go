package console

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/schedule"
	"goravel/app/console/commands"
)

type Kernel struct {
}

func (kernel *Kernel) Schedule() []schedule.Event {
	return []schedule.Event{
		//facades.Schedule().Command("command:test1").EveryMinute().OnOneServer(),
	}
}

func (kernel *Kernel) Commands() []console.Command {

	// 所有Command都需要在这里注册
	return []console.Command{
		&commands.TestCommand{},
		&commands.ShaFile{},
	}
}
