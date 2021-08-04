---
title: tinc
date: 2021-06-16 16:07:50
tags: tinc
---
### 简介

&emsp;&emsp; tinc 一个简单的组网工具

<!-- more -->

# 配置

&emsp;&emsp; tinc 不存在真正的服务器端，所有节点都是对等的。
&emsp;&emsp;
&emsp;&emsp; 在/etc/tinc/目录下建立你的网络名称目录,例如我叫MyNet

```
[root@10-9-86-69 MyNet]# tree /etc/tinc/MyNet/
/etc/tinc/MyNet/
├── hosts
│   ├── MAC
│   ├── QG
│   ├── ROOT
│   ├── SP
│   ├── TX
│   ├── UCLOUD
│   └── US
├── rsa_key.priv
├── tinc.conf
├── tinc-down
└── tinc-up
```

&emsp;&emsp;可以看到我有7个节点，每个节点配置如下

```
Address = 114.114.114.114 10080  #有就填写公网地址，没有则不需要
Subnet = 10.200.0.4/32
```

&emsp;&emsp; 创建tinc.conf
```
Name = SP  #当前节点的名称，必填，其余都选填
Interface = tinc  #生成网卡名称
Mode = switch   #模式  router switch hub  默认route
Compression=9   #压缩等级
Cipher  = aes-256-cbc
Digest = sha256
PrivateKeyFile = /etc/tinc/MyNet/rsa_key.priv
```

&emsp;&emsp; tincd -n MyNet -K4096 生成密钥对,会把公钥写入hosts目录下节点名字的配置文件,节点配置文件就变成了。
```
Address = 114.114.114.114 10080  #有就填写公网地址，没有则不需要
Subnet = 10.200.0.4/32
-----BEGIN RSA PUBLIC KEY-----
XXXXXXXXXXXXXXXX
-----END RSA PUBLIC KEY-----
```



&emsp;&emsp; down 和 up 文件是连接成功后执行的
```
[root@10-9-86-69 MyNet]# cat tinc-down 
#!/bin/sh
ifconfig $INTERFACE down
[root@10-9-86-69 MyNet]# cat tinc-up
#!/bin/sh
ifconfig $INTERFACE 10.200.0.4 netmask 255.255.255.0
```



&emsp;&emsp; m1 MacOS 已经自带了utun设备，安装tuntap也装不上，修改下设备类型
```
DeviceType = utun
ifconfig $INTERFACE 10.200.0.2 10.200.0.1 up
sudo route -n add   -net 10.200.0.0/16 10.200.0.2
```

