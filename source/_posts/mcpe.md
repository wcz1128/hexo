---
title: 我的世界岩基版搭建 mcpe
date: 2020-04-16 12:49:29
tags: mcpe
---

### 简介
&emsp;&emsp;我的世界岩基版搭建 mcpe

<!-- more -->

# 安装

&emsp;&emsp;倒腾了半天发现用docker是最快的。

&emsp;&emsp;用标准的18.04镜像

```
docker run -d -it --name mcpe -p 19132:19132/udp ubuntu:18.04
docker exec -it mcpe /bin/bash
```

```
apt-get install ssh
apt-get install unzip
apt install libcurl4-openssl-dev
wget https://minecraft.azureedge.net/bin-linux/bedrock-server-1.14.60.5.zip
unzip bedrock-server-1.14.60.5.zip
LD_LIBRARY_PATH=. ./bedrock_server
```
