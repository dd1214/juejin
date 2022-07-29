package controller

/*
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	videoID := c.Query("video_id")

	dbInit()
	defer db.Close()
	var users []dbUser
	var videos []dbVideo
	//查询
	db.Select(&users, "select ID, Name, FollowCount, FollowerCount, IsFollow from User where token=?", token)
	db.Select(&videos, "select ID, PlayUrl, CoverUrl, FavoriteCount, CommentCount, IsFavorite, Title from Video where ID=?", videoID)

	if users != nil {
		if actionType == "1" {
			text := c.Query("comment_text")
			newID := makeId()
			_, month, day := time.Now().Date()
			createData := string(month) + "-" + string(day)

			db.Exec("update Video set CommentCount=? where ID=?", videos[0].CommentCount+1, videoID)
			db.Exec("insert into Comment(CommentText, CreateDate, ID, UserID, VideoID)value(?, ?, ?, ?, ?)", text, createData, newID, users[0].ID, videoID)

			var user = User{
				Id:            users[0].ID,
				Name:          users[0].Name,
				FollowCount:   users[0].FollowCount,
				FollowerCount: users[0].FollowerCount,
				IsFollow:      users[0].IsFollow,
			}

			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0, StatusMsg: "Comment success"},
				Comment: Comment{
					Id:         newID,
					User:       user,
					Content:    text,
					CreateDate: string(month) + "-" + string(day),
				}})
			return
		} else if actionType == "2" {
			commentID := c.Query("comment_id")
			db.Exec("update Video set CommentCount=? where ID=?", videos[0].CommentCount-1, videoID)
			db.Exec("delete from Comment where ID=?", commentID)

			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "Delete comment success"})
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoID := c.Query("video_id")
	dbInit()
	defer db.Close()
	//从数据库查询该视频的评论列表
	var dbComments []dbComment
	db.Select(&dbComments, "select ID, UserID, CommentText, CreateDate from Comment where VideoID=?", videoID)
	var comments []Comment
	//填充返回的评论列表
	for _, comment := range dbComments {
		//查询评论用户
		var user []dbUser
		db.Select(&user, "select ID, Name, FollowCount, FollowerCount, IsFollow from User where ID=?", comment.UserID)
		comments = append(comments, Comment{
			Id: comment.ID,
			User: User{
				Id:            user[0].ID,
				Name:          user[0].Name,
				FollowCount:   user[0].FollowCount,
				FollowerCount: user[0].FollowerCount,
				IsFollow:      user[0].IsFollow,
			},
			Content:    comment.CommentText,
			CreateDate: comment.CreateDate,
		})
	}

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0, StatusMsg: ""},
		CommentList: comments,
	})
}


*/
