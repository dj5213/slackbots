package bot

type AvailCmdsData struct {
	Commands []Command
}

type Command struct {
	Name        string
	Command     string
	Description string
}

func getCmds() AvailCmdsData {
	cmds := AvailCmdsData{}

	mrs := Command{
		Name:    "open merge requests",
		Command: "<open mrs, get open requests>",
		Description: "Basically any text with \"open mr\" or " +
			"\"open merge request\" will trigger the " +
			"bot to get stats on open merge requests",
	}
	cmds.Commands = append(cmds.Commands, mrs)
	return cmds
}
