package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

type User struct {
	Id   int
	Name string
}

//======dao================
func mockGetUser(id int) (*User, error) {
	return nil, sql.ErrNoRows
}

// 遇到sql.ErrNoRows直接返回，个人不建议这样做
func daoGetUserV1(id int) (*User, error) {
	user, err := mockGetUser(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, "daoGetUserV1")
		} else {
			log.Println(err)
			return nil, err
		}
	}
	return user, nil
}

var ErrRecordNotFound = errors.New("dao: err record not found")
// 遇到sql.ErrNoRows时返回sentinel error
func daoGetUserV2(id int) (*User, error) {
	user, err := mockGetUser(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.WithStack(ErrRecordNotFound)
		} else {
			log.Println(err)
			return nil, err
		}
	}
	return user, nil
}

//========service===============
func serviceV1() {
	user, err := daoGetUserV1(1)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("没找到 %s %+v\n", err, err)
			return
		} else {
			log.Println("报错了", err)
			return
		}
	}
	fmt.Println(user)
}

func serviceV2() {
	user, err := daoGetUserV2(1)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			log.Println("没找到")
			return
		} else {
			log.Println("报错了", err)
			return
		}
	}
	fmt.Println(user)
}

//==============================

func main() {
	serviceV1()
	fmt.Println("=====================分割线=====================")
	serviceV2()
}