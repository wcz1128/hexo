---
title: nfs
date: 2021-11-19 09:36:10
tags: nfs sshfs iscsi
---

### 简介

找了一些网络文件系统试试

<!-- more -->

# 初衷

	在公司下国外代码太费劲，小的代码可以国外下好，打包发回国内速度也能到2~3百K，但是大的代码就由于空间不足，只能挂代理下载了。代理速度也许是小文件的缘故，也许是转发性能的缘故，速度一直上不来。所以想找一个合适的网络文件系统。

# nfs

	最早想到的就是samba和nfs。NFS端口开放比较麻烦，好在有我之前搭建的tinc，目前我所有机器都在同一个内网，所以不存在端口开放的问题。实测了下，对于大文件，基本能把我tinc的速度跑满，但是对于git一些小而量大的文件，还是速度上不来。

# sshfs

	这个是我目前零时用的最多的工具，因为使用太方便了，一条命令解决，但是缺点还是速度跑不满，甚至还不如挂代理来的快。

# iscsi

	这个最近才接触到，缺点貌似听说只能独占硬盘，其实并不属于共享，但是下完了拔过来就完事了哈。

```
apt-get install tgt
tgtadm --lld iscsi --mode target --op new --tid 1 --targetname hippo-storage-iscsi-1
tgt-admin --show  
tgtadm --lld iscsi --mode target --op bind --tid 1 -I ALL
tgt-admin --show  
tgtadm --lld iscsi --mode logicalunit --op new --tid 1 --lun 1 -b /sync57/remote/test.bin
tgt-admin --show  
删除
tgtadm --lld iscsi --op delete --mode logicalunit --tid 1 --lun 1
tgt-admin --show  
tgtadm --lld iscsi --op delete --mode target --tid 1
tgt-admin --show  
```


	客户端
```
 iscsiadm --mode discovery --type sendtargets --portal 10.200.0.1
 iscsiadm -m node --login
 
 iscsiadm -m node -p 10.200.0.1 --logout   
```
