package main

import (
	"context"
	"fmt"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Greet(name, helloText, regexp_text string) string {
	// 启动 boss_spider_main 协程
	resultChan := make(chan string)
	ws := make(chan string)
	go func() {
		if len(helloText) < 3 {
			resultChan <- "报错啦 看看是不是打招呼语忘记填啦"
			return
		} else if len(regexp_text) < 3 {
			resultChan <- "报错啦 看看是不是岗位关键词忘记填啦"
			return
		}

		go start_chrome_main(ws, resultChan)

		resultChan <- boss_spider_main(ws, helloText, regexp_text)

	}()

	// 设置 1 秒的超时时间
	select {
	case result := <-resultChan:
		// 如果 boss_spider_main 在 1 秒内返回,则返回结果
		return fmt.Sprintf(result)
	case <-time.After(1 * time.Second):
		// 如果 boss_spider_main 在 1 秒内没有返回,则返回启动成功的信息
		return fmt.Sprintf("启动成功,请在一分钟内完成操作")
	}
}
