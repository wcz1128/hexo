---
title: hexo搭建
date: 2019-04-24 10:41:58
tags: hexo
---

### 随便写的

&emsp;&emsp;一直很懒，那么大了也没留下什么记录,都是一时新鲜玩这个玩那个，除了工作需要项目持续时间没有一个超过一个月的哈哈。希望这个能让我一直坚持下去。其实原因也有很多，有想法了，去做了，休息下，等再做时候发现已经不是那个味道了，所以就暂停了。技术的进步还是飞速的，时刻都有新的技术，以前想法很有可能就不合时宜了。

<!-- more -->

# hexo搭建

&emsp;&emsp;搭建教程[官网](https://hexo.io/)和网上已经有很多了,现在写的也不一定适合以后，所以就写些我遇到的问题和处理

1. GitHub Pages

&emsp;&emsp;网上很多说托管到GitHub,需要建立一个和username一样的Repository，后来我发现并不需要，可以随意取一个名字，然后在这个仓库的Settings中把GitHub Pages开起来就好了。

2. 多个Pages服务

&emsp;&emsp;如果需要同时在github和腾讯部署需要用到两个git仓库。一开始我是这样的

```
deploy:
  type: git
  repo:
    github: git@github.com:xxxxx/hexo.git
    tencent: git@git.dev.tencent.com:xxxxx/hexo.git
  branch: master
```

&emsp;&emsp;后来发现Github上就没有master分支了，而是gh-pages分支，我感觉应该改成以下，但是不影响使用我就没去修改，没试过

```
deploy:
  type: git
  repo:
    github: git@github.com:xxxxx/hexo.git
  branch: master
  repo:
    tencent: git@git.dev.tencent.com:xxxxx/hexo.git
  branch: master
```

3. 腾讯https证书申请

&emsp;&emsp;在开启腾讯的Pages服务后，证书一直申请失败返回错误，可能和我配置了海外域名指向Github有关，暂时把海外的解析去掉了。但是发现证书只有三个月，开启境外域名到时候可能还是会有问题

4. 文章太长在首页不想全部显示

&emsp;&emsp;可以在文章中加入`<!-- more -->`，这样首页就会显示一个 readmore

5. 多地编辑

&emsp;&emsp;再增加一个分支，并且设置为默认分支。将代码(除去.deploy_git .DS_Store Thumbs.db db.json *.log node_modules/ public/)以外代码传至新分支。新机器只要git clone git@github.com:xxxxx.git。进入克隆后的目录`npm install` 安装必要的库(之前要先install hexo)，然后就可以`hexo g;hexo d`发布了。新机器改完后需要`git commit;git push` 将新的代码传回新建分支即可。代码在新分支，Pages用的是主分支。

6. 懒

&emsp;&emsp;还是有一些忘记了，不管了就写到这里吧，默认的主题也挺好看的就不换了。

