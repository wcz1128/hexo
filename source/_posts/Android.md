---
title: Android 一些记录 
date: 2021-04-22 13:37:35 
tags: Android
---

### 简介

  Android 一些简单记录

<!-- more -->

# vpn

     受保护的socket会最终通过`setsockopt(*socketFd, SOL_SOCKET, SO_MARK,...)`设置标志位，并且这个标志位并不在数据包内，而是内核的一个标志。

frameworks/base/core/java/android/net/VpnService.java

```
    /**
     * Protect a socket from VPN connections. After protecting, data sent
     * through this socket will go directly to the underlying network,
     * so its traffic will not be forwarded through the VPN.
     * This method is useful if some connections need to be kept
     * outside of VPN. For example, a VPN tunnel should protect itself if its
     * destination is covered by VPN routes. Otherwise its outgoing packets
     * will be sent back to the VPN interface and cause an infinite loop. This
     * method will fail if the application is not prepared or is revoked.
     *
     * <p class="note">The socket is NOT closed by this method.
     *
     * @return {@code true} on success.
     */
    public boolean protect(int socket) {
        return NetworkUtils.protectFromVpn(socket);
    }

    /**
     * Convenience method to protect a {@link Socket} from VPN connections.
     *
     * @return {@code true} on success.
     * @see #protect(int)
     */
    public boolean protect(Socket socket) {
        return protect(socket.getFileDescriptor$().getInt$());
    }

```

frameworks/base/core/java/android/net/NetworkUtils.java

```
    /**
     * Protect {@code fd} from VPN connections.  After protecting, data sent through
     * this socket will go directly to the underlying network, so its traffic will not be
     * forwarded through the VPN.
     */
    public static boolean protectFromVpn(FileDescriptor fd) {
        return protectFromVpn(fd.getInt$());
    }

    /**
     * Protect {@code socketfd} from VPN connections.  After protecting, data sent through
     * this socket will go directly to the underlying network, so its traffic will not be
     * forwarded through the VPN.
     */
    public native static boolean protectFromVpn(int socketfd);
```

frameworks/base/core/jni/android_net_NetUtils.cpp

```
{ "protectFromVpn", "(I)Z", (void*)android_net_utils_protectFromVpn },
static jboolean android_net_utils_protectFromVpn(JNIEnv *env, jobject thiz, jint socket)
{
    return (jboolean) !protectFromVpn(socket);
}
```

system/netd/client/NetdClient.cpp

```
extern "C" int protectFromVpn(int socketFd) {
    if (socketFd < 0) {
        return -EBADF;
    }
    FwmarkCommand command = {FwmarkCommand::PROTECT_FROM_VPN, 0, 0};
    return FwmarkClient().send(&command, socketFd, nullptr);
}

```

system/netd/server/FwmarkServer.cpp

```
int FwmarkServer::processClient(SocketClient* client, int* socketFd) {
    FwmarkCommand command;
    FwmarkConnectInfo connectInfo;

    iovec iov[2] = {
        { &command, sizeof(command) },
        { &connectInfo, sizeof(connectInfo) },
    };
    msghdr message;
    memset(&message, 0, sizeof(message));
    message.msg_iov = iov;
    message.msg_iovlen = ARRAY_SIZE(iov);

    union {
        cmsghdr cmh;
        char cmsg[CMSG_SPACE(sizeof(*socketFd))];
    } cmsgu;

    memset(cmsgu.cmsg, 0, sizeof(cmsgu.cmsg));
    message.msg_control = cmsgu.cmsg;
    message.msg_controllen = sizeof(cmsgu.cmsg);

    int messageLength = TEMP_FAILURE_RETRY(recvmsg(client->getSocket(), &message, 0));
    if (messageLength <= 0) {
        return -errno;
    }

    if (!((command.cmdId != FwmarkCommand::ON_CONNECT_COMPLETE && messageLength == sizeof(command))
            || (command.cmdId == FwmarkCommand::ON_CONNECT_COMPLETE
            && messageLength == sizeof(command) + sizeof(connectInfo)))) {
        return -EBADMSG;
    }

    Permission permission = mNetworkController->getPermissionForUser(client->getUid());

    if (command.cmdId == FwmarkCommand::QUERY_USER_ACCESS) {
        if ((permission & PERMISSION_SYSTEM) != PERMISSION_SYSTEM) {
            return -EPERM;
        }
        return mNetworkController->checkUserNetworkAccess(command.uid, command.netId);
    }

    cmsghdr* const cmsgh = CMSG_FIRSTHDR(&message);
    if (cmsgh && cmsgh->cmsg_level == SOL_SOCKET && cmsgh->cmsg_type == SCM_RIGHTS &&
        cmsgh->cmsg_len == CMSG_LEN(sizeof(*socketFd))) {
        memcpy(socketFd, CMSG_DATA(cmsgh), sizeof(*socketFd));
    }

    if (*socketFd < 0) {
        return -EBADF;
    }

    Fwmark fwmark;
    socklen_t fwmarkLen = sizeof(fwmark.intValue);
    if (getsockopt(*socketFd, SOL_SOCKET, SO_MARK, &fwmark.intValue, &fwmarkLen) == -1) {
        return -errno;
    }

    switch (command.cmdId) {
        case FwmarkCommand::ON_ACCEPT: {
            // Called after a socket accept(). The kernel would've marked the NetId and necessary
            // permissions bits, so we just add the rest of the user's permissions here.
            permission = static_cast<Permission>(permission | fwmark.permission);
            break;
        }

        case FwmarkCommand::ON_CONNECT: {
            // Called before a socket connect() happens. Set an appropriate NetId into the fwmark so
            // that the socket routes consistently over that network. Do this even if the socket
            // already has a NetId, so that calling connect() multiple times still works.
            //
            // But if the explicit bit was set, the existing NetId was explicitly preferred (and not
            // a case of connect() being called multiple times). Don't reset the NetId in that case.
            //
            // An "appropriate" NetId is the NetId of a bypassable VPN that applies to the user, or
            // failing that, the default network. We'll never set the NetId of a secure VPN here.
            // See the comments in the implementation of getNetworkForConnect() for more details.
            //
            // If the protect bit is set, this could be either a system proxy (e.g.: the dns proxy
            // or the download manager) acting on behalf of another user, or a VPN provider. If it's
            // a proxy, we shouldn't reset the NetId. If it's a VPN provider, we should set the
            // default network's NetId.
            //
            // There's no easy way to tell the difference between a proxy and a VPN app. We can't
            // use PERMISSION_SYSTEM to identify the proxy because a VPN app may also have those
            // permissions. So we use the following heuristic:
            //
            // If it's a proxy, but the existing NetId is not a VPN, that means the user (that the
            // proxy is acting on behalf of) is not subject to a VPN, so the proxy must have picked
            // the default network's NetId. So, it's okay to replace that with the current default
            // network's NetId (which in all likelihood is the same).
            //
            // Conversely, if it's a VPN provider, the existing NetId cannot be a VPN. The only time
            // we set a VPN's NetId into a socket without setting the explicit bit is here, in
            // ON_CONNECT, but we won't do that if the socket has the protect bit set. If the VPN
            // provider connect()ed (and got the VPN NetId set) and then called protect(), we
            // would've unset the NetId in PROTECT_FROM_VPN below.
            //
            // So, overall (when the explicit bit is not set but the protect bit is set), if the
            // existing NetId is a VPN, don't reset it. Else, set the default network's NetId.
            if (!fwmark.explicitlySelected) {
                if (!fwmark.protectedFromVpn) {
                    fwmark.netId = mNetworkController->getNetworkForConnect(client->getUid());
                } else if (!mNetworkController->isVirtualNetwork(fwmark.netId)) {
                    fwmark.netId = mNetworkController->getDefaultNetwork();
                }
            }
            break;
        }

        case FwmarkCommand::ON_CONNECT_COMPLETE: {
            // Called after a socket connect() completes.
            // This reports connect event including netId, destination IP address, destination port,
            // uid, connect latency, and connect errno if any.

            // Skip reporting if connect() happened on a UDP socket.
            int socketProto;
            socklen_t intSize = sizeof(socketProto);
            const int ret = getsockopt(*socketFd, SOL_SOCKET, SO_PROTOCOL, &socketProto, &intSize);
            if ((ret != 0) || (socketProto == IPPROTO_UDP)) {
                break;
            }

            android::sp<android::net::metrics::INetdEventListener> netdEventListener =
                    mEventReporter->getNetdEventListener();

            if (netdEventListener != nullptr) {
                char addrstr[INET6_ADDRSTRLEN];
                char portstr[sizeof("65536")];
                const int ret = getnameinfo((sockaddr*) &connectInfo.addr, sizeof(connectInfo.addr),
                        addrstr, sizeof(addrstr), portstr, sizeof(portstr),
                        NI_NUMERICHOST | NI_NUMERICSERV);

                netdEventListener->onConnectEvent(fwmark.netId, connectInfo.error,
                        connectInfo.latencyMs,
                        (ret == 0) ? String16(addrstr) : String16(""),
                        (ret == 0) ? strtoul(portstr, NULL, 10) : 0, client->getUid());
            }
            break;
        }

        case FwmarkCommand::SELECT_NETWORK: {
            fwmark.netId = command.netId;
            if (command.netId == NETID_UNSET) {
                fwmark.explicitlySelected = false;
                fwmark.protectedFromVpn = false;
                permission = PERMISSION_NONE;
            } else {
                if (int ret = mNetworkController->checkUserNetworkAccess(client->getUid(),
                                                                         command.netId)) {
                    return ret;
                }
                fwmark.explicitlySelected = true;
                fwmark.protectedFromVpn = mNetworkController->canProtect(client->getUid());
            }
            break;
        }

        case FwmarkCommand::PROTECT_FROM_VPN: {
            if (!mNetworkController->canProtect(client->getUid())) {
                return -EPERM;
            }
            // If a bypassable VPN's provider app calls connect() and then protect(), it will end up
            // with a socket that looks like that of a system proxy but is not (see comments for
            // ON_CONNECT above). So, reset the NetId.
            //
            // In any case, it's appropriate that if the socket has an implicit VPN NetId mark, the
            // PROTECT_FROM_VPN command should unset it.
            if (!fwmark.explicitlySelected && mNetworkController->isVirtualNetwork(fwmark.netId)) {
                fwmark.netId = mNetworkController->getDefaultNetwork();
            }
            fwmark.protectedFromVpn = true;
            permission = static_cast<Permission>(permission | fwmark.permission);
            break;
        }

        case FwmarkCommand::SELECT_FOR_USER: {
            if ((permission & PERMISSION_SYSTEM) != PERMISSION_SYSTEM) {
                return -EPERM;
            }
            fwmark.netId = mNetworkController->getNetworkForUser(command.uid);
            fwmark.protectedFromVpn = true;
            break;
        }

        default: {
            // unknown command
            return -EPROTO;
        }
    }

    fwmark.permission = permission;

    if (setsockopt(*socketFd, SOL_SOCKET, SO_MARK, &fwmark.intValue,
                   sizeof(fwmark.intValue)) == -1) {
        return -errno;
    }

    return 0;
}
```

# 输入法

AndroidManifest.xml 

这里面指定的activity是主页上点击的，和设置里面点击的不一样。

```
<!--
 Copyright (C) 2008-2012  OMRON SOFTWARE Co., Ltd.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
-->

<manifest xmlns:android="http://schemas.android.com/apk/res/android"
          package="my.com.cn.myime">

    <original-package android:name="my.com.cn.myime" />

        <uses-sdk android:minSdkVersion="19" />
    <uses-permission xmlns:android="http://schemas.android.com/apk/res/android" android:name="android.permission.VIBRATE"/>
    <uses-permission android:name="android.permission.INTERNET"></uses-permission>
        <uses-permission android:name="android.permission.WRITE_SECURE_SETTINGS" />
    <application android:label="我的输入法"
                android:icon="@drawable/ic_launcher"
        >

        <service android:name="MyIME" 
                                android:label="我的输入法"
                                android:permission="android.permission.BIND_INPUT_METHOD"
                 >
            <intent-filter>
                <action android:name="android.view.InputMethod" />
            </intent-filter>
            <meta-data android:name="android.view.im" android:resource="@xml/method" />
        </service>
                <activity android:name="MainActivity">
            <intent-filter>
                <action android:name="android.intent.action.MAIN" />
                <category android:name="android.intent.category.LAUNCHER" />
            </intent-filter>
        </activity>
    </application>

</manifest> 
```

res/xml/method.xml

这个里面主要settingsActivity是指定在设置界面内点击输入法的方法

```
<?xml version="1.0" encoding="utf-8"?>                                                                                                                 
<!--
 Copyright (C) 2008-2012  OMRON SOFTWARE Co., Ltd.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0
 
  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
 -->
<!-- The attributes in this XML file provide configuration information -->
<!-- for the Search Manager. -->

<input-method xmlns:android="http://schemas.android.com/apk/res/android"
                          android:settingsActivity="my.com.cn.myime.MainActivity"
              >
</input-method>
```

输入法如果不实现InputView在某些情况下将会有一些问题。

```
        @Override public View onCreateInputView() {
             
                return this.getLayoutInflater().inflate(R.layout.keyboard_default_main, null);
                
        }
```
