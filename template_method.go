package main

import "fmt"

type AutoResumeTemplateMethod interface {
	CanResume() bool
	Resume()
}

type DefaultAutoResume struct {
	AutoResumeTemplateMethod
}

func (d *DefaultAutoResume) CheckState() bool {
	fmt.Println("default check state")
	return true
}

func (d *DefaultAutoResume) Run() {
	if d.CheckState() && d.CanResume() {
		d.Resume()
	}
}

type SyncAutoResume struct {
	DefaultAutoResume
}

func (s *SyncAutoResume) CanResume() bool {
	fmt.Println("sync can resume")
	return true
}

func (s *SyncAutoResume) Resume() {
	fmt.Println("sync resume")
}

func main() {
	s := &SyncAutoResume{DefaultAutoResume: DefaultAutoResume{}}
	s.DefaultAutoResume.AutoResumeTemplateMethod = s
	s.Run()
}
