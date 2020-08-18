func mustGet(url string) string{
    resp,err = http.Get(url)
    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()
    var body []byte
    if body,err = ioutil.readAll(resp.Body); err != nil {
        panic(err)
    }
    return string(body)

}