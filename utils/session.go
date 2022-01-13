package utils

import (
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
)

type ORMOperation func(session *xorm.Session) error

//Transaction  ...事务提交
func Transaction(session *xorm.Session, f ORMOperation) (err error) {
	err = session.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			glog.Errorf("recover rollback:%s\r\n", p)
			session.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			glog.Errorf("error rollback:%s\r\n", err)
			session.Rollback() // err is non-nil; don't change it
		} else {
			err = session.Commit() // err is nil; if Commit returns error update err
		}
	}()
	err = f(session)
	return err
}
