package dbops

import (
	"log"
	"database/sql"
	"HqVideo/api/defs"
	"HqVideo/api/utils"
	"time"
)
//用户
func AddUserCredential(loginName string,pwd string) error {


	stmtIns ,err := dbConn.Prepare("insert into users (login_name,pwd) values(?,?)")
	if err != nil{
		return err
	}
	_,err = stmtIns.Exec(loginName,pwd)
	if err != nil {
		return nil
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string,error)  {
	stmtOut ,err := dbConn.Prepare("select pwd from users where login_name = ?")
	if err != nil {
		log.Printf("%s",err)
		 return "",err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows{

		return "",err

	}
	defer stmtOut.Close()

	return pwd,nil
}

func GetUser(loginName string) (*defs.User,error)  {
	stmtOut ,err := dbConn.Prepare("select id  from users where login_name = ?")
	if err != nil {
		log.Printf("%s",err)
		return nil,err
	}
	var uid int
	err = stmtOut.QueryRow(loginName).Scan(&uid)
	if err != nil && err != sql.ErrNoRows{

		return nil,err

	}
	user := &defs.User{Id:uid}
	defer stmtOut.Close()

	return user,nil
}

func DeleteUserCredential(loginName string,pwd string) error  {
	stmtDel ,err := dbConn.Prepare("delete  from users where login_name = ? and pwd =?")
	if err != nil {
		log.Printf("%s",err)
		return err
	}
	_,err = stmtDel.Exec(loginName,pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}
//视频
func AddNewVideo(aid int,name string) (*defs.VideoInfo,error) {

	vid,err := utils.NewUUID()
	if err != nil {
		return nil,err
	}

	t := time.Now()
	// Jan 02 2006, 15:04:05这是一个模块不能更改
	ctime := t.Format("Jan 02 2006, 15:04:05") //M D Y HH:mm:ss

	stmtIns,err := dbConn.Prepare("insert into video_info (id,author_id,name,display_time) values (?,?,?,?)")
	if err != nil {
		return nil,err
	}
	_,err = stmtIns.Exec(vid,aid,name,ctime)
	if err != nil {
		return nil,err
	}

	res := &defs.VideoInfo{Id:vid,AuthorId:aid,Name:name,DisplayTime:ctime}

	defer stmtIns.Close()

	return res,nil

}
func GetVideoInfo(vid string) (*defs.VideoInfo,error)  {
	stmtOut,err := dbConn.Prepare("select author_id,name,display_time from video_info where id = ? ")
	var aid int
	var dct string
	var name string
	err = stmtOut.QueryRow(vid).Scan(&aid,&dct,&name)
	if err == sql.ErrNoRows {
		return  nil,nil
	}
	defer stmtOut.Close()

	res := &defs.VideoInfo{Id:vid,AuthorId:aid,Name:name,DisplayTime:dct}
	return res,nil
}
func ListVideoInfo(uid string,from,to int) ([]*defs.VideoInfo ,error)  {

	stmtOut,err := dbConn.Prepare(`
	select video_info.id,video_info.author_id,video_info.name,video_info.display_time,users.login_name from video_info 
	inner join users on video_info.author_id= users.id
	where video_info.video_id = ?
	 and video_info.create_time > from_unixtime(?) 
	 and video_info.create_time <= from_unixtime(?)
`)
	var res []*defs.VideoInfo
	rows,err := stmtOut.Query(uid,from,to)
	if err != nil {
		return res,err
	}

	for rows.Next() {
		var id string
		var authorId int
		var name string
		var displayTime string

		if err := rows.Scan(&id,&authorId,&name,&displayTime); err != nil {
			return res,err
		}
		cm := &defs.VideoInfo{Id:id,AuthorId:authorId,Name:name,DisplayTime:displayTime}
		res = append(res,cm)
	}
	defer stmtOut.Close()

	return res,nil
}
func DeleteVideoInfo(vid string) error  {

	stmtDel,err := dbConn.Prepare("delete from video_info where id = ?")
	if err != nil {
		return err
	}
	_,err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

//评论

func AddNewComments(vid string,aid int,content string)  error{

	id,err := utils.NewUUID()
	if err != nil {
		return err
	}


	stmtIns,err := dbConn.Prepare("insert into comments (id,video_id,author_id,content) values (?,?,?,?)")
	if err != nil {
		return err
	}
	_,err = stmtIns.Exec(id,vid,aid,content)
	if err != nil {
		return err
	}

	return nil
}

func ListComments(vid string,from,to int) ([]*defs.Comment ,error)  {

	stmtOut,err := dbConn.Prepare(`
	select comments.id,users.login_name,comments.content from comments
 	inner join users on comments.author_id = users.id 
	where comments.video_id = ?
	 and comments.time > from_unixtime(?) 
	 and comments.time <= from_unixtime(?)
	order by comments.time desc 
`)
	var res []*defs.Comment
	rows,err := stmtOut.Query(vid,from,to)
	if err != nil {
		return res,err
	}

	for rows.Next() {
		var id,name,content string
		if err := rows.Scan(&id,&name,&content); err != nil {
			return res,err
		}
		cm := &defs.Comment{Id:id,VideoId:vid,Author:name,Content:content}
		res = append(res,cm)
	}
	defer stmtOut.Close()

	return res,nil
}

