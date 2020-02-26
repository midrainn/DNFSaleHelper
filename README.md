# DNFSaleHelper
### 简要说明
```
 此项目为WEB网页程序，一个DNF商人用的小账本
```
前端使用了VUE,有本地模式和云存档模式，云存档api和host请自行配置
云存档后端使用golang实现，如果遇到跨域问题，请在orgin.config中加入允许的跨域,特别注意需要留一个空行,不然本地将会被跨域拦截
数据库为mysql