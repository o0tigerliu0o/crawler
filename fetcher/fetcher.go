package fetcher

import (
	"bufio"
	"crawler/distributed/config"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var rateLimiter = time.Tick(time.Second / config.Qps)

func Fetch(url string) ([]byte, error) {

	<-rateLimiter
	// 将网页抓下来
	log.Printf("Fetching url %v", url)

	resp, err := http.Get(url)

	if nil != err {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	if http.StatusOK != resp.StatusCode {
		return nil, fmt.Errorf("Error : wrong status code: %d", resp.StatusCode)
	}

	// 根据页面编码对设置字符集
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)

	// 对网页进行字符集转换
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

// 根据页面的前1024个字符推测页面字符集
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	// 为了读取的1024个字节，依然保留在reader中供其他地方读取
	bytes, err := r.Peek(1024)
	if nil != err {
		log.Printf("Fetch error: %v", err)
		return unicode.UTF8
	}

	// 根据页面的前1024个字符推测页面字符集
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e

}

// 个人网页信息获取
func FetchProfile(url string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	now := time.Now().Second()
	//http访问头部信息设置
	//request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	switch int(now / 3) {
	case 0:
		request.Header.Add("user-agent",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 10_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1s")
	case 1:
		request.Header.Add("user-agent",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 11_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1s")
	case 2:
		request.Header.Add("user-agent",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 12_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1s")
	default:
		request.Header.Add("user-agent",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1s")
	}

	/*request.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	request.Header.Add("accept-encoding", "gzip, deflate, br")
	request.Header.Add("authority", "album.zhenai.com")*/

	client := &http.Client{}
	resp, err := client.Do(request)

	// 简单的网页内容获得
	//resp, err := http.Get("http://study.163.com")
	if nil != err {
		panic(err)
	}
	defer resp.Body.Close()

	if nil != err {
		panic(err)
	}
	defer resp.Body.Close()

	/*// 获取网页信息
	body, err := httputil.DumpResponse(c, true)
	if nil != err {
		panic(err)
	}*/

	// 根据页面编码对设置字符集
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)

	// 对网页进行字符集转换
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	//fmt.Printf("%s\n", body)
	return ioutil.ReadAll(utf8Reader)
}
