package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FavoriteAction(c *gin.Context) {
	userID := c.Query("user_id")
	articleID := c.Query("article_id")
	action_type := c.Query("action_type")

	dbInit()
	defer db.Close()
	var user []dbUser
	//查询用户是否存在
	db.Select(&user, "select ID, Name, FollowCount, FollowerCount, IsFollow from User where ID=?", userID)
	//获取投稿信息，若不存在直接返回
	var articles []dbArticle
	db.Select(&articles, "select ID, PlayUrl, CoverUrl, FavoriteCount, CommentCount, IsFavorite, Title, Text from Video where ID=?", articleID)
	if articles == nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Video not exist"})
		return
	}
	if user != nil {
		if action_type == "1" {
			db.Exec("update Aticle set FavoriteCount=? where ID=?", articles[0].FavoriteCount+1, articles[0].ID)
			db.Exec("update Aticle set IsFavorite=? where ID=?", true, articles[0].ID)
			db.Exec("insert into LikeList(UserID, AticleID)value(?, ?)", user[0].ID, articles[0].ID)

			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "Like success"})
		} else if action_type == "2" {
			db.Exec("update Aticle set FavoriteCount=? where ID=?", articles[0].FavoriteCount-1, articles[0].ID)
			db.Exec("update Aticle set IsFavorite=? where ID=?", false, articles[0].ID)
			db.Exec("delete from LikeList where UserID=? and AticleID=?", user[0].ID, articles[0].ID)

			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "Unlike success"})
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

func FavoriteList(c *gin.Context) {
	userID := c.Query("user_id")

	dbInit()
	defer db.Close()
	//获取用户信息
	var articleList []Article
	var users []dbUser
	db.Select(&users, "select ID, Name, FollowCount, FollowerCount, IsFollow from User where ID=?", userID)
	if users == nil {
		c.JSON(http.StatusOK, ArticleListResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "user not found",
			},
			ArticleList: articleList,
		})
		return
	}
	var user = User{
		Id:            users[0].ID,
		Name:          users[0].Name,
		FollowCount:   users[0].FollowCount,
		FollowerCount: users[0].FollowerCount,
		IsFollow:      users[0].IsFollow,
	}
	//从数据库获取点赞列表
	var likeList []dbLike
	db.Select(&likeList, "select ArticleID from LikeList where UserID=?", userID)
	for _, article := range likeList {
		//获取视频列表
		var articles []dbArticle
		db.Select(&articles, "select ID, AuthorID, Url, FavoriteCount, CommentCount, IsFavorite, Title, PublishTime, Text from Article where ID=?", article.ArticleID)
		if articles == nil {
			c.JSON(http.StatusOK, ArticleListResponse{
				Response: Response{
					StatusCode: 0,
					StatusMsg:  "video not found",
				},
				ArticleList: articleList,
			})
			return
		}
		//填充响应视频列表
		articleList = append(articleList, Article{
			Id:     articles[0].ID,
			Author: user,
			Url:    articles[0].Url,
			//PublishTime:   articles[0].PublishTime,
			FavoriteCount: articles[0].FavoriteCount,
			CommentCount:  articles[0].CommentCount,
			IsFavorite:    articles[0].IsFavorite,
			Text:          articles[0].Text,
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
