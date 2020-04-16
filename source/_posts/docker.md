---
title: docker
date: 2019-08-08 03:01:15
tags:
  - docker
---

### 随手记

&emsp;&emsp;简单做了一个Android环境

<!-- more -->

# Dockerfile

```
FROM hub.c.163.com/public/ubuntu:16.04
MAINTAINER xxxx <xxxx@163.com>
WORKDIR /home/xxxx/16.04
COPY ./sources.list /etc/apt/sources.list
COPY ./java-7-openjdk.tar /etc/
COPY ./make-3.81.tar.bz2 /tmp/

RUN sed -i 's/PasswordAuthentication no/PasswordAuthentication yes/' /etc/ssh/sshd_config \
&& sed -i 's/PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config \
&& echo 123456 > tmp \
&& echo 123456 >> tmp \
&& passwd < tmp \
&& rm tmp \
&& apt-get update \
&& apt-get install -y openjdk-8-jdk gcc g++ cpp \
&& apt-get install -y git-core gnupg flex \
&& apt-get install -y bison gperf build-essential curl libc6-dev libssl-dev \
&& apt-get install -y libxml2-utils libgl1-mesa-dev g++-multilib tofrodos python-markdown python-mako x11proto-core-dev \
&& apt-get install -y mkisofs \
&& apt-get install -y bc \
&& apt-get install -y xsltproc dpkg-dev clang ccache zip \
&& cd /tmp/ \
&& tar -xvf /tmp/make-3.81.tar.bz2 \
&& cd /tmp/make-3.81 \
&& ./configure \
&& make && make install

COPY ./java-7-openjdk-amd64.tar /usr/lib/jvm/

EXPOSE 22
CMD [ "/usr/sbin/sshd","-D" ]
```

# build

```
docker build -t android_env/16.04 .
docker run -d --name android -v /local:/in_docker android_env/16.04 
```

# 私有仓库

```
docker pull registry
docker run -d -v /local_dir:/var/lib/registry -p 5000:5000  --name registry registry:latest
```

nginx 配置

```
upstream docker-registry {
	server localhost:5000;
}

server {
	listen port;
	server_name xxx.xx.xx;

	# SSL
	ssl on;
	ssl_certificate /xxx.crt;
	ssl_certificate_key /xxx.key;
      
	client_max_body_size       512m;
	location / {
		proxy_pass http://docker-registry;
		proxy_set_header Host $http_host;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_set_header X-Forwarded-Proto $scheme;
		proxy_read_timeout 900;
	}
}
```

推送
```
docker tag android_env/16.04 docker.xxx.xxx:port/android_env/16.04:latest
docker push docker.xxx.xxx:port/android_env/16.04:latest
```

获取
```
docker pull docker.xxx.xxx:port/android_env/16.04
```

# docker杂

在docker中需要启动vpn，可以自己建立设备节点
```
&& mkdir -p /dev/net \
&& mknod /dev/net/tun c 10 200 \
```

docker中暂时没有找到/proc/sys/net/ipv4/neigh/default/gc_thresh3
需要调整


yum install -y yum-utils device-mapper-persistent-data lvm2
yum-config-manager     --add-repo     https://download.docker.com/linux/centos/docker-ce.repo
yum install docker-ce docker-ce-cli containerd.io
CentOS Linux release 7.6.1810 (Core) 
需要更新
yum install xfsprogs




pip install virtualenv
virtualenv venv
source venv/bin/activate
