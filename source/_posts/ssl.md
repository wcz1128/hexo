---
title: ssl证书
date: 2021-06-04 14:36:07
tags: certbot
---

### 简介

&emsp;&emsp;  certbot 是一个3个月免费的ssl证书签发网站

<!-- more -->

&emsp;&emsp; 阿里云貌似已经没有免费的一年的SSL证书签发了。七牛华为等还有，估计也要收费了。


&emsp;&emsp;  certbot 使用还是挺方便的。
```
certbot certonly  -d "*.wcz.pub" --manual --preferred-challenges dns-01
```
&emsp;&emsp; 成功后证书在/etc/letsencrypt/live/wcz.pub 目录下

&emsp;&emsp; 更新证书
```
certbot renew
```
