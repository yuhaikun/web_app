package mysql

import (
	"strings"
	"web_app/models"

	"github.com/jmoiron/sqlx"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
	post_id,title,content,author_id,community_id)
	values (?,?,?,?,?)
	`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

// GetPostById 根据id查询单个帖子的数据
func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id,title,content,author_id,community_id, create_time from post where post_id = ?`
	err = db.Get(post, sqlStr, pid)
	return
}

// GetPostList 查询帖子列表函数
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select 
    post_id,title,content,author_id,community_id, create_time 
	from post
    ORDER BY create_time
	DESC
	limit ?,?
    `
	//if size == 0 {
	//	size = 2
	//}
	//posts = make([]*models.Post, 0, size) // 不要写成make([]*models.Post,2)
	//if page == 0 {
	//	err = db.Select(&posts, sqlStr, 0, size)
	//	return
	//}
	posts = make([]*models.Post, 0, size)
	if size == 0 && page == 0 {
		sqlStr = `select 
			post_id,title,content,author_id,community_id, create_time 
			from post
			ORDER BY create_time
			DESC
			`
		err = db.Select(&posts, sqlStr)
		return
	} else if page == 0 {
		err = db.Select(&posts, sqlStr, 0, size)
		return
	}
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

// GetPostListByIDs 根据给定的id列表查询帖子数据
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where post_id in (?)
	ORDER BY FIND_IN_SET(post_id,?)
	`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	/*
		在给定的代码中，db.Rebind函数没有使用的原因是因为SQL查询语句中没有包含任何需要绑定的参数，因此不需要调用db.Rebind函数。
		db.Rebind函数的作用是将SQL查询语句中的占位符重新绑定到具体的数据库驱动程序的占位符上。
		这在涉及到动态生成的SQL查询语句或者使用不同的数据库驱动程序时非常有用。在该代码中，查询语句中的占位符是?，而参数列表中的参数已经通过db.Select函数的第二个参数传递，不需要重新绑定。
	*/
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
