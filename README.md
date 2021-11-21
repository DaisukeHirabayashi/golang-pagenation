# golang-pagenation
repository for Qitta article
Qitta Url:

## dependencies
```
gin: version 1.7.4
gorm version 1.9.16
```

## set up
### db set up
```
$ docker-compose up
```

### go set up
```
$ go mod tidy
$ go run main.go
```

### sample url
pagenation only
```
http://localhost:8080/shops?size=2&&page=1
```

pagenation and sort
```
http://localhost:8080/shops?size=2&&page=1&&direction=asc&&orderby=name
```

