package dao

import (
	"im/models"
	"im/server"
)

// ------------------用户操作---------------------------

// 创建用户
func CreateUser(n string, pwd string, e string, p string) {
	db := server.GetMysqlDB()
	a := new(models.User)
	a.Name = n
	a.Pwd = pwd
	a.Email = e
	a.Phone = p
	db.Create(a)
}

// 创建好友申请记录
func Create_friend_application(n string, u *models.User) {
	v := &models.Friend_application_list{
		FromId: u.Name,
		Name:   n,
	}
	db := server.GetMysqlDB()
	db.Create(v)
}

// 创建好友
func Create_friend(friend string, myself string) {
	v := &models.Hail_fellow{
		FellowName: friend,
		Name:       myself,
	}
	db := server.GetMysqlDB()
	db.Create(v)
}

// 删除用户
func DeleteUser(n string) {
	db := server.GetMysqlDB()
	db.Where("name = ?", n).Delete(&models.User{})
}
func Delete_friend_application(n string, s string) {
	db := server.GetMysqlDB()
	db.Where("name = ? and from_id = ?", n, s).Delete(&models.Friend_application_list{})
}

// 查找用户
func SearchUser(n string) *models.User {
	db := server.GetMysqlDB()
	a := new(models.User)
	db.Where("name = ?", n).Find(a)
	return a
}

// 获取好友添加请求
func SearchPostList(n string) []models.Friend_application_list {
	db := server.GetMysqlDB()
	a := make([]models.Friend_application_list, 0)
	db.Where("Name = ?", n).Find(&a)
	return a
}

// 获取好友列表
func SearchFriendList(n string) []models.Hail_fellow {
	db := server.GetMysqlDB()
	a := make([]models.Hail_fellow, 0)
	db.Where("Name = ?", n).Find(&a)
	return a
}

// 更新用户信息
func UpdateUser(n string, a *models.User) {
	db := server.GetMysqlDB()
	b := SearchUser(n)
	b.Name = a.Name
	b.Pwd = a.Pwd
	b.Email = a.Email
	b.Phone = a.Phone
	db.Save(b)
}
