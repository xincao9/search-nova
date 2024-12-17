package crawler

import (
	"github.com/panjf2000/ants/v2"
	"search-nova/internal/constant"
	"search-nova/internal/logger"
	"search-nova/service/page"
	"sync"
	"sync/atomic"
	"time"
)

var (
	C *crawler
)

type crawler struct {
	running atomic.Bool
	pool    *ants.Pool
}

func init() {
	var err error
	C, err = new()
	if err != nil {
		logger.L.Fatalf("Fatal error crawler: %v\n", err)
	}
}
func new() (*crawler, error) {
	c := &crawler{}
	var err error
	c.running.Store(false)
	c.pool, err = ants.NewPool(4)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *crawler) Refresh() error {
	maxId, err := page.P.MaxId()
	if err != nil {
		return err
	}
	var id int64 = 1
	var wg sync.WaitGroup
	for ; id <= maxId; id++ {
		logger.L.Infof("crawler.Refresh(): %d/%d\n", id, maxId)
		p, err := page.P.GetPageById(id)
		if err != nil {
			logger.L.Errorf("page.GetPageById(%d) err %v\n", id, err)
			continue
		}
		if p == nil {
			continue
		}
		if p.Status != constant.NewStatus {
			continue
		}
		wg.Add(1)
		err = c.pool.Submit(func() {
			defer wg.Done()
			err = page.P.TextAnalysis(p)
			if err == nil {
				return
			}
			logger.L.Errorf("crawler.TextAnalysis(%s) err %v\n", p.Url, err)
			p.Status = constant.FailureStatus
			err = page.P.Save(p)
			if err != nil {
				logger.L.Errorf("page.Save(%v) err %v\n", p, err)
			}
		})
	}
	wg.Wait()
	return nil
}

func (c *crawler) Loop() {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		if !c.running.CompareAndSwap(false, true) {
			logger.L.Infoln("crawler.Refresh() running")
			continue
		}
		go func() {
			defer c.running.Store(false)
			err := c.Refresh()
			if err != nil {
				logger.L.Errorf("crawler.Refresh() err: %v\n", err)
			}
		}()
	}
}
