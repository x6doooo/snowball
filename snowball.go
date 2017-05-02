package snowball

import (
    "net/http/cookiejar"
    "net/http"
    "crypto/md5"
    "io"
    "encoding/hex"
    "strings"
    "net/url"
    "strconv"
    "io/ioutil"
    "encoding/json"
    "time"
)

const (
    host = "https://xueqiu.com"
    apiCsrf = host + "/service/csrf?api=/user/login"
    apiLogin = host + "/user/login"
    apiStockList = host + "/stock/cata/stocklist.json"
    apiStockDetail = host + "/v4/stock/quote.json"
    apiEvents = host + "/calendar/cal/events.json"
)

type Client struct {
    jar        *cookiejar.Jar
    httpClient *http.Client
    username   string
    password   string
}

func md5hex(str string) string {
    h := md5.New()
    io.WriteString(h, str)
    resBytes := h.Sum(nil)
    resStr := hex.EncodeToString(resBytes)
    return strings.ToUpper(resStr)
}

func (me *Client) Login() error {
    // csrf request
    csrfReq, err := http.NewRequest("GET", apiCsrf, nil)
    if err != nil {
        return err
    }
    resp, err := me.httpClient.Do(csrfReq)
    if err != nil {
        return err
    }
    resp.Body.Close()

    // login request
    postData := url.Values{}
    postData.Set("telephone", me.username)
    postData.Set("remember_me", "on")
    postData.Set("areacode", "86")
    postData.Set("password", me.password)
    postDataReader := strings.NewReader(postData.Encode())
    loginReq, err := http.NewRequest("POST", apiLogin, postDataReader)
    loginReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    if err != nil {
        return err
    }

    resp, err = me.httpClient.Do(loginReq)
    if err != nil {
        return err
    }
    resp.Body.Close()
    return nil
}

func (me *Client) GetCodeList() (list []string) {
    pageNum := 0
    count := -1
    size := 100
    params := url.Values{}
    params.Set("size", strconv.Itoa(size))
    params.Set("order", "asc")
    params.Set("orderby", "code")
    params.Set("type", "0,1,2")
    for {
        codeList := CodeList{}
        pageNum += 1
        params.Set("page", strconv.Itoa(pageNum))
        req, _ := http.NewRequest("GET", apiStockList + "?" + params.Encode(), nil)
        resp, _ := me.httpClient.Do(req)
        body, _ := ioutil.ReadAll(resp.Body)
        resp.Body.Close()
        json.Unmarshal(body, &codeList)
        for _, item := range codeList.Stocks {
            if code, ok := item["code"]; ok {
                list = append(list, code)
            }
        }
        if count == -1 {
            if c, ok := codeList.Count["count"]; ok {
                count = c
            }
        }
        count -= len(codeList.Stocks)
        if size > count {
            params.Set("size", strconv.Itoa(count))
        }
        if count <= 0 {
            break
        }
    }
    return
}

var detailFloatFieldsMap = make(map[string]bool)

func (me *Client) GetDetail(codes string) (list []map[string]interface{}) {
    req, _ := http.NewRequest("GET", apiStockDetail + "?code=" + codes, nil)
    resp, _ := me.httpClient.Do(req)
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    stocks := map[string](map[string]string){}
    json.Unmarshal(body, &stocks)

    if len(detailFloatFieldsMap) == 0 {
        for _, v := range detailFloatFields {
            detailFloatFieldsMap[v] = true
        }
    }

    for _, item := range stocks {
        itemCast := map[string]interface{}{}
        for k, v := range item {
            if _, ok := detailFloatFieldsMap[k]; ok {
                val, err := strconv.ParseFloat(v, 64)
                if err != nil {
                    itemCast[k] = 0
                } else {
                    itemCast[k] = val
                }
            } else {
                itemCast[k] = v
            }
        }

        list = append(list, itemCast)
    }
    return list
}

func New(username, password string) (*Client) {
    jar, _ := cookiejar.New(nil)
    httpClient := &http.Client{
        Jar: jar,
        Timeout: time.Second * 10,
    }
    return &Client{
        jar: jar,
        httpClient: httpClient,
        username: username,
        password: md5hex(password),
    }
}
