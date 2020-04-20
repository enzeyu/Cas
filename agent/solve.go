package agent

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"Cas/dataobj"
)

var reply_frontend []dataobj.ReponseItem

func Get_Result(context *gin.Context){
	receive_url := context.PostForm("url")
	receive_url = strings.Replace(receive_url,"\"","",-1)

	request := &dataobj.RequestItem{receive_url,time.Now(),false,1,""}

	Wg.Add(1)
	replys := Crawl(*request)
	Wg.Add(len(replys.RespUrls))

	for _ , j := range replys.RespUrls{
		go Crawl(j)
	}
	Wg.Wait()

	//_, temp_res := request.Get_res(request.Url)
	fmt.Println("reply_frontend",len(reply_frontend))
	context.JSON(200, gin.H{"message": reply_frontend})
	reply_frontend = []dataobj.ReponseItem{}
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}


