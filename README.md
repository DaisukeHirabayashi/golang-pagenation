# golang-pagenation
repository for Qitta article
Qitta Url:

## dependencies
```
gin: version 1.7.4
gorm version 1.9.16
```

## sample usage
### set up
#### db set up
```
$ docker-compose up
```

#### go set up
```
$ go mod tidy
$ go run main.go
```

## sample url
pagenation only
```
http://localhost:8080/shops?size=2&&page=1
```

pagenation and sort
```
http://localhost:8080/shops?size=2&&page=1&&direction=asc&&orderby=name
```

SampleResponse
```
{
  "page": {
    "number": 2, # pageNumber
    "size": 2, #ncontents size
    "total_elements": 7, # total size
    "total_pages": 4 # total pages
  },
  "shops": [
    {
      "id": 7,
      "name": "Beauty-Beauty",
      "created_at": "2021-11-21T16:42:43.655756Z"
    },
    {
      "id": 3,
      "name": "Ceauty-Beauty",
      "created_at": "2021-11-21T16:42:43.651779Z"
    }
  ]
}
```
