package lecloud

import (
	"testing"
	"fmt"
)

func TestGetVideoInfoParam(t *testing.T) {
	var p = GetVideoInfoParam{}
	p.VideoId = "27647891"
	fmt.Println(RequestWithKey("e05b65109e8ecbeec5a61e276ac51592", "xxx", p))
}


func TestGetVideoListParam(t *testing.T) {
	var p = GetVideoListParam{}
	fmt.Println(RequestWithKey("e05b65109e8ecbeec5a61e276ac51592", "xxx", p))
}