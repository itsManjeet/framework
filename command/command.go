package command

import (
	"fmt"

	"github.com/itsmanjeet/framework/command/flag"
)

type Handler func(*Command, []string, interface{}) error

type InitMethod func() (interface{}, error)

type Command struct {
	id        string
	about     string
	shortName string
	usage     string
	selfPath  string
	handler   Handler

	flags       []*flag.Flag
	initMethod  InitMethod
	subCommands []*Command
}

func New(id string) *Command {
	return &Command{
		id:         id,
		initMethod: nil,
	}
}

func (c *Command) ShortName(shortName string) *Command {
	c.shortName = shortName
	return c
}

func (c *Command) About(about string) *Command {
	c.about = about
	return c
}

func (c *Command) Usage(usage string) *Command {
	c.usage = usage
	return c
}

func (c *Command) Handler(handler Handler) *Command {
	c.handler = handler
	return c
}

func (c *Command) Sub(sub *Command) *Command {
	c.subCommands = append(c.subCommands, sub)
	return c
}

func (c *Command) Flag(f *flag.Flag) *Command {
	c.flags = append(c.flags, f)
	return c
}

func (c *Command) Init(i InitMethod) *Command {
	c.initMethod = i
	return c
}

func (c *Command) handleFlag(args []string) (int, error) {
	for _, i := range c.flags {
		if "-"+i.GetId() == args[0] {
			if i.GetCount() > len(args[1:]) {
				return 0, fmt.Errorf("%s expect %d arguments but %d provided", i.GetId(), i.GetCount(), len(args[1:]))
			}
			if err := i.GetHandler()(args[1 : i.GetCount()+1]); err != nil {
				return 0, err
			}
			return i.GetCount(), nil
		}
	}
	return 0, fmt.Errorf("invalid flag %s", args[0])
}

func (c *Command) Handle(args []string, iface interface{}) error {
	return c.handler(c, args, iface)
}
