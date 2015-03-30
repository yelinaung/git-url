package giturl

import (
	"fmt"
	str "strings"
)

var dotGit = ".git"

func Http2Ssh(httpUrl string) string {
	userNameRepoPair := str.Split(httpUrl, "https://")[1]

	tempPair := str.Split(userNameRepoPair, "/")
	host := tempPair[0]
	userName := tempPair[1]
	repoName := str.Split(tempPair[2], ".git")[0]

	return fmt.Sprintf("git@%s:%s/%s.git", host, userName, repoName)
}

func Ssh2Http(sshUrl string) string {
	// fmt.Printf("%s is http ? %t\n", sshUrl, str.Contains(sshUrl, http))

	hostRepoPair := str.Split(sshUrl, "@")[1]
	host := str.Split(hostRepoPair, ":")[0]
	userNameRepoPair := str.Split(hostRepoPair, ":")[1]
	userName := str.Split(userNameRepoPair, "/")[0]
	repoName := str.Split(userNameRepoPair, "/")[1]

	if str.Contains(sshUrl, dotGit) {
		return fmt.Sprintf("https://%s/%s/%s", host, userName, repoName)
	} else {
		return fmt.Sprintf("https://%s/%s/%s.git", host, userName, repoName)
	}
}

//func main() {
//	sshUrl := "git@github.com:yelinaung/git-url.git"
//	fmt.Println(Ssh2Http(sshUrl))
//
//	httpUrl := "https://bitbucket.org/myname/myrepo.git"
//	fmt.Println(Http2Ssh(httpUrl))
//}
