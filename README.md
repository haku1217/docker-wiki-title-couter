# Lesson 01

## 環境

使用言語： Python3
イメージ： alpine:3.7

## Image Build

```sh
$ docker build -t test .
```
Image Size: 141MB

## Container Run

```sh
$ docker run --rm -it test aaa,bbb,ccc
```

## 環境

使用言語： GO
イメージ： golang:alpine
実行イメージ：　scratch

## Image Build

```sh
$ docker build -t test -f Dockerfile2 .
```
Image Size: 24.7MB

## Container Run

```sh
$ docker run --rm -it test aaa,bbb,ccc,ddd,eee
```
35秒ほど# docker-wiki-title-couter
