// Copyright 2017 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vhost

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/9sr/frp/utils/version"
)

const (
	NotFound = `<!DOCTYPE html>
<html>
<head>
<title>访问错误</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>访问的页面出错。</h1>
<p>抱歉, 您访问的页面不存在或暂时不可用，可以稍后重试。</p>
<p>服务由 "杰思网络®" 提供.</p>
</body>
</html>
`
)

func notFoundResponse() *http.Response {
	header := make(http.Header)
	header.Set("server", "frp/"+version.Full())
	header.Set("Content-Type", "text/html")
	res := &http.Response{
		Status:     "Not Found",
		StatusCode: 404,
		Proto:      "HTTP/1.0",
		ProtoMajor: 1,
		ProtoMinor: 0,
		Header:     header,
		Body:       ioutil.NopCloser(strings.NewReader(NotFound)),
	}
	return res
}
