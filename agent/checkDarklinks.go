package agent

import (
	"fmt"
	"regexp"
	"Cas/dataobj"
)

func CheckDarklinks() []string{
	var return_lists = []string{}
	content := `222222 <div id="1">i love liu lei</div>  <script>`
	CheckReg := dataobj.CheckReg
	fmt.Println(CheckReg,len(CheckReg))
	for _,reg_rule := range CheckReg{
		reg := regexp.MustCompile(reg_rule)
		if reg.MatchString(content) {
			return_lists = append(return_lists, reg_rule)
			continue
		}
	}
	return return_lists
}
