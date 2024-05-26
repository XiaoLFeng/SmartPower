package startup

import (
	"context"
	"fmt"
)

// SystemPreparationFinal
//
// 初始化系统 - 系统准备完成
func SystemPreparationFinal(ctx context.Context) {
	fmt.Println("\033[1;35m" + `
	 ____                       _   ____
	/ ___| _ __ ___   __ _ _ __| |_|  _ \ _____      _____ _ __
	\___ \| '_ ` + "`" + ` _ \ / _` + "`" + ` | '__| __| |_) / _ \ \ /\ / / _ \ '__|
	 ___) | | | | | | (_| | |  | |_|  __/ (_) \ V  V /  __/ |
	|____/|_| |_| |_|\__,_|_|   \__|_|   \___/ \_/\_/ \___|_|`)
	fmt.Println("\033[32m	  	::: SmartPower :::\033[33m		v1.0.0")
	fmt.Println("\033[0m")
}
