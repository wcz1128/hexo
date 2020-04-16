---
title: 证书制作
date: 2019-08-05 02:32:07
tags:
  - easy-rsa
---

### 简介
&emsp;&emsp;记录下证书制作
<!-- more -->

### 服务器证书

cp easy-rsa3 server
cd server
cp vars.example vars
./easyrsa init-pki
./easyrsa build-ca
./easyrsa gen-req server nopass
./easyrsa sign server server
./easyrsa gen-dh

### 客户端证书

cp easy-rsa3 client
cd client
cp vars.example vars
./easyrsa init-pki
./easyrsa gen-req xxxxxxx

### 签约

cd server
./easyrsa import-req ../client/pki/reqs/xxxxxx.req xxxxxxx
./easyrsa sign client xxxxxx


