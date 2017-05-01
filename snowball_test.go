package snowball

import (
    "testing"
)

var (
    username = "username"
    password = "password"
)

func TestClient_GetDetail(t *testing.T) {
    client := New(username, password)
    client.Login()
    list := client.GetDetail("AMD,RGSE")
    if len(list) != 2 {
        t.Error("failed")
    }
}

