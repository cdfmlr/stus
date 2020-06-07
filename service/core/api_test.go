// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package core

import (
	"net/http"
	"testing"
)

func TestCoreApi(t *testing.T) {
	s := NewCoreApi(db)
	_ = http.ListenAndServe(":9001", s)
}
