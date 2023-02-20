package myGrom

import (
	"fmt"
	"testing"
)

func TestDb(t *testing.T) {
	str := ""

	Db.Table("tb_blog").Select("title").Where("id = 4").Find(&str)
	fmt.Println(str)
}
