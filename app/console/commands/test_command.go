package commands

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
)

type TestCommand struct {
}

// Signature The name and signature of the console command.
func (receiver *TestCommand) Signature() string {
	return "command:test1"
}

// Description The console command description.
func (receiver *TestCommand) Description() string {
	return "test"
}

// Extend The console command extend.
func (receiver *TestCommand) Extend() command.Extend {
	return command.Extend{}
}

// Handle Execute the console command.
func (receiver *TestCommand) Handle(ctx console.Context) error {

	//ctx.Info("hello world~")
	//
	//tools.ExecTime(func() interface{} {
	//
	//	facades.Cache().Store("redis").Put("bbb", 1, 3*time.Second)
	//
	//	return nil
	//
	//}, "flag")
	//
	//tools.ExecTime(func() interface{} {
	//
	//	facades.Cache().Put("aaa", 1, 100*time.Second)
	//
	//	return nil
	//
	//}, "flag")

	return nil
}
