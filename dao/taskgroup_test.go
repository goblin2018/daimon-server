package dao

import (
	"daimon/pkg/mysql"
	"testing"
)

func TestABC(t *testing.T) {

	a := 2 / 65536
	t.Log(a + 2)
}

func TestGetSerial(t *testing.T) {
	mysql.Init()
	d := NewTaskGroupDao()
	serial := d.GetTaskGroupSerialById(3)
	t.Logf("serial %d", serial)
}
