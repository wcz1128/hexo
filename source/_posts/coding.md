---
title: coding 代码托管和动态网站
date: 2021-03-24 10:35:56
tags:
---


&emsp;&emsp;代码托管的一些记录

<!-- more -->

# 起因

&emsp;&emsp;有很多原因叠加，一来ucloud的虚拟机要到期了，又没找到合适的便宜的主机替代，二来自己家是可以做网站没错，但是访问时候总要带个端口令人不爽，三来最近买了台Macmini做服务器，貌似经常休眠，访问网站第一次时候经常需要等待很长时间，Mac刚玩还不熟悉，所以体验也不太友好。四来公司服务器当然有公网地址，服务器资源也很多，但是公私还是分清楚的好哈。

&emsp;&emsp;家里Mac还是会继续做服务器的，毕竟有些私人的东西需要放家里的，放公网上的东西就托管了吧。


# 虚拟环境

&emsp;&emsp;

```
python3 -m venv myweb
source bin/activate

pip freeze > requirements.txt
```
