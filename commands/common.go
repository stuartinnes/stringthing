package commands

const (
	//EXIT_NOK is used to indicate a failing test
	EXIT_NOK = iota + 1
	//EXIT_OTHER is used in all other abnormal exit conditions
	EXIT_OTHER

	//verbose ENV flag
	ENV_VERBOSE = "ST_VERBOSE"
)
