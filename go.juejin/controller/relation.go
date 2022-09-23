package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

func RelationAction(c *gin.Context) {
	userID := c.Query("user_id")
	actionType := c.Query("action_type") // 1表示关注，2表示取消关注
	toUserID := c.Query("to_user_id")
	dbInit()
	defer db.Close()
	var user []dbUser
	var toUser []dbUser
	//查询用户是否存在
	db.Select(&user, "select ID, Name, FollowCount, FollowerCount, IsFollow from User where ID=?", userID)
	db.Select(&toUser, "select ID, Name, FollowCount, FollowerCount, IsFollow from User where ID=?", toUserID)

	//若对自己操作，则无效
	if user[0].ID == toUser[0].ID {
		return
	}
	if user != nil && toUser != nil {
		if actionType == "1" {
			db.Exec("update User set IsFollow=? where ID=?", true, toUser[0].ID)
			//判断是否是已关注的，若是则直接返回
			var users []dbFollower
			db.Select(&users, "select Name from FollowList where UserID=? and FollowerID=?", toUser[0].ID, user[0].ID)
			if users != nil {
				c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "You have followed this user before"})
				return
			}

			//修改用户关注数和粉丝数，并在FollowList新增一行
			db.Exec("update User set FollowerCount=? where ID=?", toUser[0].FollowerCount+1, toUser[0].ID)
			db.Exec("update User set FollowCount=? where ID=?", user[0].FollowCount+1, user[0].ID)
			db.Exec("insert into FollowList(FollowCount, FollowerCount, FollowerID, UserID, IsFollow, Name, token)value(?, ?, ?, ?, ?, ?, ?)", 0, 0, user[0].ID, toUserID, 1, toUser[0].Name, "1")
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "Follow success"})
		} else if actionType == "2" {
			db.Exec("update User set IsFollow=? where ID=?", false, toUser[0].ID)
			//修改用户关注数和粉丝数，并在FollowList删除对应行
			db.Exec("update User set FollowerCount=? where ID=?", toUser[0].FollowerCount-1, toUser[0].ID)
			db.Exec("update User set FollowCount=? where ID=?", user[0].FollowCount-1, user[0].ID)
			db.Exec("delete from FollowList where UserID=? and FollowerID=?", toUser[0].ID, user[0].ID)
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "Unfollow success"})
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

func FollowList(c *gin.Context) {
	userID := c.Query("user_id")
	dbInit()
	defer db.Close()
	var userList []User
	//从数据库获取关注列表
	var followList []dbFollower
	db.Select(&followList, "select UserID, Name, FollowCount, FollowerCount, IsFollow from FollowList where FollowerID=?", userID)
	//填充至返回的列表
	for _, user := range followList {
		userList = append(userList, User{
			Id:            user.UserID,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      user.IsFollow,
		})
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "",
		},
		UserList: userList,
	})
}

func FollowerList(c *gin.Context) {
	userID := c.Query("user_id")
	dbInit()
	defer db.Close()
	var userList []User
	//从数据库获取粉丝列表
	var followList []dbFollower
	db.Select(&followList, "select FollowerID from FollowList where UserID=?", userID)
	//填充至返回的列表
	for _, follower := range followList {
		var users []dbUser
		db.Select(&users, "select ID, Name, FollowCount, FollowerCount, IsFollow from User where ID=?", follower.FollowerID)
		userList = append(userList, User{
			Id:            users[0].ID,
			Name:          users[0].Name,
			FollowCount:   users[0].FollowCount,
			FollowerCount: users[0].FollowerCount,
			IsFollow:      users[0].IsFollow,
		})
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "",
		},
		UserList: userList,
	})
}
