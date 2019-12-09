package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte,error){

	client:=&http.Client{}
	req,err:=http.NewRequest("GET",url,nil)
	if err!=nil{
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")

	resp,err:=client.Do(req)
	if err!=nil{
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,fmt.Errorf("wrong status code: %d\n",resp.StatusCode)
	}
	bodyReader :=bufio.NewReader(resp.Body)
	e:=determineEncoding(bodyReader)
	utf8Reader:=transform.NewReader(bodyReader,
		e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}
func determineEncoding(r *bufio.Reader) encoding.Encoding{
	bytes,err:=r.Peek(1024)
	if err!=nil{
		log.Printf("Fetcher error: %v",err)
		return unicode.UTF8
	}
	e,_,_:=charset.DetermineEncoding(bytes,"")
	return e
}
