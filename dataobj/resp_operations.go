package dataobj

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/rpcx-master/log"
	"Cas/db"
)

func return_req() string{
	return "request_"
}

func return_resp() string{
	return "response_"
}

func (this *RequestItem) Add_req(req RequestItem) error{
	c := db.Redgo.Get()
	value,err := json.Marshal(req)
	if err != nil {
		log.Errorf("request json marshal err,%s", err)
		return err
	}
	_  , err = c.Do("set",return_req() + req.Url, value)
	if err != nil {
		fmt.Println("request set failed", err.Error())
		return err
	}
	return nil
}

func (this *ReponseItem) Add_res(res ReponseItem) error{
	c := db.Redgo.Get()
	value,err := json.Marshal(res)
	if err != nil {
		log.Errorf("response json marshal err,%s", err)
		return err
	}
	_  , err = c.Do("set",return_resp() + res.RequestUrl, value)
	if err != nil {
		fmt.Println("response set failed", err.Error())
		return err
	}
	return nil
}

func (this *RequestItem) Get_req(s string) (err error, item RequestItem){
	req := RequestItem{}
	c := db.Redgo.Get()
	val , err := c.Do("get",return_req() + s)
	//fmt.Printf("%s",content)
	t := gconv.Bytes(val)
	err = json.Unmarshal(t,&req)
	if err != nil {
		fmt.Println("unmarshal failed", err.Error())
		return err, req
	}
	return nil,req
}

/*func Testdb(t *testing.T){
	req := RequestItem{}
	e := req.("https://www.seu.edu.cn/")

}*/

func (this *RequestItem) Get_res(s string) (err error, item ReponseItem){
	res := ReponseItem{}
	c := db.Redgo.Get()
	val  , err := c.Do("get",return_resp() + s)
	t := gconv.Bytes(val)
	err = json.Unmarshal(t,&res)
	if err != nil {
		fmt.Println("unmarshal failed", err.Error())
		return err, res
	}
	return nil,res
}


