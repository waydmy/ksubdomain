package runner

import (
	util "github.com/GhostTroops/go-utils"
	"strings"
)

func (r *runner) handleResult() {
	go util.DoRunning()
	defer util.CloseLogBigDb()
	var szSkp = "0.0.0.1"
	//log.Println("r.options.Writer len:", len(r.options.Writer))
	wg := util.NewSizedWaitGroup(0)
	for result := range r.recver {
		if -1 < strings.Index(result.Subdomain, szSkp) {
			continue
		}
		for _, x := range result.Answers {
			if x == szSkp { // 跳过无效的
				return
			}
		}

		var m1 = map[string]interface{}{"ip": result.Answers, "subdomain": result.Subdomain, "tags": "subdomain"}
		//KvCc.KvCc.Put(result.Subdomain, []byte("1"))

		func(m09 *map[string]interface{}) {
			util.WaitFunc4Wg(&wg, func() {
				util.PushLog(&m09)
			})
		}(&m1)
		for _, out := range r.options.Writer {
			_ = out.WriteDomainResult(result)
		}
		r.printStatus()
	}
	wg.Wait()
}
