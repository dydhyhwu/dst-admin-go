package schedule

import (
	"dst-admin-go/utils/zip"
	"fmt"
	"time"
)

type BackupSchedule struct {
	cron  string
	state chan int
}

func NewBackSchedule(cron string) *BackupSchedule {
	backupSchedule := &BackupSchedule{
		cron:  cron,
		state: make(chan int),
	}
	backupSchedule.state <- 1
	return backupSchedule
}

func (b *BackupSchedule) StartSchedule() {
	go func() {
		for {
			select {
			case <-b.state:
				// 开始 定时任务
			}
		}
	}()
}
func (b *BackupSchedule) ReSchedule(cron string) {

}

func (b *BackupSchedule) TimingBackup(sourceDir, targetZip string, hour, minute int) {
	for {
		now := time.Now()
		year, month, day := now.Date()

		targetTime := time.Date(year, month, day, hour, minute, 0, 0, now.Location())

		if targetTime.Before(now) {
			targetTime = targetTime.Add(24 * time.Hour) // 如果目标时间已过，则设置为下一天的目标时间
		}

		timeToWait := targetTime.Sub(now)
		fmt.Printf("下一次压缩将在 %v 后进行\n", timeToWait)

		timer := time.NewTimer(timeToWait)
		<-timer.C

		err := zip.Zip(sourceDir, targetZip)
		if err != nil {
			fmt.Println("压缩目录时发生错误:", err)
		} else {
			fmt.Println("目录已成功压缩！")
		}
	}
}
