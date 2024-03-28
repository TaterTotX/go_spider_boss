package main

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"regexp"
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

func boss_spider_main(u, hello_text, regexp_text string) string {

	if len(u) < 10 {
		return "报错啦 看看是不是ws值没填对呀"
	}
	//定义常量
	next_page_element_xpath, job_list_elements_xpath, chat_with_boss_elements_xpath, chat_input_element_xpath, send_message_xpath, message_history_xpath, chat_speed := get_const_info()

	boss_spider, err := NewBrowserController(u)

	if err != nil {
		return "报错啦 看看是不是ws值没填对呀"
	} else if len(hello_text) < 3 {
		return "报错啦 看看是不是打招呼语忘记填啦"
	} else if len(regexp_text) < 3 {
		return "报错啦 看看是不是岗位关键词忘记填啦"
	}

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
