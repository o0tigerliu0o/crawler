package fetcher

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	content, err := Fetch("http://album.zhenai.com/u/1814582139")
	// 内容判断
	if nil != err {
		t.Errorf("expected : nil;\n but was %v \n content=[%v]", err, content)
	}
	fmt.Println(string(content))
}
