package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func apiTestHelper(t *testing.T, addr string) map[string]interface{} {
	result, err := http.Get("http://localhost:8080/api/" + addr)
	if err != nil {
		t.Error(err)
	}

	data, err2 := ioutil.ReadAll(result.Body)
	if err2 != nil {
		t.Error(err2)
	}

	var ret map[string]interface{}

	if err := json.Unmarshal(data, &ret); err != nil {
		panic(err)
	}

	return ret
}

func TestAPI(t *testing.T) {
	go Serve()
	time.Sleep(500 * time.Millisecond)

	d1 := apiTestHelper(t, "google.com")
	if d1["up"] != true {
		t.Error("Google should be up!")
	}

	d2 := apiTestHelper(t, "https://ya.ru:80/kdoakdais")
	if d2["up"] != true {
		t.Error("Yandex should be up!")
	}

	d3 := apiTestHelper(t, "dsjadajdiajdaiodja.com")
	if d3["up"] != false {
		t.Error("dsjadajdiajdaiodja.com should be down!")
	}

	d4 := apiTestHelper(t, "8319208L:2131312!;№2;")
	if d4["errorCode"].(float64) != 0.0 {
		t.Error("8319208L:2131312!;№2; should return zero error code!")
	}
}
