---
title: Home Assistant
date: 2019-07-08 07:03:41
tags:
  - Home Assistant
---

### 简介

&emsp;&emsp;Home Assistant是一个第三方开源的家庭智能设备接入系统。
&emsp;&emsp;家里为了能统一管理一直用的都是小米设备，但是偶尔也想用用其他的怎么办？试试能不能统一管理。

<!-- more -->

# 安装HASS

&emsp;&emsp;官方建议用python虚拟环境安装，这样可以减少和现有环境的冲突，以后其他项目也可以借鉴下。另外这个虚拟环境是全部拷贝出来的，对于权限管理也会很方便。

```
python3 -m venv homeassistant
cd homeassistant
source bin/activate
python3 -m pip install --upgrade homeassistant
hass --open-ui
hass -c /8T/homeassistant/ --daemon
```
&emsp;&emsp;虚拟环境只是隔离了python库，并不隔离项目，所以默认的配置文件还在~/.homeassistant。启动后会监听8123。可以通过web来管理了。


# 增加设备

&emsp;&emsp;简单学习了下，小米的设备都要token，摄像头需要hack，要接入还是要一番折腾的。暂时放下，有空回来搞

&emsp;&emsp;yeelight在最新的版本中已经默认支持了，只需要在yeelight自己APP中开启局域网即可，配置如下：
```
discovery:
  ignore:
    - yeelight

yeelight:
  devices:
    192.168.3.152:
      name: 小房间灯
```
&emsp;&emsp;默认支持中文也很nice，但是如果我修改了UI配置反而会出不来，所以没去动他

&emsp;&emsp;回来了，今天把小米的都接入了。获取token最便捷的方法就是安装米家[v5.4.54](https://android-apk.org/com.xiaomi.smarthome/43397902-mi-home/),又快又方便。另外有个插曲是，华为路由器客人WIFI，设备间是互相无法访问的。开始设备挂在那个WIFI下，导致一直找不到。

```
fan:
  - platform: xiaomi_miio
    name: 小米2S
    host: 192.168.3.153
    token: xxxxxxxxxxxxxxxxxxxx
    model: zhimi.airpurifier.ma2
remote:
  - platform: xiaomi_miio
    host: 192.168.3.135
    token: xxxxxxxxxxxxxxxxxxxx
    name: 红外遥控
    slot: 1
    timeout: 30
    hidden: false
    commands:
      tv_next_chanle:
        command:
          - raw:Z6VLADYCAACRBgAAxQgAAIkRAAAgIwAACJ0AAFh5AQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA0AAAAEAAAEAAQEBAAEBAQAAAAAAAAAAAAEBAQEBAQEBBQJGAkAA==
```
&emsp;&emsp;红外这个需要对码，可以在网页的 开发者工具->服务 中启动一个remote.learn_command服务。这个时候对码，在通知哪里就有红外码出现了，降红外码放到command里面就可以了。需要使用时候可以新增一个脚本。

```
tv_next_chanle:
  sequence:
    - service: remote.send_command
      entity_id: 'remote.hong_wai_yao_kong'
      data:
        command:
          - 'tv_next_chanle'
```
&emsp;&emsp;另外还增加了一个MQTT的模拟设备（用PC模拟的），可以在界面上操作开关，同时也可以增加其他设备

PC模拟
```
import paho.mqtt.client as mqtt
 
 
def on_connect(client, userdata, flags, rc):
    print("Connected with result code "+str(rc))
    client.subscribe("cmd_test")
 
def on_message(client, userdata, msg):
    print(msg.topic+" " + ":" + str(msg.payload))
    client.publish("state_test",str(msg.payload),2)
 
client = mqtt.Client()
client.on_connect = on_connect
client.on_message = on_message
client.connect(ip, 1883, 60)
client.loop_forever()
```

HASS配置文件
```
light: 
  - platform: mqtt
    name: test_light
    command_topic: "cmd_test"
    state_topic: "state_test"
    optimistic: false

mqtt:
  broker: 127.0.0.1
  username: xxxxx
  password: xxxxx

```
&emsp;&emsp;如果增加了MQTT在开发者工具里也有一个MQTT发送的测试工具

&emsp;&emsp;另外简单的试了下检测，有空再写了

```
device_tracker:
  - platform: ping
    hosts:
      wang: 192.168.3.101

```

# 控客灯泡

&emsp;&emsp;在论坛中看到控客在[GitHub](https://github.com/jedmeng/homeassistant-konke)有第三方支持，而且价格又较小米的便宜，所以入手了一个灯泡试试，（在杭州快递非常快哈哈）。实际用起来才发现这个库作者可能不维护了，所以有各种问题。

&emsp;&emsp;第一个非常明显的问题就是这个库依赖

```
REQUIREMENTS = ['pykonkeio>=2.1.8']
```
而我去[pypi](https://pypi.org/project/pykonkeio/)里面查好像最新的也只到了2.1.7。修改版本要求改到2.1.7。

&emsp;&emsp;第二个应该是我对HASS不熟悉的关系，我看网上文章，以及github作者自己的readme都是 copy the custom_components to your home-assistant config directory.并且按要求配置了configuration.yaml。但是系统始终认不出来，没办法就走了一些偏路。
&emsp;&emsp;我仿照xiaomi_miio在python3.6/site-packages/homeassistant/components/(更新了下变成dist-packages,反正就是库所在目录)目录下建了konke目录，并且里面加了作者的light.py和从xiaomi_miio拷贝来的manifest.json。再次尝试，这次有出错日志，说是找不到uuid。在代码中查了，确实pykonkeio定义的设备就没有uuid（可能是版本对不上？但是我看作者提交那次，其他都没修改，只是修改了版本）。看了下应该就是一个标志符修改成MAC地址吧，但是再跑后发现是空的。分析了下原来没有刷新信息，这个时候MAC还没拿到，再加上刷新设备信息。于是变成以下

light.py
```
async def async_setup_platform(hass, config, async_add_entities, discovery_info=None):
    stat = await device.fetch_info()
    entity = KonkeLight(device, name, model)
    #注释掉LOG

class KonkeLight(Light):
    @property
    def unique_id(self) -> str:
        """Return unique ID for light."""
        if self._model == MODEL_K2_LIGHT:
            return self._device.mac + ':light'
        else:
            return self._device.mac
   #修改uuid为mac
```
configuration.yaml
```
light:
  - platform: konke
    name: konke
    model: kbulb
    host: 192.168.3.177
```

&emsp;&emsp;再启动设备终于出现了，并且开关和亮度也没有问题了。但是返现控制色温有问题。直接调用pykonkeio里面的方法也是一样有问题。所以应该是pykonkeio的问题了，时间有点晚下次再解决吧。

&emsp;&emsp;发现单独设置色温和亮度有问题，一起设置就没问题。比较土的解决方法就是每次设置色温同时也设置一下亮度，问题解决，虽然比较土，但是能用就算了哈，毕竟不是频繁操作。

修改pykonkeio/device/kbulb.py
```
    """
        调整亮度
        req: lan_phone%28-d9-8a-xx-xx-xx%XXXXXXXX%%set#lum#xxx%kbulb
        res: lan_device%28-d9-8a-xx-xx-xx%XXXXXXXX%%set#lum#xxx%kback
    """
    async def set_brightness(self, w, x=-1):
        try:
            utils.check_number(w, 0, 100)
        except ValueError:
            raise error.IllegalValue('brightness should between 0 and 100')

        if x == -1:
            await self.update()
            #print(self.ct)
            await self.set_ct(self.ct, x=1)
        #if self.brightness == int(w):
        #    return

        await self.set_mode(1)
        #await self.send_message('set#lum#%s' % w)
        await self.send_message('set#lum#%s#ctp#5000' % w)
        self.brightness = int(w)
        self.status = 'open'

        if self.brightness == 0:
            await self.turn_off()

    """
        调整色温
        req: lan_phone%28-d9-8a-xx-xx-xx%XXXXXXXX%%set#ctp#xxx%kbulb
        res: lan_device%28-d9-8a-xx-xx-xx%XXXXXXXX%%set#ctp#xxx%kback
    """
    async def set_ct(self, ct, x=-1):
        try:
            utils.check_number(ct, 2700, 6500)
        except ValueError:
            raise error.IllegalValue('illegal color temperature')

        if x == -1:
            await self.update()
            #print(self.brightness)
            await self.set_brightness(self.brightness,x=1)
        #if self.ct != int(ct):

        await self.set_mode(1)
        await self.send_message('set#ctp#%s' % ct)
        self.ct = int(ct)
        self.status = 'open'

        if ct == 0:
            await self.turn_off()
```


&emsp;&emsp;最近还买了ESP32的板子，打算随便玩玩，下次有时间再写了。

# 小爱语音技能

&emsp;&emsp;最近搞了小爱的语音技能，可以通过小爱同学控制家里的其他设备了，大致流程是，

1、通过小爱交互意图识别后，调用小爱提供的函数计算，函数计算给MQTT服务器发送一条指令。
2、HASS新增一个传感器订阅这个指令，
3、增加对应自动化，在传感器值变成不通值时候执行不同的命令。
3、小爱同学进入开发者模式，这样就可以通过类似用XXX打开XXX的方式控制米家以外的设备了。

&emsp;&emsp;当我想在米家做自动化的时候发现，没有上架的调试中的技能是无法调用的，所以米家自动化调用米家外的设备还不行（但是这个却有需求，因为米家的一些设备例如蓝牙锁还加不到HASS中，如果要联动就还不行）


