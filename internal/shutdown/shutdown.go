package shutdown

import (
	"os"
	"os/signal"
	"search-nova/internal/logger"
	"syscall"
)

var (
	S *shutdown
)

func init() {
	S = new()
	S.watch()
}

func new() *shutdown {
	return &shutdown{
		done: make(chan bool, 1),
	}
}

type shutdown struct {
	funcs []func()
	done  chan bool
}

func (s *shutdown) Await() {
	<-s.done
}

func (s *shutdown) Add(f func()) {
	s.funcs = append(s.funcs, f)
}

func (s *shutdown) watch() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		logger.L.Infof("shutdown.watch() 捕捉到 %v 信号\n", sig)
		total := len(s.funcs)
		for no, f := range s.funcs {
			logger.L.Infof("执行第 %d/%d 个终止任务\n", no+1, total)
			f()
		}
		s.done <- true
		os.Exit(0)
		logger.L.Infof("shutdown.watch() 完成\n")
	}()
}
