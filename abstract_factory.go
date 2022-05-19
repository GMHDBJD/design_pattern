package main

import "fmt"

type Job struct {
	componetFactor ComponetFactory
}

func NewJob(c ComponetFactory) *Job {
	return &Job{
		componetFactor: c,
	}
}

func (j *Job) doJob() {
	master := j.componetFactor.CreateMaster()
	worker := j.componetFactor.CreateWorker()
	master.Do()
	worker.Do()
}

type Master interface {
	Do()
}

type DMMaster struct {
	Master
}

type CDCMaster struct {
	Master
}

type Worker interface {
	Do()
}

type DMWorker struct {
	Worker
}

type CDCWorker struct {
	Worker
}

type ComponetFactory interface {
	CreateMaster() Master
	CreateWorker() Worker
}

type DMFactory struct {
}

type CDCFactory struct {
}

func (f *DMFactory) CreateMaster() Master {
	return &DMMaster{}
}

func (f *CDCFactory) CreateMaster() Master {
	return &CDCMaster{}
}

func (f *DMFactory) CreateWorker() Worker {
	return &DMWorker{}
}

func (f *CDCFactory) CreateWorker() Worker {
	return &CDCWorker{}
}

func (f *DMMaster) Do() {
	fmt.Println("dm master")
}

func (f *CDCMaster) Do() {
	fmt.Println("cdc master")
}

func (f *CDCWorker) Do() {
	fmt.Println("cdc worker")
}

func (f *DMWorker) Do() {
	fmt.Println("dm worker")
}

func main() {
	job := NewJob(&DMFactory{})
	job.doJob()

	job = NewJob(&CDCFactory{})
	job.doJob()
}
