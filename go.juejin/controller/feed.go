package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type FeedResponse struct {
	Response
	ArticleList []Article `json:"video_list,omitempty"`
	//NextTime    int64     `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")

	dbInit()
	defer db.Close()
	var userLogin []dbUser
	//查询登录用户信息
	db.Select(&userLogin, "select ID, Name, FollowCount, FollowerCount, IsFollow from User where token=?", token)
	var articleList []Article
	//获取文章列表
	rows, _ := db.Query("select ID, AuthorID, PlayUrl, CoverUrl, FavoriteCount, CommentCount, IsFavorite, Title, Text from Article where ID>?", 0)
	//填充文章列表
	if rows != nil {
		for rows.Next() {
			var article dbArticle
			rows.Scan(&article.ID, &article.AuthorID, &article.Url, &article.FavoriteCount, &article.CommentCount, &article.IsFavorite, &article.Title, &article.PublishTime, &article.Text)
			//获取用户信息
			var users []dbUser
			db.Select(&users, "select ID, Name, FollowCount, FollowerCount, IsFollow from User where ID=?", article.AuthorID)
			//若用户已登录，则判断是否已关注文章作者，否则默认未关注
			if userLogin != nil {
				var follow []dbFollower
				db.Select(&follow, "select IsFollow from FollowList where FollowerID=? and UserID=?", userLogin[0].ID, users[0].ID)
				if follow != nil {
					users[0].IsFollow = true
				} else {
					users[0].IsFollow = false
				}
			} else {
				users[0].IsFollow = false
			}
			var user = User{
				Id:            users[0].ID,
				Name:          users[0].Name,
				FollowCount:   users[0].FollowCount,
				FollowerCount: users[0].FollowerCount,
				IsFollow:      users[0].IsFollow,
			}
			articleList = append([]Article{
				{
					Id:     article.ID,
					Author: user,
					Url:    article.Url,
					//PublishTime:   article.PublishTime,
					FavoriteCount: article.FavoriteCount,
					CommentCount:  article.CommentCount,
					IsFavorite:    article.IsFavorite,
					Text:          article.Text,
				},
			}, articleList...)
		}
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:    Response{StatusCode: 0, StatusMsg: ""},
		ArticleList: articleList,
		//NextTime:  time.Now().Unix(),
	})
}
