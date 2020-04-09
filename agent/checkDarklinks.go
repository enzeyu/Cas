package agent

import (
	"Cas/dataobj"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)


func Get_Check(context *gin.Context) {
	return_lists := []string{}
	receive_url := context.PostForm("url")
	resp, err := http.Get(receive_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get wrong: %v\n", err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	CheckReg := dataobj.CheckReg
	//fmt.Println(CheckReg,len(CheckReg))
	for _,reg_rule := range CheckReg{
		reg := regexp.MustCompile(reg_rule)
		if reg.MatchString(string(content)) {
			return_lists = append(return_lists, reg_rule)
			continue
		}
	}
	context.JSON(200, gin.H{"message": return_lists})
}
