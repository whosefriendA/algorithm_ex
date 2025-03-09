package main

import (
	"fmt"
	"io"
	"net/http"
)

func main(){
	res,err :=http.Get("https://blog.csdn.net/weixin_42998312/article/details/144015015?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522adc3dabbee77fc61b23e04d16c8439af%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=adc3dabbee77fc61b23e04d16c8439af&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_ecpm_v1~hot_rank-1-144015015-null-null.142^v102^control&utm_term=go%20net%2Fhttp&spm=1018.2226.3001.4187")
	if err !=nil{
		fmt.Println("Error:",err)
		return 
	}

	defer res.Body.Close()

	body,err :=io.ReadAll(res.Body)

	if err !=nil{
		fmt.Println("Error reading body:",err)
	}
	fmt.Println("res:",string(body))
}