---
title: Mac mini
date: 2021-03-11 13:10:34
tags: 
  - Mac mini
---
### Mac mini

&emsp;&emsp;Mac mini 购入后的一些配置

<!-- more -->

# Mac mini

&emsp;&emsp;最近家里的NAS挂了，想着继续折腾mini-itx主板的，一时也没找到合适的。前端时间不是M1 芯片的Mac mini发布了么，按照文档说的待机功耗才6W，性能又强悍，这个不就是我的要的家庭服务器理想的选择么哈哈。之前也想过rk3322这类ARM开发板作为中心，但是没有雷电接口，扩展外置硬盘总会有瓶颈。这次看到M1 Mac mini想着拿回来试试哈哈。

# brew

&emsp;&emsp;作为苹果小白，之前从未接触过MacOS。装brew就花了半天。因为是M1芯片同时，网络环境也不好。最后总结以下方法是最靠谱的

```
sudo mkdir -p /opt/homebrew
sudo chown -R $(whoami) /opt/homebrew
cd /opt
curl -L https://github.com/Homebrew/brew/tarball/master | tar xz --strip 1 -C homebrew
export PATH=$PATH:/opt/homebrew/bin/:/opt/homebrew/sbin/
```

&emsp;&emsp;通过brew就可以安装很多工具了，例如

```
brew install proxychains-ng
brew install v2ray
```

&emsp;&emsp;经过熟悉后发现，其实brew下载的软件会在<font color='red'>/opt/homebrew/</font>目录下建立一个类似linux的根目录。配置文件就在熟悉的etc目录下

&emsp;&emsp;要启动服务可以通过 services相关命令实现

```
brew services --help
```


