package main

import (
	"encoding/json"
	"testing"
)

func checkTestHelper(t *testing.T, addr string) map[string]interface{} {
	result := Check(addr)

	var ret map[string]interface{}

	if err := json.Unmarshal([]byte(result), &ret); err != nil {
		panic(err)
	}

	return ret
}

func TestCheck(t *testing.T) {
	d1 := checkTestHelper(t, "google.com")
	if d1["up"] != true {
		t.Error("Google should be up!")
	}

	d2 := checkTestHelper(t, "https://ya.ru:80/kdoakdais")
	if d2["up"] != true {
		t.Error("Yandex should be up!")
	}

	d3 := checkTestHelper(t, "dsjadajdiajdaiodja.com")
	if d3["up"] != false {
		t.Error("dsjadajdiajdaiodja.com should be down!")
	}

	d4 := checkTestHelper(t, "8319208L:2131312!;№2;")
	if d4["errorCode"].(float64) != 0.0 {
		t.Error("8319208L:2131312!;№2; should return zero error code!")
	}
}
