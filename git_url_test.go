package giturl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	sshUrl  = "git@github.com:yelinaung/git-url.git"
	httpUrl = "https://github.com/yelinaung/git-url.git"

	bitbucketSshUrl  = "git@bitbucket.org:myname/myrepo.git"
	bitbucketHttpUrl = "https://bitbucket.org/myname/myrepo.git"
)

func TestAreMenNamesRamdonOrNot(t *testing.T) {
	assert := assert.New(t)

	ssh2http := Ssh2Http(sshUrl)
	assert.Equal(ssh2http, httpUrl, "they should be equanl")

	http2ssh := Http2Ssh(httpUrl)
	assert.Equal(http2ssh, sshUrl, "they should be equanl")

	bitbucketSsh2Http := Ssh2Http(bitbucketSshUrl)
	assert.Equal(bitbucketSsh2Http, bitbucketHttpUrl, "they should be equanl")

	bitbucketHttp2Ssh := Http2Ssh(bitbucketHttpUrl)
	assert.Equal(bitbucketHttp2Ssh, bitbucketSshUrl, "they should be equanl")
}
