package main

import "fmt"

type JobType int

const (
	DM JobType = iota
	CDC
	Dumpling
	Lightning
)

type Job interface {
	Work()
}

type DMJob struct {
	Job
}

type CDCJob struct {
	Job
}

type DumplingJob struct {
	Job
}

type LightningJob struct {
	Job
}

type JobFactory interface {
	CreateJob(JobType) Job
}

type StreamJobFactory struct {
}

type BatchJobFactory struct {
}

func (f *StreamJobFactory) CreateJob(t JobType) Job {
	switch t{
	case DM:
		return &DMJob{}
	default:
		return &CDCJob{}
	}
}

func (f *BatchJobFactory) CreateJob(t JobType) Job{
	switch t{
	case Dumpling:
		return &DumplingJob{}
	default:
		return &LightningJob{}
	}
}

func (m *DMJob) Work() {
	fmt.Println("dm job")
}

func (m *CDCJob) Work() {
	fmt.Println("cdc job")
}

func (m *DumplingJob) Work() {
	fmt.Println("dumpling job")
}

func (m *LightningJob) Work() {
	fmt.Println("lightning job")
}

func main() {
	streamJobFactory := &StreamJobFactory{}
	dmJob := streamJobFactory.CreateJob(DM)
	dmJob.Work()
	cdcJob := streamJobFactory.CreateJob(CDC)
	cdcJob.Work()

	batchJobFactory := &BatchJobFactory{}
	dumplingJob := batchJobFactory.CreateJob(Dumpling)
	dumplingJob.Work()
	lightningJob := batchJobFactory.CreateJob(Lightning)
	lightningJob.Work()
}
