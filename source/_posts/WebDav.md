---
title: WebDav 搭建
date: 2019-05-24 07:39:49
tags:
  - WebDav
  - NAS
  - Aria2
---

### 简介

&emsp;&emsp;一直想做一个NAS，机器早就买了但是一直没有用起来，机器是华擎的J3455低功耗，高性能，带4个SATA非常适合做NAS。
&emsp;&emsp;目标是做一个私人的云盘，最好能通用的在线播放(现在家里都是百兆宽带，带宽没有问题),需要简单的认证不想公开，给自己带来不必要的麻烦。以前还想着能够下载BT什么的，虽然我也搭建了，但是感觉已经是鸡肋了。现在是一个尊重版权的年代，盗版电影已经不好下了。

<!-- more -->

# NAS服务

&emsp;&emsp;选来选去，最后选择了[WebDav](https://baike.baidu.com/item/WebDAV/4610909),这个是一个比较老的协议，但是好处是Windows系统天然支持，在Linux，Android，IOS各个平台的支持也非常好。所以可以做的很通用。搭建过程非常简单，只要是支持http协议的服务器就可以，一开始我选的是nginx做的，但是效果不是很好，最后发现还是apache的支持比较好，所以最后选择了apache作为web服务器。

```
<Directory /xxxx/xxxx>
        Options Indexes MultiViews
        AllowOverride None
        Require all granted
</Directory>

Alias /xxx /xxxx

<Location /xxxx>
        DAV On
        AuthType Basic
        AuthName "webdav"
        AuthUserFile /xxxx/passwd
        Require valid-user
</Location>
```

### 客户端选择

&emsp;&emsp;Android选择了ES文件浏览器，被归为了FTP一类中可以添加。虽然被曝ES有网络漏洞，但是确实ES是我在Android上使用比较习惯的一个工具软件。主要好处是可以在其他应用直接分享保存到ES，这样省去很多麻烦。

&emsp;&emsp;IOS下次补充，忘记名字了，因为ipad不在身边

&emsp;&emsp;Windows 使用自带的映射网络驱动器即可

&emsp;&emsp;Linux mount -t davfs https://ip:port/dir /mnt  然后输入用户名密码即可

&emsp;&emsp;Mac 下使用自带的资源管理器finder即可


### Windows 50M 文件大小限制

&emsp;&emsp;可以修改注册表 计算机\HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Services\WebClient\Parameters\FileSizeLimitInBytes 来解决，需要重启Webclient服务，或者简单暴力重启系统也可以哈。


# 下载服务

&emsp;&emsp;现在离线下载已经不是什么强需求了，什么东西都可以通过在线解决了，可能对于一些收藏控还有点意义吧，而且现在版权意识也很高，盗版的也很难下载了。我还是简单的搭建了一个[Aria2](https://aria2.github.io)的下载服务器，和[人人影视](http://www.zmz2019.com/)的下载服务器。

### Aria2

&emsp;&emsp;aria2 可以通过 apt 安装，这个只是一个命令行的程序用起来不是很好用。所以搭配了AriaNg这个静态网页一起使用，这样就可以用图形化的界面来方便的管理，具体也没有细用，能下载就可以了，所以也没有遇到多大的坑。

&emsp;&emsp;我做了一个简单的启动脚本
```
aria2c --enable-rpc --rpc-listen-all --rpc-allow-origin-all --rpc-secret=miyao -c --dir /dir -D
```

### 人人影视

&emsp;&emsp;人人影视的安装完全是机缘巧合，因为当时说可以通过做节点共享带宽方式获取收入，所以就下载了。结果挂了好几天，自动下载了好多内容，半分钱也没分到。但是发现用来下载一些需要版权的视频倒是不错的，所以就留下了。（请支持正版，我买了腾讯，优酷，爱奇艺的会员哈哈哈，但是有些很老的，或者想找个国语配音的给小朋友看，确实在线的很难找，发现人人影视的资源还是很丰富的）

# 在线流媒体播放器

&emsp;&emsp;这个一直没有找到合适，我希望既可以网页浏览，又可以在线点播。