---
title: rsync
date: 2020-06-12 15:26:33
tags:
---

### 简介

&emsp;&emsp;rsync 一些简单记录

<!-- more -->

# 安装AriaNG

&emsp;&emsp;云服务器要到期了，打算备份下里面的东西
&emsp;&emsp;yum -y install rsync
&emsp;&emsp;服务器上配置文件/etc/rsyncd.conf       我就按示例来的
```
# /etc/rsyncd: configuration file for rsync daemon mode
# See rsyncd.conf man page for more options.
# configuration example:

uid = nobody
gid = nobody
use chroot = yes
max connections = 4
pid file = /var/run/rsyncd.pid
exclude = lost+found/
transfer logging = yes
timeout = 900
ignore nonreadable = yes
dont compress   = *.gz *.tgz *.zip *.z *.Z *.rpm *.deb *.bz2

[ftp]
        path = /root/upload
        comment = ftp export area
        read only=yes
        list=no
        auth users=username
        secrets file=/etc/rsyncd.passwd

[www]
        path = /www
        comment = www export area
        read only=yes
        list=no
        auth users=username
        secrets file=/etc/rsyncd.passwd

```

&emsp;&emsp;配置密码
```
echo 'username:123456'>/etc/rsyncd.passwd
chmod 600 /etc/rsyncd.passwd  
```

&emsp;&emsp;客户端更简单
&emsp;&emsp;配置密码
```
echo '123456'>/local/passwd
rsync -auv --password-file=/local/passwd username@ip::www /localdir/www
```
&emsp;&emsp;rsync.conf 必须是600权限



