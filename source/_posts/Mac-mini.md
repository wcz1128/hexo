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

&emsp;&emsp;经过熟悉后发现，其实brew下载的软件会在`/opt/homebrew/`目录下建立一个类似linux的根目录。配置文件就在熟悉的etc目录下

&emsp;&emsp;要启动服务可以通过 services相关命令实现

```
brew services --help
```


# hexo
 
&emsp;&emsp;代码提交`hexo d`报如下错误

```
TypeError [ERR_INVALID_ARG_TYPE]: The "mode" argument must be integer. Received an instance of Object
```
&emsp;&emsp;原因是node版本和原本的hexo不符合。可以通过降级node解决，但是M1版本的nvm还下不到M1版本的老版本node，所以反过来，升级hexo一样可以解决。我的方法是重新安装新版本的hexo，然后`hexo init hexo`将里面的`package.json`拷贝出来替换老的配置就好了。
```
{
  "name": "hexo-site",
  "version": "0.0.0",
  "private": true,
  "scripts": {
    "build": "hexo generate",
    "clean": "hexo clean",
    "deploy": "hexo deploy",
    "server": "hexo server"
  },
  "hexo": {
    "version": "5.4.0"
  },
  "dependencies": {
    "hexo": "^5.0.0",
    "hexo-deployer-git": "^3.0.0",
    "hexo-generator-archive": "^1.0.0",
    "hexo-generator-category": "^1.0.0",
    "hexo-generator-index": "^2.0.0",
    "hexo-generator-tag": "^1.0.0",
    "hexo-renderer-ejs": "^1.0.0",
    "hexo-renderer-marked": "^4.0.0",
    "hexo-renderer-stylus": "^2.0.0",
    "hexo-server": "^2.0.0",
    "hexo-theme-landscape": "^0.0.3"
  }
}
```

# 安装IOS应用

&emsp;&emsp;M1一大吸引我的地方是可以安装IOS应用，直接在应用市场下载就可以了。但是有些应用在市场已经下架了。例如微信手机版本。如果要安装，就要通过IPA包的方式安装，获得IPA包的方式网上推荐比较多的是iMazing。我安装了它的免费版，发现它的导出IPA按钮是灰色的。
&emsp;&emsp;最终方法是通过`Apple Configurator 2`,通过它向手机添加APP时候，它会零时将IPA下载到`~/Library/GroupContainersAlias/K36BKF7T3D.group.com.apple.configurator/Library/Caches/Assets/TemporaryItems/MobileApps/`目录下，另外安装完，执行会报没有权限，可以通过`sudo xattr -rd com.apple.quarantine /Applications/xxxxxx`解决

# 挂载tmpfs

&emsp;&emsp;ramdisk 相当于linux的tmpfs
```
diskutil erasevolume HFS+ RamDisk `hdiutil attach -nomount ram://204800`
```


# 接着慢慢玩吧
&emsp;&emsp;对于我这个MacOS小白来说，M1芯片的MacMini还是挺有意思的。

# 开启转发
```
sysctl -w net.inet.ip.forwarding=1
#nat on utun2 from en0:network to 172.19.24.0/24 -> (utun2)
nat on en1 from en0:network -> (en1)
```
