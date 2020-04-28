---
title: VirtualBox
date: 2020-04-27 13:37:04
tags:
  - VirtualBox
---
### 简介
&emsp;&emsp;需要一个Windows系统。平时带pad外出时候远程连接使用，
&emsp;&emsp;KVM VMWare VBOX最后选择了VBOX。因为更接近于日常使用，而不是做服务器使用。

<!-- more -->

# 安装VBOX
&emsp;&emsp;这里遇到了不少问题，主要是我的电脑开了SecretBoot，导致无法插入模块。一开始报错也没显示问题，手动插模块才发现。然后按照网上做法进行签名，但是死活不出mok management的界面。最后还是在BIOS里面关了安全启动的选项。
&emsp;&emsp;启动虚拟机后，远程始终连接不上，后来发现远程是由扩展包支持的。我只安装了Base，所以启动不了


```
VBoxManage extpack install Oracle_VM_VirtualBox_Extension_Pack-6.1.6.vbox-extpack 
```

&emsp;&emsp;安装后就可以远程安装系统了。

```
VBoxManage createmedium disk --filename win10.vdi --size 102400
VBoxManage createvm --name windows10 --ostype "Windows10_64" --register
vboxmanage storagectl windows10 --name "SATA Controller" --add sata --controller IntelAHCI
vboxmanage storagectl windows10 --name "IDE Controller" --add ide
vboxmanage modifyvm windows10 --cpus 4 --memory 4096 --vram 256 --hwvirtex on
vboxmanage modifyvm windows10 --ioapic on
vboxmanage modifyvm windows10 --boot1 disk --boot2 dvd --boot3 none --boot4 none
vboxmanage modifyvm windows10 --vrde on
vboxmanage modifyvm windows10 --vrdeport 3389
vboxmanage modifyvm windows10 --vrdeaddress 0.0.0.0
VBoxManage  modifyvm windows10 --natpf1  rdp,tcp,,3388,,3389
vboxmanage storageattach windows10 --storagectl "IDE Controller" --port 1 --device 0 --type dvddrive --medium /work1/tools/ISO/cn_windows_10_consumer_editions_version_1909_x64_dvd_76365bf8.iso
vboxmanage storageattach windows10 --storagectl "SATA Controller" --port 0 --device 0 --type hdd --medium win10.vdi

#vboxmanage startvm windows10 --type=headless
#vboxmanage controlvm windows10 poweroff
#vboxmanage controlvm windows10 reset
#vboxmanage controlvm windows10 screenshotpng test.png
#VBoxManage closemedium win10.vdi
#VBoxManage unregistervm  --delete windows10
```
