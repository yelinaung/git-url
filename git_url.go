package giturl

import (
	"fmt"
	str "strings"
)

func Http2Ssh(httpUrl string) string {

	userNameRepoPair := str.Split(httpUrl, "https://")[1]
	tempPair := str.Split(userNameRepoPair, "/")
	userName := tempPair[1]
	repoName := str.Split(tempPair[2], ".git")[0]

	return fmt.Sprintf("git@github.com:%s/%s.git", userName, repoName)
}

func Ssh2Http(sshUrl string) string {
	// fmt.Printf("%s is http ? %t\n", sshUrl, str.Contains(sshUrl, http))

	userNameRepoPair := str.Split(sshUrl, ":")[1]
	tempPair := str.Split(userNameRepoPair, "/")
	userName := tempPair[0]
	repoName := str.Split(tempPair[1], ".git")[0]

	return fmt.Sprintf("https://github.com/%s/%s.git", userName, repoName)
}
