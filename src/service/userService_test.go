package service

import "testing"

func TestUser(t *testing.T) {
	var name = "'12345';INSERT INTO tb_user(`name`) VALUES('sql注入')"
	Login(name, "")
}
