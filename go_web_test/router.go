package main

import (
	"net/http"
	"strings"
)

type router struct {
	// Key : Http Method
	// Value: URL HandlerFunc
	handlers map[string]map[string]http.HandlerFunc
}

func (r *router) HandleFunc(method, pattern string, h http.HandlerFunc) {
	// http 메서드로 등록된 맵이 있는지 확인
	if m, ok := r.handlers[method]; !ok {
		m = make(map[string]http.HandlerFunc)
		r.handlers[method] = m
	} else {
		m[pattern] = h
	}
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if m, ok := r.handlers[req.Method]; ok {
		if h, ok := m[req.URL.Path]; ok {
			h(w, req)
			return
		}
	}

	http.NotFound(w, req)
}

func match(pattern, path string) (bool, map[string]string) {
	if pattern == path {
		return true, nil
	}

	patterns := strings.Split(pattern, "/")
	paths := strings.Split(path, "/")

	if len(patterns) != len(paths) {
		return false, nil
	}

	params := make(map[string]string)

	for i := 0; i < len(patterns); i++ {
		switch {
		case patterns[i] == paths[i]:
			// TODO
		case len(patterns[i]) > 0 &&
		}
	}
}