package mysql

import (
	"testing"
	"web_app/models"
	"web_app/settings"
)

func init() {
	dbcfg := &settings.MysqlConfig{
		Host:         "127.0.0.1",
		Port:         3306,
		User:         "root",
		Password:     "8871527yhk",
		DbName:       "bulebell",
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	err := Init(dbcfg)
	if err != nil {

	}
}
func TestCreatePost(t *testing.T) {
	post := &models.Post{
		ID:          1,
		AuthorID:    123,
		CommunityID: 1,
		Title:       "test",
		Content:     "just a test",
	}
	err := CreatePost(post)
	if err != nil {
		t.Fatalf("CreatePost insert record into mysql failed,err:%v\n", err)
	}
	t.Logf("CreatePost insert record into mysql success")
}
