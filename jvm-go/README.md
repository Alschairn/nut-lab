## 问题记录

---
> 以下问题仅是本人验证方法，不代表正确go mod使用方式


- 引用报错
````
引用子包也必须是 github.com/xxxx 类型，不可以是 github.com/xxxx/xxxx 这种描述，go mod解析会不识别
```` 

- 同项目包互相引用
````
例如我们有A、B两个包，B包需要引用A包，需要在go mod中指定

require github.com/A version

replace github.com/A => ../A
````

- 多级引用
````
- A
	- A1
	- A2
- B
	依赖A1
- C
	依赖B（go Mod依赖B，并且也必须引用A1）
````