package views

import (
	"errors"
	"go-blog/common"
	"go-blog/service"
	"log"
	"net/http"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	hr, err := service.GetAllIndexInfo()
	if err != nil {
		log.Println("首页获取数据出错: ", err)
		index.WriteError(w, errors.New("please call Mason"))
	}
	index.WriteData(w, hr)
}
