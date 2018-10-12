package dbops

import (
	"testing"
	"strconv"
	"time"
	"fmt"
)
var tempVid string
func clearTables()  {

	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")

}
//这里相当于普通go文件的init初始化函数
func TestMain(m *testing.M)  {

	clearTables()
	m.Run()
	clearTables()
}
//用户测试
func TestUserWorkFlow(t *testing.T)  {

	t.Run("Add",testAddUser)
	t.Run("Get",testGetUser)
	t.Run("Del",testDelUser)
	t.Run("Reget",testRegetUser)

}

func testAddUser(t *testing.T)  {
	err := AddUserCredential("hehuiqi","123456")
	if err != nil {
		t.Errorf("添加用户错误  %v",err)
	}

}
func testGetUser(t *testing.T)  {
	pwd, err := GetUserCredential("hehuiqi")
	if pwd != "123456" || err != nil {
		t.Errorf("获取用户错误  %v",err)
	}
}

func testDelUser(t *testing.T)  {

	err := DeleteUserCredential("hehuiqi","123456")
	if err != nil {
		t.Errorf("删除用户错误  %v",err)
	}

}

func testRegetUser(t *testing.T)  {
	pwd, err := GetUserCredential("hehuiqi")
	if  err != nil {
		t.Errorf("测试清除用户错误  %v",err)
	}
	if pwd != "" {
		t.Errorf("测试清除用户错误")

	}
}

//视频测试
func TestVideoWorkFlow(t *testing.T)  {
	clearTables()
	t.Run("PrepareUse",testAddUser)
	t.Run("AddVideo",testAddVideoInfo)
	t.Run("GetVideo",testGetVideoInfo)
	t.Run("DeleteVideo",testDeleteVideoInfo)
	t.Run("RegetVideo",testRegetVideo)
}
func testAddVideoInfo(t *testing.T)  {
	vi,err := AddNewVideo(1,"hq-video")
	if err != nil {
		t.Errorf("添加视频错误--%v",err)
	}
	tempVid = vi.Id

}

func testGetVideoInfo(t *testing.T) {

	_,err := GetVideoInfo(tempVid)
	if err != nil {
		t.Errorf("获取视频信息错误--%v",err)
	}

}
func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempVid)
	if err != nil {
		t.Errorf("删除视频信息错误--%v",err)

	}
}

func testRegetVideo(t *testing.T)  {
	vi,err := GetVideoInfo(tempVid)
	if err != nil || vi != nil {
		t.Errorf("检测视频信息错误--%v",err)
	}
}
//视频评论测试
func TestCommentsWorkFlow(t *testing.T) {
	clearTables()
	t.Run("AddUser",testAddUser)
	t.Run("AddComments",testAddNewComments)
	t.Run("ListComments",testListComments)

}

func testAddNewComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "这个视频不错"
	err := AddNewComments(vid,aid,content)
	if err != nil {
		t.Errorf("添加评论错误--%v",err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to,_ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	res,err := ListComments(vid,from,to)
	if err != nil {
		t.Errorf("查询评论列表错误--%v",err)
	}
	for i,ele := range res {
		fmt.Printf("评论：%d,%v\n",i,ele.Content)
	}
}

//session测试
func TestSessionWorkFlow(t *testing.T)  {
	clearTables()
	t.Run("InsertSession",testInsertSession)
	t.Run("SimpleSession",testRetrieveSession)
	t.Run("AllSession",testRetrieveAllSessions)
	t.Run("DelSession",testDeleteSession)

}
func testInsertSession(t *testing.T) {

	sid := "wwwwwwww"
	ttl := int64(3333333)
	uname := "qqqqqq"

	err := InsertSession(sid,ttl,uname)

	if err != nil {
		t.Errorf("创建session错误--%v",err)
	}
}
func testRetrieveSession(t *testing.T) {
	sid := "wwwwwwww"
	ss,err := RetrieveSession(sid)
	if err != nil {
		t.Errorf("查询session错误--%v",err)
	}
	fmt.Println("查询ss:",ss.Username)
}
func testRetrieveAllSessions(t *testing.T) {

	cm,err := RetrieveAllSessions()
	if err != nil {
		t.Errorf("查询所有session错误--%v",err)
	}
	fmt.Println("查询所有",cm)

}
func testDeleteSession(t *testing.T) {
	sid := "wwwwwwww"

	err := DeleteSession(sid)
	if err != nil {
		t.Errorf("删除session错误--%v",err)
	}
}