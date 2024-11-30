# works-hai-backend

## 実行手順

1. install go and goimports

```
go install golang.org/x/tools/cmd/goimports@latest
```

3. create image directory for tmp file placement

```
mkdir image
```
5. copy config files and fill your neccessary information
```
cd config
mkdir private
cp sample/* private/*
```
6. makefile
```
make
```
8. run ./main and server!!

## for more...

you need to setup your firebase on CORS settings

check your current settings

```
gsutil cors get gs://*****.appspot.com
```

in private dir, run bellow, then your cors.json setting is reflected in your firebase

```
gsutil cors set cors.json gs://*****.appspot.com
```
