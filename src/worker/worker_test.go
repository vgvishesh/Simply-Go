package worker

import "testing"

func TestWhatAmI(t *testing.T) {
	result := WhatAmI(true)
	if result != BOOL {
		t.Fail()
	}
	result = WhatAmI(1)
	if result != INT {
		t.Fail()
	}
	result = WhatAmI("hey")
	if result != STRING {
		t.Fail()
	}
}
