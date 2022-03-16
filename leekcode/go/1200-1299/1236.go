package _200_1299



/*
1236. 网络爬虫
给定一个链接 startUrl 和一个接口 HtmlParser ，请你实现一个网络爬虫，以实现爬取同 startUrl 拥有相同 域名标签 的全部链接。该爬虫得到的全部链接可以 任何顺序 返回结果。
你的网络爬虫应当按照如下模式工作：
自链接 startUrl 开始爬取
调用 HtmlParser.getUrls(url) 来获得链接url页面中的全部链接
同一个链接最多只爬取一次
只输出 域名 与 startUrl 相同 的链接集合
*/

func crawl(startUrl string, htmlParser HtmlParser) []string {
	var m = map[string]struct{}{}

	u,_:=url.Parse(startUrl)
	dfs(startUrl,htmlParser,m,u.Host)
	var result []string
	for k:=range m{
		result = append(result,k)
	}

	return result
}

func dfs(startUrl string,htmlParser HtmlParser,m map[string]struct{},host string){
	if _,ok:=m[startUrl];!ok{
		m[startUrl] = struct{}{}
		for _,v:=range htmlParser.GetUrls(startUrl){
			if strings.Index(v,host) >=0 {
				dfs(v, htmlParser, m,host)
			}
		}
	}
}
