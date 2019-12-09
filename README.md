---
description: Golang 爬取珍爱网数据
---

# Golang 爬虫

![&#x722C;&#x866B;&#x7ED3;&#x6784;](.gitbook/assets/image%20%282%29.png)



### 获取珍爱网所有城市第一页用户的详情

* **Fetcher 获取目标页面代码**

```text
http.Get(url) 获取页面源代码
源代码可能不是utf8的编码类型
所以引入库golang.org/x/net/html/charset，golang.org/x/text/transform
charset.DetermineEncoding确定代码的编码类型
transform.NewReader将指定编码类型转换成utf8
```

{% hint style="info" %}
若浏览器可以直接访问, 但http.Get\(url\)报403, 则需要在请求设置User-Agent  （[https://www.cnblogs.com/haoprogrammer/p/10696027.html](https://www.cnblogs.com/haoprogrammer/p/10696027.html)）
{% endhint %}



* **城市列表parser \(获取城市列表与城市详情页url\)**

```text
从上一步中能获取到整个页面的源代码
页面中城市列表的代码
<a href="http://www.zhenai.com/zhenghun/aomen" data-v-5e16505f>中国澳门</a>
<a href="http://www.zhenai.com/zhenghun/baicheng1" data-v-5e16505f>白城</a>
...
要匹配的就是:
<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" data-v-5e16505f>([^<]+)</a>
```

![&#x57CE;&#x5E02;&#x9875;&#x9762;&#x8BE6;&#x60C5;&#x7684;url](.gitbook/assets/image.png)

* **城市详情parser \(获取用户列表与用户详情页url\)**

```text
该parser和城市列表parse思路一致
列表中某个用户页面的代码
<a href="http://album.zhenai.com/u/1141852519" target="_blank">向往</a>
正则匹配：<a href="http://album.zhenai.com/u/[0-9]+".*>[^<]+</a>
```



