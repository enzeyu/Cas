package agent

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
	"Cas/dataobj"
)
var Wg sync.WaitGroup

func Crawl(request dataobj.RequestItem) *dataobj.ReponseItem{
	defer Wg.Done()
	if request.Deepth > 2{
		return nil
	}
	request.ReqTime = time.Now()
	request.Url = preProcess(request.Url)
	request.IsCheck = true
	domain , _ := url.Parse(request.Url)
	//request.Domain = domain.Host
	request.Domain = strings.Replace(domain.Host, "www.", "", 1)
	request.Add_req(request)

	var reply = new(dataobj.ReponseItem)
	resp, err := http.Get(request.Url)
	if err != nil{
		fmt.Println("request this ", request.Url ,"occured a problem")
		fmt.Println(err)
		reply.RespCode = 404
		reply.RequestUrl = request.Url
		reply.RespUrls = []dataobj.RequestItem{}
		reply.Deepth = request.Deepth + 1
		reply.IsCheck = true
		reply.Domain = request.Domain
		reply.Add_res(*reply)
		reply_copy:= *reply
		reply_copy.RespUrls = []dataobj.RequestItem{}
		reply_frontend = append(reply_frontend,reply_copy)
		return reply
	}

	reply.RequestUrl = request.Url
	reply.RespCode = resp.StatusCode
	reply.RespTime = time.Since(request.ReqTime)

	//fmt.Println(reply.RespTime)
	defer resp.Body.Close() 	//函数结束后关闭Body
	body, err := ioutil.ReadAll(resp.Body)
	content := string(body)

	html, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil{
		fmt.Println("parse this ", request.Url ,"occured a problem")
		fmt.Println(err)
	}
	reply.RespUrls = getUrls(html,request)
	reply.Deepth = request.Deepth + 1
	reply.IsCheck = true
	reply.Domain = request.Domain
	reply.Add_res(*reply)
	reply_copy:= *reply
	reply_copy.RespUrls = []dataobj.RequestItem{}
	reply_frontend = append(reply_frontend,reply_copy)
	return reply
}

func getUrls(html *goquery.Document, request dataobj.RequestItem) []dataobj.RequestItem {
	newlist := []dataobj.RequestItem{}
	temp_url := []string{}
	temp_url = append(temp_url,request.Url)
	html.Find("a").Each(func(i int, selection *goquery.Selection) {  //此处选择器仍然可以优化
		url, _ := selection.Attr("href")

		url = parseUrls(url,request.Domain)    //对URL进行筛选
		if url != "" && url != request.Domain {
			temp := dataobj.RequestItem{}
			temp.Url = "http://" + url
			if isExist(temp.Url,temp_url) && temp.Domain != request.Domain{

			}else{
				temp.IsCheck = false
				temp.Domain = url
				newlist = append(newlist,temp)
				temp_url = append(temp_url,temp.Url)
			}

		}
	})
	return newlist
}

func parseUrls(s string,domain string) string{
	s = preProcess(s)
	//fmt.Println(s)
	u, err := url.Parse(s)
	if err != nil{
		return ""
	}
	if u.Host != "" &&  domainMatch(u.Host, domain){
		return u.Host
	} else{
		return ""
	}
}

func preProcess(s string) string{
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "\t", "", -1)

	return s
}

func domainMatch(main_s string,ass_s string)bool{
	r_main_s := []rune(main_s)
	r_ass_s  := []rune(ass_s)
	num := 0.0
	k := 1
	if len(r_main_s)>len(r_ass_s){
		for i:=len(r_ass_s)-1;i>0;i--{
			j:=len(r_main_s)-k
			if(r_main_s[j]==r_ass_s[i]){
				num = num + 1.0
			}
			k = k+1
		}
		if(num/float64(len(r_ass_s))>0.8){
			return true
		}else{
			return false
		}
	}else{
		for i:=len(r_main_s)-1;i>0;i--{
			j:=len(r_ass_s)-k
			if(r_main_s[i]==r_ass_s[j]){
				num = num + 1.0
			}
			k = k+1
		}
		if(num/float64(len(r_main_s))>0.8){
			return true
		}else{
			return false
		}
	}
	return false
}

func isExist(s string, array []string) bool{ //检查URL是否存在返回的列表里
	for _,v := range array{
		if v == s{
			return true
		}
	}
	return false
}



