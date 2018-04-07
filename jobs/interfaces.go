package jobs

type Job interface {
	Run(args []string)
}
