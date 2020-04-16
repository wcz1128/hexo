---
title: SeLinux
date: 2019-07-17 08:43:27
tags:
  - sshd
  - SeLinux
---

### 随手记
&emsp;&emsp;今天服务器升级ssh服务，但是升级完成后，如果开着SeLinux就无法执行ifconfig等命令。网上搜了很多，都是建议关闭SeLinux。最后发现执行以下就好了。

```
chcon -t bin_t sshd
```

