package abstractfactory

// ExampleAbstractFactory 委托者(client)
func ExampleAbstractFactory() {
	factory := GetFactory("ListFactory")
	people := factory.CreateLink("人民日报", "http://www.people.com.cn/")
	gmw := factory.CreateLink("光明日报", "http://www.gmw.cn/")
	us_yahoo := factory.CreateLink("Yahoo!", "http://www.yahoo.com/")
	jp_yahoo := factory.CreateLink("Yahoo!Japan", "http://www.yahoo.co.jp/")
	excite := factory.CreateLink("Excite", "http://www.excite.com/")
	google := factory.CreateLink("Google", "http://www.google.com/")

	trayNews := factory.CreateTray("日报")
	trayNews.Add(people)
	trayNews.Add(gmw)

	trayYahoo := factory.CreateTray("Yahoo!")
	trayYahoo.Add(us_yahoo)
	trayYahoo.Add(jp_yahoo)

	traySearch := factory.CreateTray("检索引擎")
	traySearch.Add(trayYahoo)
	traySearch.Add(excite)
	traySearch.Add(google)

	page := factory.CreatePage("LinkPage", "胡启航")
	page.Add(trayNews)
	page.Add(traySearch)
	page.Output()

	// Output:
	// LinkPage.html 编写完成。
}
