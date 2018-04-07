package jobs

type Job interface {
	Run(args []string)
}

type JobResult struct {
	Status string
}
