package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"time"
)

func main() {
	//初始化应用程序
	fmt.Println(`欢迎使用学习通课表获取程序，请检查自己的电脑是否安装chrome。
- 已安装Chrome，请回车。
- 未安装Chrome，请在以下网址下载安装后回车：
https://www.google.cn/intl/zh-CN/chrome/`)
	var a string
	_, _ = fmt.Scanln(&a)
	var userData = new(UserData)
	fmt.Println("为进行学习通登录操作，请输入手机号和密码。【用户名及密码不会发送至除学习通及教务系统外的任何网站。程序也不会保存】")
	fmt.Print("请输入手机号：")
	_, _ = fmt.Scanln(&userData.UserName)
	fmt.Print("请输入密码：")
	_, _ = fmt.Scanln(&userData.Pwd)
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:], chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", true),
		chromedp.Flag("ignore-certificate-errors", true), //忽略错误
		chromedp.Flag("disable-web-security", true),      //禁用网络安全标志
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`))
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(
		allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	var html string
	var id string
	fmt.Println("开始获取用户信息...")
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate("http://service.wut.edu.cn/login"),
		chromedp.WaitVisible(`#uname`, chromedp.ByID),
		chromedp.SendKeys(`#uname`, userData.UserName, chromedp.ByID),
		chromedp.SendKeys(`#pwd`, userData.Pwd, chromedp.ByID),
		chromedp.Click(`#form > div.item.loginBtn > input[type=button]`, chromedp.ByID),
		chromedp.Sleep(4 * time.Second),
		chromedp.Click(`body > div > div.mainTopCon.marginTop > div.often.backgroundRadius.topHeight > ul > li:nth-child(2) > a`, chromedp.BySearch),
		chromedp.Sleep(5 * time.Second),
		chromedp.Navigate("http://jwxt.wut.edu.cn/admin/pkgl/xskb/queryKbForXsd?xnxq=2022-2023-1"),
		chromedp.WaitNotVisible(`#xhid`, chromedp.ByID),
		chromedp.Value(`#xhid`, &id, chromedp.ByID),
	})
	fmt.Println("开始获取课表信息...")
	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(fmt.Sprintf("http://jwxt.wut.edu.cn/admin/api/getXskb?xnxq=2022-2023-1&userId=%s&xqid=&role=xs", id)),
		chromedp.Sleep(2 * time.Second),
		chromedp.Text(`body > pre`, &html, chromedp.ByQuery),
	})
	if err != nil {
		return
	}
	fmt.Println("获取完成")
	err = os.WriteFile("./data.json", []byte(html), 0777)
	if err != nil {
		return
	}
	fmt.Println("获取的数据已经放在本地文件[data.json]中...输入回车结束程序")
	fmt.Scanln(&a)
	//var courseData = new(CourseData)
	//
	//_ = json.Unmarshal([]byte(html), &courseData)
	//
	//if err != nil {
	//	return
	//}
}
