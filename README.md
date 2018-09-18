Yandex.XML client: search and limits.

# Client

```
c := yaxml.NewClient()
c.SetAuth("login", "key")
c.SetLR(213)
```

# Search

```
req := yaxml.NewYandexSearchRequest()
req.Query = "hello"

ys, err := c.GetYandexSearch(req)
if err != nil {
	fmt.Println(err)
}

res := ys.Results()
for _, g := range *res {
	fmt.Println(g.Doc[0].URL)
}
```

# Limits

```
ls, err := c.GetLimits()
if err != nil {
	fmt.Println(err)
}

fmt.Println(ls)
fmt.Println(ls.All())
fmt.Println(ls.Now())
fmt.Println(ls.Next())
fmt.Println(ls.RPS())
```