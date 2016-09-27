# Troubleshoot

#### UnicodeEncodeError: 'cp950' codec can't encode character
> 中文的 windows「命令提示字元」(cmd) 編碼預設是：cp950 ，而Python3 的預設程式碼編碼是：utf-8 (cp65001)，另外通常網頁的編碼也是 utf-8 (cp65001)，這時候就很有可能在執行 print() 指令時出現 「UnicodeEncodeError: 'cp950' codec can't encode character ... ...」的錯誤，進而導致程式腳本執行中斷

兩個解決方法：
1. 改「命令提示字元」(cmd) 的編碼，執行 `$chcp 65001`，硬改「命令提示字元」(cmd) 的編碼成 unicode (UTF-8)，這樣它就不會在輸出文字時需要轉換了，不過這對要封裝程式給別人用的人是一個很大的麻煩，因為並不是每個使用者都有足夠的能力去改這東西。

2. 使用 encode 的 replace 功能，把文字改成 "cp950" 的編碼，同時把 "cp950" 不認識的字替換掉 (變成 '?' )，這之後再把它用 "cp950" decode 回去就可以了。


* [UnicodeEncodeError: 'cp950' codec can't encode character...](http://marsray.pixnet.net/blog/post/61040521-%5Bpython3%5D-%E7%94%A8-python3-%E5%AF%AB%E4%B8%80%E5%80%8B%E7%B6%B2%E8%B7%AF%E7%88%AC%E8%9F%B2)


#### configparser '%' in value will cause ValueError
> [section]<br/>
> name = 'value%'<br/>
> ValueError: invalid interpolation syntax in 'value%'<br/>
> workaround:<br/>
> name = 'value%'.replace('%', '%%')<br/>

* 造成原因，看起來是'%'會被當作替換字元，如果要把'%'當作*值*寫出，可以用兩個'%%'
> [Paths]<br/>
> home_dir: /Users<br/>
> my_dir: %(home_dir)s/lumberjack<br/>
> my_pictures: %(my_dir)s/Pictures
* [Interpolation of values](https://docs.python.org/3/library/configparser.html#interpolation-of-values)
