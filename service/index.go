package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/models"
)

func GetAllIndexInfo() (*models.HomeResponse, error) {
	categories, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}

	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "张三",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &models.HomeResponse{
		Viewer:     config.Cfg.Viewer,
		Categories: categories,
		Posts:      posts,
		Total:      1,
		Page:       1,
		Pages:      []int{1},
		PageEnd:    true,
	}

	return hr, nil
}
