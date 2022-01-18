package initialize

import (
	"fmt"

	"github.com/pddzl/kubeops/server/config"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/utils"
)

func Timer() {
	if global.KOP_CONFIG.Timer.Start {
		for i := range global.KOP_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				global.KOP_Timer.AddTaskByFunc("ClearDB", global.KOP_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.KOP_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(global.KOP_CONFIG.Timer.Detail[i])
		}
	}
}
