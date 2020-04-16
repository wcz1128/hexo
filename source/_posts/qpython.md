---
title: uiautomator2 && qpython
date: 2019-04-25 14:17:29
tags: 
  - Python
  - Android
  - uiautomator
---
### 简介

&emsp;&emsp;[UI Automator](https://developer.android.com/topic/libraries/testing-support-library#UIAutomator)是安卓自带的一个测试框架（现在也有2了）。而uiautomator2是一个方便我们使用的python的库（当然uiautomator也是）。
&emsp;&emsp;[Python](http://www.qpython.com/)一款可以在Android上运行python的小引擎，目前支持python2.7和3.6两个版本。

<!-- more -->

# uiautomator2

&emsp;&emsp;Android自动化测试工具很多，我目前用过appium和uiautomator2。客户端用的都是python。appium风格和Webdriver很像（本来就是用的那套哈），如果以前用的Webdriver基本可以很方便使用。但是appium的服务端是用nodejs写的，在我一台很老的电脑上环境也很乱，装appium时候费了很多力才装上，感觉这点不太友好。当然如果用docker统一环境这也不是什么大问题，可惜我那台老电脑docker也装不上。所以在我第二次使用时候就没有选择appium。而换到了uiautomator上。这个安装相当简单，至少我没有遇到问题。

&emsp;&emsp;但是不管是appium还是uiautomator都是无法测试游戏的。因为游戏对于Android来说都是一块画布，只有一个控件，至于里面多彩的画面是由游戏引擎绘制的，测试框架目前是拿不到的。

# uiautomator基本操作

初始化

```
python -m uiautomator2 init
2019-04-25 15:57:29,527 - __main__.py:327 - INFO - Detect pluged devices: [u'10.2.11.70:5555']
2019-04-25 15:57:29,527 - __main__.py:343 - INFO - Device(10.2.11.70:5555) initialing ...
2019-04-25 15:57:30,082 - __main__.py:133 - INFO - install minicap
2019-04-25 15:57:30,608 - __main__.py:140 - INFO - install minitouch
2019-04-25 15:57:31,078 - __main__.py:155 - INFO - app-uiautomator.apk(1.1.7) downloading ...
2019-04-25 15:57:31,078 - __main__.py:158 - INFO - app-uiautomator-test.apk downloading ...
2019-04-25 15:57:37,129 - __main__.py:350 - INFO - atx-agent is already running, force stop
2019-04-25 15:57:37,684 - __main__.py:213 - INFO - atx-agent(0.5.2) already installed, skip
2019-04-25 15:57:37,887 - __main__.py:254 - INFO - launch atx-agent daemon
2019-04-25 15:57:39,223 - __main__.py:273 - INFO - atx-agent version: 0.5.2
atx-agent output: 2019/04/25 15:57:38 [INFO][github.com/openatx/atx-agent] main.go:508: atx-agent listening on 10.2.11.70:7912
2019-04-25 15:57:39,868 - __main__.py:279 - INFO - success
```

可以看到在我们手机上安装了以下程序

```shell
pm list packages -3                                                                                       
package:com.github.uiautomator
package:com.github.uiautomator.test

ls -l /data/local/tmp/                                                                                                    
total 19912
-rwxr-xr-x 1 shell shell 9525184 2018-12-11 22:34 atx-agent
-rwxr-xr-x 1 shell shell  580128 2019-04-25 14:34 minicap
-rw-rw-rw- 1 shell shell   22576 2019-04-25 14:34 minicap.so
-rwxr-xr-x 1 shell shell   34736 2019-04-25 14:34 minitouch
```

并且启动了`/data/local/tmp/atx-agentserver -d`服务监听7912端口
这个时候我们就可以使用啦

```python
import uiautomator2 as ut2
import time


d = ut2.connect('http://10.2.11.70:7912')

print d.info

d.watchers.remove()
d.watchers.reset()
print d.watchers
d.watcher("yaliceshi").when(resourceId='com.antutu.ABenchMark:id/test_more_stress').click(resourceId='com.antutu.ABenchMark:id/test_more_stress')
d.watcher("kaishiceshi").when(resourceId='com.antutu.ABenchMark:id/test_stress_start').click(resourceId='com.antutu.ABenchMark:id/test_stress_start')
print d.watchers
d.watchers.watched = True
for x in d.watchers:
    print x,d.watcher(x).triggered

d.app_start('com.antutu.ABenchMark')

#sess=d.session('com.antutu.ABenchMark',attach = True)
#print help(sess)
#print sess.running()

print d.dump_hierarchy()

d(text="Settings").wait(timeout=3.0) 

while True:
    for x in d.watchers:
        print x,d.watcher(x).triggered
        if d.watcher(x).triggered:
            d.watchers.remove(x)
    time.sleep(5)

```

个人感觉这个watch非常的好用。dump_hierarchy可以dump出整个布局文件。

# qpython

首先在apk界面点击运行终端,让qpython初始化一下，在第一行可以看到启动命令，以后就可以后台启动了
```
4.4以上系统
/data/data/com.hipipal.qpyplus/files/bin/qpython-android5.sh
在4.4系统下
/data/data/com.hipipal.qpyplus/files/bin/qpython.sh
```

打开pip安装界面
```
pip install humanize
pip install progress
pip install retry
pip install requests
```

从pc拷贝whichcraft.py 和 uiautomator2 到机器上

如果不要前面的python -m uiautomator2 init初始化，也可以手动安装两个包，然后手动启动atx服务

然后连接手机时候不要用adb用http

```
d = u2.connect('http://127.0.0.1:7912')
```
