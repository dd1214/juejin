package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ArticleListResponse struct {
	Response
	ArticleList []Article `json:"article_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	ID := c.Query("id")

	//生成id并获取投稿时间
	newID := makeId()
	publishTime := time.Now().Unix()
	dbInit()
	defer db.Close()
	var users []dbUser
	//查询
	db.Select(&users, "select ID, Name from User where ID=?", ID)
	if users == nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	//生成url
	newUrlID := makeId()
	Url := "http://localhost/news?" + string(newUrlID)

	//获取文章标题及文章内容
	title := c.Query("title")
	text := c.Query("text")
	db.Exec("insert into Article(ID, AuthorID, Url, FavoriteCount, CommentCount, IsFavorite, Title, PublishTime, Text)value(?, ?, ?, ?, ?, ?, ?, ?, ?)", newID, users[0].ID, Url, 0, 0, 0, title, publishTime, text)
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  title + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	userID := c.Query("user_id")

	dbInit()
	defer db.Close()
	var articleList []Article
	//获取用户信息
	var user []dbUser
	db.Select(&user, "select ID, Name, FollowCount, FollowerCount, IsFollow from User where ID=?", userID)
	if user == nil {
		return
	}
	var ResponseUser = User{
		Id:            user[0].ID,
		Name:          user[0].Name,
		FollowCount:   user[0].FollowCount,
		FollowerCount: user[0].FollowerCount,
		IsFollow:      user[0].IsFollow,
	}
	//获取视频列表
	var articles []dbArticle
	db.Select(&articles, "select ID, AuthorID, Url, FavoriteCount, CommentCount, IsFavorite, Title, PublishTime from Article where AuthorID=?", user[0].ID)
	//填充视频列表
	for _, article := range articles {
		articleList = append(articleList, Article{
			Id:     article.ID,
			Author: ResponseUser,
			Url:    article.Url,
			//PublishTime:   article.PublishTime,
			FavoriteCount: article.FavoriteCount,
			CommentCount:  article.CommentCount,
			IsFavorite:    article.IsFavorite,
			Text:          article.Text,
			Title:         article.Title,
		})
	}
	c.JSON(http.StatusOK, ArticleListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "",
		},
		ArticleList: articleList,
	})
}
