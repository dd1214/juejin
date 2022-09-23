package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	//获取输入参数username, password
	username := c.Query("username")
	password := c.Query("password")

	//生成唯一ID
	newID := makeId()
	//用户鉴权token, 为用户名+密码的长字符串
	token := username + password

	dbInit()
	defer db.Close()
	var user []dbUser
	//查询,保证用户名唯一
	db.Select(&user, "select ID from User where Name=?", username)
	//若查询到则直接返回
	if user != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		_, err := db.Exec("insert into User(FollowCount, FollowerCount, ID, IsFollow, Name, token)value(?, ?, ?, ?, ?, ?)", 0, 0, newID, 0, username, token)
		if err != nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "User register fail"},
			})
		} else {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0, StatusMsg: "User register success"},
				UserId:   newID,
				Token:    username + password,
			})
		}
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	dbInit()
	defer db.Close()
	var user []dbUser
	//查询
	db.Select(&user, "select ID from User where token=?", token)

	//返回响应
	if user != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: ""},
			UserId:   user[0].ID,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist or password error"},
		})
	}
}

func UserInfo(c *gin.Context) {
	ID := c.Query("id")

	dbInit()
	defer db.Close()
	var user []dbUser
	//查询
	db.Select(&user, "select ID, Name, FollowCount, FollowerCount from User where ID=?", ID)

	//返回响应
	if user != nil {
		var ResponseUser = User{
			Id:            user[0].ID,
			Name:          user[0].Name,
			FollowCount:   user[0].FollowCount,
			FollowerCount: user[0].FollowerCount,
		}
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0, StatusMsg: ""},
			User:     ResponseUser,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
