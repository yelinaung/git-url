package giturl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	sshUrl  = "git@github.com:yelinaung/git-url.git"
	httpUrl = "https://github.com/yelinaung/git-url.git"
)

func TestAreMenNamesRamdonOrNot(t *testing.T) {

	assert := assert.New(t)

	assert.Equal(123, 123, "they should be equal")

	ssh2http := Ssh2Http(sshUrl)
	assert.Equal(ssh2http, httpUrl, "they should be equanl")

	http2ssh := Http2Ssh(httpUrl)
	assert.Equal(http2ssh, sshUrl, "they should be equanl")

}
