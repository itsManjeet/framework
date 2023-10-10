package command

import "fmt"

func (c *Command) Help() error {
	fmt.Printf("%s: %s\n", c.selfPath, c.usage)
	fmt.Println(c.about)
	fmt.Printf("TASK:\n")
	for _, s := range c.subCommands {
		fmt.Printf("  %s%*.s %s\n", s.id, 20-len(s.id), " ", s.about)
	}
	fmt.Printf("\nFLAGS:\n")
	for _, f := range c.flags {
		fmt.Printf("  -%s%*.s%s\n", f.GetId(), 20-len(f.GetId()), " ", f.GetAbout())
	}
	return nil
}
