package api

import (
	"errors"
	"fmt"
	"im/models"
	"im/service"

	"im/vo"
	websokcet "im/ws"

	"github.com/gin-gonic/gin"
)

func Chat(c *gin.Context) {
	//u, _ := GetUser(c)
	cc := websokcet.Up(c)
	cl := websokcet.NewClient(c.Query("name"), cc.RemoteAddr().String(), cc)
	go cl.Read()
	go cl.Write()
	go cl.TimeOutClose()
	websokcet.Manager.Register <- cl
}

// 用户登录
func Login(c *gin.Context) {
	u := new(vo.UserInfo)
	c.Bind(u)
	code, data := service.Login(u.N, u.P)
	switch code {
	case 401:
		c.JSON(code, data)
	case 200:
		c.JSON(code, gin.H{
			"msg":  "登录成功",
			"data": data,
		})
	}
}

// 用户注册
func Register(c *gin.Context) {
	u := new(vo.UserR)
	c.Bind(u)
	code := service.Register(u)
	if code == 200 {
		c.JSON(code, "注册成功")
	} else {
		c.JSON(code, "注册失败")
	}
}

// 申请好友接口
func Addfriend(c *gin.Context) {
	v := c.PostForm("name")
	u, err := GetUser(c)
	if err != nil {
		c.JSON(500, err)
		return
	}
	code, data := service.Add(v, u)
	c.JSON(code, data)
}

// 同意好友申请
func AgreeFriendPost(c *gin.Context) {
	v := c.Query("name")
	u, err := GetUser(c)
	if err != nil {
		c.JSON(500, err)
		return
	}
	code, data := service.Agree(v, u.Name)
	c.JSON(code, data)
}

// 拉取好友请求列表
func GetPostList(c *gin.Context) {
	u, err := GetUser(c)
	if err != nil {
		c.JSON(500, err)
		return
	}
	code, data := service.GetPostList(u.Name)
	c.JSON(code, data)
}

// 拉取好友列表
func GetList(c *gin.Context) {
	u, err := GetUser(c)
	if err != nil {
		c.JSON(500, err)
		return
	}
	code, data := service.GetList(u.Name)
	c.JSON(code, data)
}
func GetUser(c *gin.Context) (*models.User, error) {
	y, err := c.Get("user") //得到登录用户的信息
	fmt.Println(y)
	if !err {

		e := errors.New("提取用户失败")
		return nil, e
	}
	user1 := y.(models.User) //将any格式转化为USER格式
	return &user1, nil
}

// 获取消息记录
func Get_messages_list(c *gin.Context) {
	u, err := GetUser(c)
	if err != nil {
		c.JSON(500, err)
		return
	}
	code, data := service.Get_messages_list(u.Name, c.Query("name"))
	c.JSON(code, data)
}

// 创建群聊
func CreateGroup(c *gin.Context) {
	u, err := GetUser(c)
	if err != nil {
		c.JSON(500, err)
		return
	}
	a := new(vo.GroupName)
	c.Bind(a)
	//fmt.Println(a.Gname)
	code, data := service.CreateGroup(u.Name, a.Gname)
	if data != nil {
		c.JSON(code, data)
	} else {
		c.JSON(code, "创建成功")
	}
}

// 加入群聊
func JoinGroup(c *gin.Context) {
	u, err := GetUser(c)
	if err != nil {
		c.JSON(500, err)
		return
	}
	a := new(vo.GroupName)
	c.Bind(a)

	code, data := service.JoinGroup(a.Gname, u.Name)
	if data != nil {
		c.JSON(code, data)
	} else {
		c.JSON(code, "加群成功")
	}
}

// 查看自己所加群聊列表
func SearchGrouplist(c *gin.Context) {
	u, err := GetUser(c)
	if err != nil {
		c.JSON(500, err)
		return
	}
	code, data := service.SearchGrouplist(u.Name)
	c.JSON(code, data)
}

func Get_group_messages_list(c *gin.Context) {
	code, data := service.Get_group_messages_list(c.Query("gname"))
	c.JSON(code, data)
}
