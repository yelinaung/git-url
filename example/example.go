package main

import (
	"fmt"
	"github.com/yelinaung/git-url"
)

func main() {
	sshUrl := "git@github.com:yelinaung/git-url.git"
	fmt.Println(giturl.Ssh2Http(sshUrl))

	httpUrl := "https://github.com/yelinaung/git-url.git"
	fmt.Println(giturl.Http2Ssh(httpUrl))
}
