---
title: Ubuntu20_04
date: 2020-06-29 13:07:50
tags:
---


### 简介

&emsp;&emsp;今天新装了台电脑，纯记录下

<!-- more -->

# 安装Ubutnu20.04

&emsp;&emsp;今天看到Ubuntu出20.04了，记录下安装遇到的问题。
&emsp;&emsp;第一个老的ssh客户端不再支持新的加密了，修改/etc/ssh/sshd_config，增加

```
KexAlgorithms diffie-hellman-group1-sha1,diffie-hellman-group14-sha1,diffie-hellman-group-exchange-sha1,diffie-hellman-group-exchange-sha256,ecdh-sha2-nistp256,ecdh-sha2-nistp384,ecdh-sha2-nistp521,diffie-hellman-group1-sha1,curve25519-sha256@libssh.org
```

&emsp;&emsp;其实最好是把客户端升级了，这样用老的可能存在安全风险。


# 安装VBOX
&emsp;&emsp;我把安全启动关闭了，否则需要做下MOK。

/etc/apt/sources.list
```
deb [arch=amd64] https://download.virtualbox.org/virtualbox/debian focal contrib
```
安装
```
wget -q https://www.virtualbox.org/download/oracle_vbox_2016.asc -O- | sudo apt-key add -
apt-get update
apt-get install virtualbox-6.1
wget https://download.virtualbox.org/virtualbox/6.1.10/Oracle_VM_VirtualBox_Extension_Pack-6.1.10.vbox-extpack
VBoxManage extpack install Oracle_VM_VirtualBox_Extension_Pack-6.1.10.vbox-extpack
```

create.sh
```
VBoxManage createmedium disk --filename win10.vdi --size 102400
VBoxManage createvm --name windows10 --ostype "Windows10_64" --register
vboxmanage storagectl windows10 --name "SATA Controller" --add sata --controller IntelAHCI
vboxmanage storagectl windows10 --name "IDE Controller" --add ide
vboxmanage modifyvm windows10 --cpus 4 --memory 4096 --vram 256 --hwvirtex on
vboxmanage modifyvm windows10 --ioapic on
vboxmanage modifyvm windows10 --boot1 disk --boot2 dvd --boot3 none --boot4 none
vboxmanage modifyvm windows10 --vrde on
vboxmanage modifyvm windows10 --vrdeport 18701
vboxmanage modifyvm windows10 --vrdeaddress 0.0.0.0
vboxmanage sharedfolder add windows10 --name ubuntu --hostpath /work/VBOX/win10share/
vboxmanage storageattach windows10 --storagectl "IDE Controller" --port 1 --device 0 --type dvddrive --medium /work/ISO/cn_windows_10_consumer_editions_ver
sion_1909_x64_dvd_76365bf8.iso
vboxmanage storageattach windows10 --storagectl "SATA Controller" --port 0 --device 0 --type hdd --medium win10.vdi
vboxmanage modifyvm "windows10" --vrdeauthtype external
vboxmanage modifyvm "windows10" --vrdeauthlibrary VBoxAuthSimple
#vboxmanage internalcommands passwordhash "xxxxxxx"  
vboxmanage setextradata "windows10" "VBoxAuthSimple/users/myname" `vboxmanage internalcommands passwordhash xxxxxxxx | awk '{print $3}'`
```

start.sh
```
vboxmanage startvm windows10 --type=headless
```

stop.sh
```
vboxmanage controlvm windows10 poweroff
```

del.sh
```
VBoxManage unregistervm  --delete windows10
VBoxManage closemedium win10.vdi
```


# 安装docker
&emsp;&emsp;一开始用snap装的docker，后来遇到很多问题，后来换成apt 安装

```
apt-get install docker-io
docker load < xxxx.tar
docker run -itd --name maker --privileged --tmpfs /tmp -v /home/android:/work --ip 172.17.0.20 -p 0.0.0.0:22:22 longene_maker:v1
```

# 安装zabbix
```
wget https://repo.zabbix.com/zabbix/5.0/ubuntu/pool/main/z/zabbix-release/zabbix-release_5.0-1+focal_all.deb
dpkg -i zabbix-release_5.0-1+focal_all.deb
apt update
apt install zabbix-server-mysql zabbix-frontend-php zabbix-nginx-conf zabbix-agent
```

数据库
```
mysql -uroot -p
password
mysql> create database zabbix character set utf8 collate utf8_bin;
mysql> create user zabbix@localhost identified by 'password';
mysql> grant all privileges on zabbix.* to zabbix@localhost;
mysql> quit;
```
 
```
zcat /usr/share/doc/zabbix-server-mysql*/create.sql.gz | mysql -uzabbix -p zabbix
```
/etc/zabbix/zabbix_server.conf
```
DBPassword=password

```
/etc/zabbix/nginx.conf
```
listen 80;
server_name example.com;
```

/etc/zabbix/php-fpm.conf
```
php_value[date.timezone] = Asia/Shanghai
```

```
systemctl restart zabbix-server zabbix-agent nginx php7.4-fpm
systemctl enable zabbix-server zabbix-agent nginx php7.4-fpm
```

默认密码Admin zabbix
