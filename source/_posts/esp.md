---
title: esp
date: 2019-07-27 15:21:25
tags:
  - ESP32
  - ESP8266
  - MicroPython
---

### 简介
&emsp;&emsp;最近买了一些ESP的板子很便宜拿来玩下MicroPython和传感器.

<!-- more -->

# 安装umqtt

import upip
upip.install('micropython-umqtt.simple')

# 远程连接esp8266
可以通过WebREPL实现
```
import webrepl
webrepl.start()
```

在线连接
http://micropython.org/webrepl 

# esp32 cam
连接图如下

![连接图](/image/esp32_1.png)


# esp32 主机开发环境




我的esp-idf有点老，导致virtualenv参数有点对不上，退回老的版本
```
pip uninstall virtualenv
pip install virtualenv==16.7.9
```

另外utf8的判断有点问题，中文就不要decode了
```
#in_str = version_cmd_result.decode()
in_str = version_cmd_result
```


ESP32

下载编译器 esp-idf.tar
解压后
install.sh 
. ./export.sh
就配置完环境了。


