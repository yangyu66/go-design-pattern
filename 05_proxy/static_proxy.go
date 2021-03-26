/*
*What is this pattern about?
Proxy is used in places where you want to add functionality to a class without
changing its interface. The main class is called `Real Subject`. A client should
use the proxy or the real subject without any code change, so both must have the
same interface. Logging and controlling access to the real subject are some of
the proxy pattern usages.
*References:
https://refactoring.guru/design-patterns/proxy/python/example
https://python-3-patterns-idioms-test.readthedocs.io/en/latest/Fronting.html
*TL;DR
Add functionality or logic (e.g. logging, caching, authorization) to a resource
without changing its interface.

	proxy := NewUserProxy(&User{})

	err := proxy.Login("test", "password")
*/

package proxy

import (
	"log"
	"time"
)

// IUser IUser
type IUser interface {
	Login(username, password string) error
}

// User 用户
// @proxy IUser
type User struct {
}

// Login 用户登录
func (u *User) Login(username, password string) error {
	// 不实现细节
	return nil
}

// proxy should keep raw GetName
func (u *User) GetName() string {
	return "ca"
}

// UserProxy 代理类
type UserProxy struct {
	*User
}

// NewUserProxy NewUserProxy
func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{
		User: user,
	}
}

// Login 登录，和 user 实现相同的接口
func (p *UserProxy) Login(username, password string) error {
	// before 这里可能会有一些统计的逻辑
	start := time.Now()

	// 这里是原有的业务逻辑
	if err := p.User.Login(username, password); err != nil {
		return err
	}

	// after 这里可能也有一些监控统计的逻辑
	log.Printf("user login cost time: %s", time.Now().Sub(start))

	return nil
}
