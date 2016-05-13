package gitio

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestGenerateShortURL(t *testing.T) {

	base := "https://git.io"

	u, err := url.ParseRequestURI("https://github.com/osamingo/gitio")
	require.NoError(t, err)

	defer gock.Off()
	gock.New(base).Post("").ReplyError(errors.New("dummy error"))

	_, err = GenerateShortURL(u, "")
	require.Error(t, err)
	assert.EqualError(t, err, "Post https://git.io: dummy error")

	gock.Flush()
	gock.New(base).Post("").Reply(http.StatusPreconditionFailed).BodyString("dummy body")

	_, err = GenerateShortURL(u, "")
	require.Error(t, err)
	assert.EqualError(t, err, "invalid http status code\nstatusCode: 412\nmessage: dummy body")

	gock.Flush()
	gock.New(base).Post("").Reply(http.StatusCreated).SetHeader("location", "https://git.io/gitio-cmd")

	r, err := GenerateShortURL(u, "")
	require.NoError(t, err)
	assert.Equal(t, "https://git.io/gitio-cmd", r)
}
