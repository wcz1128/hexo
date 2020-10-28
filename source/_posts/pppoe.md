---
title: pppoe
date: 2020-10-27 09:52:12
tags:
---

### 简介

&emsp;&emsp;随手记录一些docker网络 pppoe拨号 vlan方面的内容

<!-- more -->

# pppoe拨号

&emsp;&emsp;安装后通过pppoe-setup十分方便;其实他只是修改了两个文件

/etc/sysconfig/network-scripts/ifcfg-ppp0
/etc/ppp/chap-secrets

&emsp;&emsp;手动修改这两个文件也是一样的效果

# docker bridge

'''
docker network create -o "com.docker.network.bridge.name"="br200" vlan200
创建一个bridge

ip link add link enp10s0 name enp10s0.200 tpye vlan id 200
ip link set dev enp10s0.200 master br200
ifconfig enp10s0.200 up
把enp10s0网卡设置vlan200，并且加入br200，然后启动

docker run -itd --name vlan200_1 --privileged --network vlan200 mycentos /usr/sbin/init
启动docker，这样网卡已经不用设置vlan了，因为在网桥处已经去掉tag了
'''



# H3C  VLAN

&emsp;&emsp;串口模式下

```
system-view
```
进入system


```
vlan 200
```
创建200的VLAN


```
dis cur
```
显示当前的配置


```
interface GigabitEthernet1/0/1
```
进入端口


```
port access vlan 200
```
or
```
port link-type trunk
port trunk permit vlan 100 200 
```
设置端口vlan属性
