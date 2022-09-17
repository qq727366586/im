package time

import (
	"log"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := NewHeapTimer(100)
	tds := make([]*TimerData, 100)
	for i := 0; i < 100; i++ {
		tds[i] = timer.Add(time.Duration(i)*time.Second+5*time.Minute, nil)
	}
	printTimer(timer)
	for i := 0; i < 100; i++ {
		timer.Del(tds[i])
	}
	printTimer(timer)
	for i := 0; i < 100; i++ {
		tds[i] = timer.Add(time.Duration(i)*time.Second+5*time.Minute, nil)
	}
	printTimer(timer)
	for i := 0; i < 100; i++ {
		timer.Del(tds[i])
	}
	printTimer(timer)
	timer.Add(time.Second, nil)
	time.Sleep(time.Second * 2)
	if len(timer.timers) != 0 {
		t.FailNow()
	}
}

func printTimer(timer *HeapTimer) {
	log.Printf("----------timers: %d ----------", len(timer.timers))
	log.Println()
	for i := 0; i < len(timer.timers); i++ {

		log.Printf("timer: %s, %s, index: %d", timer.timers[i].key, timer.timers[i].expire.Format("2006-01-02 15:03:04"), timer.timers[i].index)
		log.Println()

	}
	log.Printf("--------------------")
	log.Println()

}
