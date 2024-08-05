package periodicjob

import (
	"time"
)

type PeriodicJob interface {
	SetInterval(interval time.Duration)
	Cancel()
}

type Impl struct {
	interval time.Duration
	callback func()
	canceled bool
}

func New(interval time.Duration, callback func()) *Impl {
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
		time.Sleep(job.interval - elapsedTime)
		if job.canceled {
			break
		}
	}
}

func (job *Impl) Cancel() {
	job.canceled = true
}

func (job *Impl) SetInterval(interval time.Duration) {
	job.interval = interval
}
