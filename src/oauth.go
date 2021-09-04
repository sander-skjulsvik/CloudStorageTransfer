package src

import (
	"fmt"
	"time"
)

type Oauth struct {
	Timeout time.Time
	token   string
}

func NewOauthPwd(id string, pwd string) {

}

func (a *Oauth) refreshToken() {
	fmt.Println("Token Graph token expired enter new one: ")
	fmt.Scanln(a.token)
	for !a.testToken() {
		fmt.Printf("Token: %s does not work. \nEnter another one:\n", a.token)
		fmt.Scanln(a.token)
	}
}

func (a *Oauth) testToken() bool {
}

func (a *Oauth) GetToken() string {
	if a.Timeout.After(time.Now()) {
		fmt.Println("Token expired geting new one.")
		a.refreshToken()
	}
	return a.token
}
