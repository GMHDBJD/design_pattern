package main

import "fmt"

type Job struct {
	master string
	worker string
	task   string
}

type JobBuilder interface {
	BuildMaster()
	BuildWorker()
	BuildTask()
	CreateJob()
	GetJob() *Job
}

type DMJobBuilder struct {
	dmJob *Job
	JobBuilder
}

type CDCJobBuilder struct {
	cdcJob *Job
	JobBuilder
}

func (b *DMJobBuilder) CreateJob() {
	b.dmJob = &Job{}
}

func (b *DMJobBuilder) BuildMaster() {
	b.dmJob.master = "dm-master"
}

func (b *DMJobBuilder) BuildWorker() {
	b.dmJob.worker = "dm-worker"
}

func (b *DMJobBuilder) BuildTask() {
	b.dmJob.task = "dm-task"
}

func (b *DMJobBuilder) GetJob() *Job {
	return b.dmJob
}

func (b *CDCJobBuilder) CreateJob() {
	b.cdcJob = &Job{}
}

func (b *CDCJobBuilder) BuildMaster() {
	b.cdcJob.master = "cdc-master"
}

func (b *CDCJobBuilder) BuildWorker() {
	b.cdcJob.worker = "cdc-worker"
}

func (b *CDCJobBuilder) BuildTask() {
	b.cdcJob.task = "cdc-task"
}

func (b *CDCJobBuilder) GetJob() *Job {
	return b.cdcJob
}

func main() {
	dmJobBuilder := &DMJobBuilder{}
	dmJobBuilder.CreateJob()
	dmJobBuilder.BuildMaster()
	dmJobBuilder.BuildWorker()
	dmJobBuilder.BuildTask()
	j := dmJobBuilder.GetJob()
	fmt.Println(j)

	cdcJobBuilder := &CDCJobBuilder{}
	cdcJobBuilder.CreateJob()
	cdcJobBuilder.BuildMaster()
	cdcJobBuilder.BuildWorker()
	cdcJobBuilder.BuildTask()
	j = cdcJobBuilder.GetJob()
	fmt.Println(j)
}
