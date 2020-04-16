---
title: MQTT
date: 2019-07-11 03:59:02
tags:
  - MQTT
  - Erlang
---
### 简介

&emsp;&emsp;昨天捣鼓了HASS，把小米设备和天猫的插座都接了进去，今天想来自己捣鼓下MQTT协议。
&emsp;&emsp;目标很简单自己搭建一个MQTT代理，能够控制家里设备。代理很多，暂时选取了emqttd。（貌似现在很多都用Erlang,而这个环境折腾了我一个上午）

<!-- more -->

# Erlang环境搭建

&emsp;&emsp;不管是我选的emqttd或者RabbitMQ都会用到Erlang，（另外吐槽下emqttd的官网，他们的docker下载链接竟然是无效的）一开始直接apt-get install 安装erlang，或者git代码安装遇到了各种问题，不是版本太低，就是编译出错。后来发现了一个简单的方法。访问[官网](https://www.erlang-solutions.com/resources/download.html)下载deb包,然后安装最新的版本即可。至于rebar3就是个脚本下一个就是了。

# emqx安装

&emsp;&emsp;一个简单的方法。访问[官网](https://www.emqx.io/downloads/broker/)下载。自己编译的话，我在dep时候遇到好多问题，暂时没去搞，毕竟我只是需要服务而已，并不进行优化。

