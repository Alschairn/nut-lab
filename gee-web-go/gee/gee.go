package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) // 用户定义的处理方法

// 实现http.Handler
type Engine struct {
	router map[string]HandlerFunc // 路由映射表 key: 请求方法+Path  value: 处理方法
}

// 实现接口Handler的方法
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 请求类型(GET/POST/PUT/DELETE)加路径构建路由key
	key := req.Method + "-" + req.URL.Path
	// 根据拼装的key从路由中获取定义的处理方法
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "Bro, I can't find this page, 404 NOT FOUND: %s\n", req.URL)
	}
}


// 获取一个处理引擎实例
func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

// 添加路由信息
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// 添加GET路由信息
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// 添加POST路由信息
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// 启动http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
