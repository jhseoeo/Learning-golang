package main

import "sync"

type ChannelScoreboardManager chan func(map[string]int)

// ----------------------------------------------------------------------------------

func NewChannelScoreboardManager() (ChannelScoreboardManager, func()) {
	ch := make(ChannelScoreboardManager)
	done := make(chan struct{})
	go scoreboardManager(ch, done)
	return ch, func() {
		close(done)
	}
}

func (csm ChannelScoreboardManager) Update(name string, val int) {
	csm <- func(m map[string]int) {
		m[name] = val
	}
}

func (csm ChannelScoreboardManager) Read(name string) (int, bool) {
	var out int
	var ok bool
	done := make(chan struct{})

	csm <- func(m map[string]int) {
		out, ok = m[name]
		close(done)
	}
	<-done
	return out, ok
}

// ----------------------------------------------------------------------------------

type MutexScoreboardManager struct {
	l          sync.RWMutex
	scoreboard map[string]int
}

func NewMutexScoreboardManager() *MutexScoreboardManager {
	return &MutexScoreboardManager{
		scoreboard: map[string]int{},
	}
}

func (msm *MutexScoreboardManager) Update(name string, val int) {
	msm.l.Lock()
	defer msm.l.Unlock()
	msm.scoreboard[name] = val
}

func (msm *MutexScoreboardManager) Read(name string) (int, bool) {
	msm.l.RLock()
	defer msm.l.RUnlock()
	val, ok := msm.scoreboard[name]
	return val, ok
}

// ----------------------------------------------------------------------------------

func scoreboardManager(in <-chan func(map[string]int), done <-chan struct{}) {
	scoreboard := map[string]int{}
	for {
		select {
		case <-done:
			return
		case f := <-in:
			f(scoreboard)
		}
	}
}
