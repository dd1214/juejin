package controller

import (
	"context"
	"fmt"
	idworker "github.com/gitstliu/go-id-worker"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/ssh"
	"net"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Article struct {
	Id     int64  `json:"id,omitempty"`
	Author User   `json:"author"`
	Url    string `json:"url" json:"url,omitempty"`
	//PublishTime   int64  `json:"publish_time,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Text          string `json:"text,omitempty"`
	Introduction  string `json:"introduction,omitempty"`
	Title         string `json:"title,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}
type ViaSSHDialer struct {
	client *ssh.Client
}

func (self *ViaSSHDialer) Dial(context context.Context, addr string) (net.Conn, error) {
	return self.client.Dial("tcp", addr)
}

var db *sqlx.DB

//数据库初始化，连接远程数据库
func dbInit() {
	// 一个ClientConfig指针,指向的对象需要包含ssh登录的信息
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("flash123"), //用户的密码
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //服务端验证
	}

	client, err := ssh.Dial("tcp", "192.168.1.9:22", config) //服务器的的IP地址
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	defer client.Close()

	mysql.RegisterDialContext("mysql+tcp", (&ViaSSHDialer{client}).Dial)
	dsn := "root:flash123@tcp(127.0.0.1:3306)/juejin?charset=utf8&parseTime=True"

	database, err := sqlx.Open("mysql", dsn)
	//本地数据库，测试用
	//database, err := sqlx.Open("mysql", "root:flash123@tcp(localhost:3306)/juejin") //此行需修改本地MySQL密码和数据库名称
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	db = database
}

type dbUser struct {
	ID            int64  `db:"ID"`
	Name          string `db:"Name"`
	FollowCount   int64  `db:"FollowCount"`
	FollowerCount int64  `db:"FollowerCount"`
	IsFollow      bool   `db:"IsFollow"`
	token         string `db:"token"`
}

type dbFollower struct {
	UserID        int64  `db:"UserID"`
	FollowerID    int64  `db:"FollowerID"`
	Name          string `db:"Name"`
	FollowCount   int64  `db:"FollowCount"`
	FollowerCount int64  `db:"FollowerCount"`
	IsFollow      bool   `db:"IsFollow"`
	token         string `db:"token"`
}

type dbArticle struct {
	ID            int64  `db:"ID"`
	AuthorID      int64  `db:"AuthorID"`
	Url           string `db:"Url"`
	FavoriteCount int64  `db:"FavoriteCount"`
	CommentCount  int64  `db:"CommentCount"`
	IsFavorite    bool   `db:"IsFavorite"`
	Title         string `db:"Title"`
	PublishTime   int64  `db:"PublishTime"`
	Text          string `db:"Text"`
	Introduction  string `db:"Introduction"`
}

type dbComment struct {
	ID          int64  `db:"ID"`
	UserID      int64  `db:"UserID"`
	ArticleID   int64  `db:"ArticleID"`
	CommentText string `db:"CommentText"`
	CreateDate  string `db:"CreateDate"`
}

type dbLike struct {
	UserID    int64 `db:"UserID"`
	ArticleID int64 `db:"ArticleID"`
}

//生成唯一ID
func makeId() int64 {
	currWoker := &idworker.IdWorker{}
	currWoker.InitIdWorker(1000, 1)
	newID, _ := currWoker.NextId()
	return newID
}
