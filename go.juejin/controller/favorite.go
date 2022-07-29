package controller

/*
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoID := c.Query("video_id")
	action_type := c.Query("action_type")

	dbInit()
	defer db.Close()
	var user []dbUser
	//查询用户是否存在
	db.Select(&user, "select ID, Name, FollowCount, FollowerCount, IsFollow from User where token=?", token)
	//获取视频信息，若视频不存在直接返回
	var videos []dbVideo
	db.Select(&videos, "select ID, PlayUrl, CoverUrl, FavoriteCount, CommentCount, IsFavorite, Title from Video where ID=?", videoID)
	if videos == nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Video not exist"})
		return
	}
	if user != nil {
		if action_type == "1" {
			db.Exec("update Video set FavoriteCount=? where ID=?", videos[0].FavoriteCount+1, videos[0].ID)
			db.Exec("update Video set IsFavorite=? where ID=?", true, videos[0].ID)
			db.Exec("insert into LikeList(UserID, VideoID)value(?, ?)", user[0].ID, videos[0].ID)

			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "Like success"})
		} else if action_type == "2" {
			db.Exec("update Video set FavoriteCount=? where ID=?", videos[0].FavoriteCount-1, videos[0].ID)
			db.Exec("update Video set IsFavorite=? where ID=?", false, videos[0].ID)
			db.Exec("delete from LikeList where UserID=? and VideoID=?", user[0].ID, videos[0].ID)

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
	var videoList []Video
	var users []dbUser
	db.Select(&users, "select ID, Name, FollowCount, FollowerCount, IsFollow from User where ID=?", userID)
	if users == nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "user not found",
			},
			VideoList: videoList,
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
	db.Select(&likeList, "select VideoID from LikeList where UserID=?", userID)
	for _, video := range likeList {
		//获取视频列表
		var videos []dbVideo
		db.Select(&videos, "select ID, PlayUrl, CoverUrl, FavoriteCount, CommentCount, IsFavorite, Title from Video where ID=?", video.VideoID)
		if videos == nil {
			c.JSON(http.StatusOK, VideoListResponse{
				Response: Response{
					StatusCode: 0,
					StatusMsg:  "video not found",
				},
				VideoList: videoList,
			})
			return
		}
		//填充响应视频列表
		videoList = append(videoList, Video{
			Id:            videos[0].ID,
			Author:        user,
			PlayUrl:       videos[0].PlayUrl,
			CoverUrl:      videos[0].CoverUrl,
			FavoriteCount: videos[0].FavoriteCount,
			CommentCount:  videos[0].CommentCount,
			IsFavorite:    videos[0].IsFavorite,
		})
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "",
		},
		VideoList: videoList,
	})
}


*/
