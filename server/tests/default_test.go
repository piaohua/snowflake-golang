package tests

import (
	"io/ioutil"
	"testing"

	beetest "github.com/astaxie/beego/testing"
)

func TestUrl(t *testing.T) {
	request := beetest.Post("/")
	request.Param("count", "10")
	response, _ := request.Response()
	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	t.Logf("contents %v", contents)
}

func TestPing(t *testing.T) {
	request := beetest.Get("/v1/id")
	response, _ := request.Response()
	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	t.Logf("contents %v", contents)
}
