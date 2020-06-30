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

修改字体 替换 /usr/share/zabbix/assets/fonts/graphfont.ttf
默认密码Admin zabbix

# 安装nextcloud
```
upstream php-handler {
    server 127.0.0.1:9000;
    #server unix:/var/run/php/php7.2-fpm.sock;
}

server {
    listen 18703 ssl http2;
    listen [::]:18703 ssl http2;
    server_name test.com;

    # Use Mozilla's guidelines for SSL/TLS settings
    # https://mozilla.github.io/server-side-tls/ssl-config-generator/
    # NOTE: some settings below might be redundant
    ssl_certificate /etc/nginx/test.com.crt;
    ssl_certificate_key /etc/nginx/test.com.key;

    # Add headers to serve security related headers
    # Before enabling Strict-Transport-Security headers please read into this
    # topic first.
    #add_header Strict-Transport-Security "max-age=15768000; includeSubDomains; preload;" always;
    #
    # WARNING: Only add the preload option once you read about
    # the consequences in https://hstspreload.org/. This option
    # will add the domain to a hardcoded list that is shipped
    # in all major browsers and getting removed from this list
    # could take several months.
    add_header Referrer-Policy "no-referrer" always;
    add_header X-Content-Type-Options "nosniff" always;
    add_header X-Download-Options "noopen" always;
    add_header X-Frame-Options "SAMEORIGIN" always;
    add_header X-Permitted-Cross-Domain-Policies "none" always;
    add_header X-Robots-Tag "none" always;
    add_header X-XSS-Protection "1; mode=block" always;

    # Remove X-Powered-By, which is an information leak
    fastcgi_hide_header X-Powered-By;

    # Path to the root of your installation
    root /var/www/nextcloud;

    location = /robots.txt {
        allow all;
        log_not_found off;
        access_log off;
    }

    # The following 2 rules are only needed for the user_webfinger app.
    # Uncomment it if you're planning to use this app.
    #rewrite ^/.well-known/host-meta /public.php?service=host-meta last;
    #rewrite ^/.well-known/host-meta.json /public.php?service=host-meta-json last;

    # The following rule is only needed for the Social app.
    # Uncomment it if you're planning to use this app.
    #rewrite ^/.well-known/webfinger /public.php?service=webfinger last;

    location = /.well-known/carddav {
      return 301 $scheme://$host:$server_port/remote.php/dav;
    }
    location = /.well-known/caldav {
      return 301 $scheme://$host:$server_port/remote.php/dav;
    }

    # set max upload size
    client_max_body_size 512M;
    fastcgi_buffers 64 4K;

    # Enable gzip but do not remove ETag headers
    gzip on;
    gzip_vary on;
    gzip_comp_level 4;
    gzip_min_length 256;
    gzip_proxied expired no-cache no-store private no_last_modified no_etag auth;
    gzip_types application/atom+xml application/javascript application/json application/ld+json application/manifest+json application/rss+xml application/vnd.geo+json application/vnd.ms-fontobject application/x-font-ttf application/x-web-app-manifest+json application/xhtml+xml application/xml font/opentype image/bmp image/svg+xml image/x-icon text/cache-manifest text/css text/plain text/vcard text/vnd.rim.location.xloc text/vtt text/x-component text/x-cross-domain-policy;

    # Uncomment if your server is build with the ngx_pagespeed module
    # This module is currently not supported.
    #pagespeed off;

    location / {
        rewrite ^ /index.php;
    }

    location ~ ^\/(?:build|tests|config|lib|3rdparty|templates|data)\/ {
        deny all;
    }
    location ~ ^\/(?:\.|autotest|occ|issue|indie|db_|console) {
        deny all;
    }

    location ~ ^\/(?:index|remote|public|cron|core\/ajax\/update|status|ocs\/v[12]|updater\/.+|oc[ms]-provider\/.+|.+\/richdocumentscode\/proxy)\.php(?:$|\/) {
        fastcgi_split_path_info ^(.+?\.php)(\/.*|)$;
        set $path_info $fastcgi_path_info;
        try_files $fastcgi_script_name =404;
        include fastcgi_params;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        fastcgi_param PATH_INFO $path_info;
        fastcgi_param HTTPS on;
        # Avoid sending the security headers twice
        fastcgi_param modHeadersAvailable true;
        # Enable pretty urls
        fastcgi_param front_controller_active true;
        fastcgi_pass php-handler;
        fastcgi_intercept_errors on;
        fastcgi_request_buffering off;
    }

    location ~ ^\/(?:updater|oc[ms]-provider)(?:$|\/) {
        try_files $uri/ =404;
        index index.php;
    }

    # Adding the cache control header for js, css and map files
    # Make sure it is BELOW the PHP block
    location ~ \.(?:css|js|woff2?|svg|gif|map)$ {
        try_files $uri /index.php$request_uri;
        add_header Cache-Control "public, max-age=15778463";
        # Add headers to serve security related headers (It is intended to
        # have those duplicated to the ones above)
        # Before enabling Strict-Transport-Security headers please read into
        # this topic first.
        #add_header Strict-Transport-Security "max-age=15768000; includeSubDomains; preload;" always;
        #
        # WARNING: Only add the preload option once you read about
        # the consequences in https://hstspreload.org/. This option
        # will add the domain to a hardcoded list that is shipped
        # in all major browsers and getting removed from this list
        # could take several months.
        add_header Referrer-Policy "no-referrer" always;
        add_header X-Content-Type-Options "nosniff" always;
        add_header X-Download-Options "noopen" always;
        add_header X-Frame-Options "SAMEORIGIN" always;
        add_header X-Permitted-Cross-Domain-Policies "none" always;
        add_header X-Robots-Tag "none" always;
        add_header X-XSS-Protection "1; mode=block" always;

        # Optional: Don't log access to assets
        access_log off;
    }

    location ~ \.(?:png|html|ttf|ico|jpg|jpeg|bcmap|mp4|webm)$ {
        try_files $uri /index.php$request_uri;
        # Optional: Don't log access to other assets
        access_log off;
    }
}
```
数据库
```
create database nextcloud;
create user namexxxx@localhost identified by 'mimaxxxxx';
grant all privileges on nextcloud.* to namexxxx@localhost identified by 'mimaxxxxx'; flush privileges; 
```

修改目录权限www-data
```
apt-get install php-zip php-curl  php-intl php-gmp php-imagick php-redis
```

/etc/php/7.4/fpm/php.ini
```
memory_limit = 1024M
```

/etc/php/7.4/fpm/pool.d/www.conf
```
env[PATH] = /usr/local/bin:/usr/local/sbin:/usr/bin:/usr/sbin:/bin:/sbin:/usr/bin/php
```

nextcloud/config/config.php
```
  'memcache.local' => '\OC\Memcache\Redis',
  'redis' => array(
          'host' => 'localhost',
          'port' => 1128,
  ),`
```
