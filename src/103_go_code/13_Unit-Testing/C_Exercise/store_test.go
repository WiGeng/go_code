package store

import "testing"

func TestStore(t *testing.T) {
	monster := &Monster{
		Name:  "sunwukong",
		Age:   26,
		Skill: "DevOps",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store() 错误，希望为=%v 实际为=%v", true, res)
	}
	t.Logf("monster.Store() 测试成功!")
}

func TestReStore(t *testing.T) {

	//测试数据是很多，测试很多次，才确定函数，模块..

	//先创建一个 Monster 实例 ， 不需要指定字段的值
	var monster = &Monster{}
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore() 错误，希望为=%v 实际为=%v", true, res)
	}

	//进一步判断
	if monster.Name != "sunwukong" {
		t.Fatalf("monster.ReStore() 错误，希望为=%v 实际为=%v", "sunwukong", monster.Name)
	}

	t.Logf("monster.ReStore() 测试成功!")
}
