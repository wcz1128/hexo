---
title: VSCode
date: 2019-11-27 14:36:35
tags:
  - VSCode
---
### 简介

&emsp;&emsp;听说VSCode很香，我也来试试。

<!-- more -->

# VSCode环境搭建

&emsp;&emsp;看到有远程功能所以想试下，之前一直ssh+vim。没什么问题，vim插件也够丰富，存粹是为了体验一下VSCode。

&emsp;&emsp;配置和安装都比较人性，点几下就好了。导入Android工程时候遇到一些问题，头文件需要指定下，在工作目录的.vscode/c_cpp_properties.json

```
{
    "configurations": [
        {
            "name": "Linux",
            "includePath": [
                "${workspaceFolder}/**",
                "/work/kernel/include/uapi/"
            ],
            "defines": [],
            "compilerPath": "/usr/bin/clang",
            "cStandard": "c11",
            "cppStandard": "c++17",
            "intelliSenseMode": "clang-x64"
        }
    ],
    "version": 4
}
```


# 后记

&emsp;&emsp;体验了一段时间后，又退回到ssh+vim了，不是说不好用，是改变不了习惯。另外不知道是不是配置的有问题，每次连接后都会在服务器的tmp目录下留下一大堆垃圾，貌似没有自动清理。反正不用，也没再去看配置了。
