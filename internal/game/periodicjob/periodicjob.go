package periodicjob

import (
	"time"
)

type PeriodicJob interface {
	SetInterval(interval int)
	Cancel()
}

type Impl struct {
	interval int
	callback func()
	canceled bool
}

func New(interval int, callback func()) *Impl {
	job := Impl{
		interval: interval,
		callback: callback,
	}
	go job.periodicExecute()
	return &job
}

func (job *Impl) periodicExecute() {
	for {
		startTime := time.Now()
		job.callback()
		elapsedTime := time.Now().Sub(startTime)
		time.Sleep(time.Duration(job.interval) - elapsedTime)
		if job.canceled {
			break
		}
	}
}

func (job *Impl) Cancel() {
	job.canceled = true
}

func (job *Impl) SetInterval(interval int) {
	job.interval = interval
}
