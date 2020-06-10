// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package response

import (
	"encoding/json"
	"net/http"
)

// ResponseJson 将传过来的 resp Marshal 成 Json，写到 w
func ResponseJson(w *http.ResponseWriter, resp interface{}) {
	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 👇这行代码解决前端开发过程中 No 'Access-Control-Allow-Origin' header is present on the requested resource 的不便
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	// 👆在生产环境应该禁用
	(*w).Header().Set("Content-Type", "application/json")
	if _, err = (*w).Write(js); err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
	}
}
