package main

import (
	"bufio"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"log"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"
)

// 获取常量信息方法
func get_const_info() (string, string, string, string, string, string, time.Duration) {

	chat_speed := 3 * time.Second
	const next_page_element_xpath string = "//*[@class=\"ui-icon-arrow-right\"]"
	const job_list_elements_xpath string = "//*[@id=\"wrap\"]//li/div[1]/a/div[1]/span[1]"
	const chat_with_boss_elements_xpath string = "//*[@class=\"info-public\"]"
	const chat_input_element_xpath string = "//*[@id=\"chat-input\"]"
	const send_message_xpath string = "//*[@class=\"btn-v2 btn-sure-v2 btn-send\"]"
	const message_history_xpath string = "//*[@class=\"item-myself\"]"
	return next_page_element_xpath, job_list_elements_xpath, chat_with_boss_elements_xpath, chat_input_element_xpath, send_message_xpath, message_history_xpath, chat_speed
}

// BrowserController 结构体，包含一个rod浏览器实例
type BrowserController struct {
	Browser *rod.Browser
}

// NewBrowserController 用于创建BrowserController实例的函数
func NewBrowserController(url string) (*BrowserController, error) {
	// 使用提供的URL创建一个新的浏览器实例
	browser := rod.New().ControlURL(url).MustConnect()
	return &BrowserController{Browser: browser}, nil
}

// 打招呼方法
func page_message_send(bc *BrowserController, message_history_xpath, chat_input_element_xpath, hello_text, send_message_xpath string, chat_speed time.Duration) {
	pages, _ := bc.Browser.Pages()
	page := pages[0]
	message_history_elements, _ := page.ElementsX(message_history_xpath)
	//判断是否有消息发送记录
	if len(message_history_elements) == 0 {
		chat_input_element := page.MustElementX(chat_input_element_xpath)
		err := chat_input_element.Input(hello_text)
		if err != nil {
			fmt.Println(err.Error())
		}
		time.Sleep(chat_speed)

		send_message_element := page.MustElementX(send_message_xpath)
		send_message_element.MustClick()
		time.Sleep(chat_speed)
		page.MustNavigateBack()

	} else {
		page.MustNavigateBack()
	}

}

func boss_spider_main(ws chan string, hello_text, regexp_text string) string {
	u := <-ws

	//定义常量
	next_page_element_xpath, job_list_elements_xpath, chat_with_boss_elements_xpath, chat_input_element_xpath, send_message_xpath, message_history_xpath, chat_speed := get_const_info()

	boss_spider, _ := NewBrowserController(u)
	time.Sleep(5 * time.Second)
	boss_spider.Browser.MustPage().MustNavigate("https://www.zhipin.com/web/geek/job?query=")
	time.Sleep(60 * time.Second)

	////翻页次数
	for i := 0; i < 11; i++ {

		//获取页面元素信息
		pages, err := boss_spider.Browser.Pages()
		if err != nil {
			return "报错啦 看看是不是信息没填对呀"
		} else if len(pages) == 0 {
			// 如果没有错误，但是pages的长度为0，说明没有获取到页面
			return "报错啦，没有获取到页面信息 看看是不是信息没填对呀"
		}
		page := pages[0]
		job_list_elements, _ := page.ElementsX(job_list_elements_xpath)
		next_page_element, _ := page.ElementX(next_page_element_xpath)
		chat_with_boss_elements, _ := page.ElementsX(chat_with_boss_elements_xpath)

		//获取当前页面岗位信息
		for j := 0; j < len(job_list_elements); j++ {

			//重新定位
			pages, _ = boss_spider.Browser.Pages()
			page = pages[0]
			job_list_elements, _ = page.ElementsX(job_list_elements_xpath)
			chat_with_boss_elements, _ = page.ElementsX(chat_with_boss_elements_xpath)
			key_todo := page.Keyboard
			mouse_todo := page.Mouse

			//获取岗位和聊天信息
			job_name, _ := job_list_elements[j].Text()
			chat_with_boss_element := chat_with_boss_elements[j]

			if j > 16 {
				for n := 0; n < 4; n++ {
					key_todo.MustType(input.PageDown)
					time.Sleep(time.Second * 1)
					mouse_todo.MustMoveTo(1.2, 1.2)
				}

			}

			//判断岗位是否符合条件
			pattern := regexp.MustCompile(regexp_text)
			if pattern.MatchString(job_name) {
				fmt.Println(job_name, chat_with_boss_element)

				//点击跳转
				wait := page.MustWaitNavigation()
				chat_with_boss_element.MustClick()
				pages, _ = boss_spider.Browser.Pages()
				page = pages[0]
				wait()
				time.Sleep(chat_speed)

				//发送消息并返回
				page_message_send(boss_spider, message_history_xpath, chat_input_element_xpath, hello_text, send_message_xpath, chat_speed)
				time.Sleep(chat_speed)

			}
		}
		pages, _ = boss_spider.Browser.Pages()
		page = pages[0]
		next_page_element, _ = page.ElementX(next_page_element_xpath)
		next_page_element.MustClick()
		time.Sleep(chat_speed)

	}
	return ""

}

func start_chrome_main(wsChan, resultChan chan string) {
	// 创建一个字符串通道用于在协程之间传递 WebSocket URL

	chrome := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	if runtime.GOOS == "darwin" {
		chrome = "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
	}
	cmd := exec.Command(chrome, "--remote-debugging-port=9222")

	// 获取命令的标准输出管道
	stdoutPipe, _ := cmd.StdoutPipe()
	// 获取命令的标准错误输出管道
	stderrPipe, _ := cmd.StderrPipe()

	// 启动命令
	if err := cmd.Start(); err != nil {

		log.Printf("命令启动失败:", err)
		resultChan <- "浏览器启动失败,请检查浏览器启动路径是否在C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe  并以管理员身份运行此程序"

		return
	}

	// 创建一个 go 协程来读取标准输出
	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	// 创建另一个 go 协程来读取标准错误输出
	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			line := scanner.Text()
			// 检查行是否包含特定的前缀
			if strings.HasPrefix(line, "DevTools listening on") {
				// 使用 Split 函数分割字符串并获取 "on" 之后的部分
				parts := strings.Split(line, "on ")
				if len(parts) > 1 {
					// 打印 "on" 之后的部分
					wsChan <- parts[1]
					break // 如果您只想打印第一次出现的地址，请加上 break
				}
			}
		}
		close(wsChan) // 关闭通道，表示没有更多的值会被发送
	}()

	// 等待命令完成
	if err := cmd.Wait(); err != nil {
		fmt.Println("命令执行出错:", err)
	}

}
