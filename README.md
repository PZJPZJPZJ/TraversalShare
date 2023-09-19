# TraversalShare
## 介绍
打洞型点对点互传工具
## 实现目标
使用NAT穿透技术，通过STUN服务器检测NAT类型，筛选最优方案进行文字、文件等点对点互传通信
## 穿透方案
从上至下优先度递减
### IPv6
#### 无防火墙IPv6-任意IPv6
直连
#### 有防火墙IPv6-有防火墙IPv6
双方均进行防火墙打洞
### IPv4
#### NAT1-任意NAT
NAT1向STUN服务器获取公网地址端口，转发到对方客户端
#### NAT2-NAT3
##### NAT2(192.168.0.1:2000)->CGNAT2(2.2.2.2:12000)->STUN[转发到NAT3客户端]
##### NAT3(192.168.0.1:3000)->CGNAT3(3.3.3.3:13000)->STUN[转发到NAT2客户端]
##### NAT2(192.168.0.1:2000)->CGNAT2(2.2.2.2:12000)-*>CGNAT3(3.3.3.3:13000)
##### NAT3(192.168.0.1:3000)->CGNAT3(3.3.3.3:13000)->CGNAT2(2.2.2.2:12000)->NAT2(192.168.0.1:2000)[持续发送心跳包保活]
##### NAT2(192.168.0.1:2000)->CGNAT2(2.2.2.2:12000)->CGNAT3(3.3.3.3:13000)->NAT3(192.168.0.1:3000)[持续发送心跳包保活]
#### NAT2-NAT4
##### NAT2(192.168.0.1:2000)->CGNAT2(2.2.2.2:12000)->STUN[转发到NAT3客户端]
##### NAT4(192.168.0.1:4000)->CGNAT3(4.4.4.4:14000)->STUN[转发到NAT2客户端]
##### NAT2(192.168.0.1:2000)->CGNAT2(2.2.2.2:12000)-*>CGNAT4(4.4.4.4:14000)
##### NAT4(192.168.0.1:4000)->CGNAT3(4.4.4.4:14000)->CGNAT2(2.2.2.2:12000)->NAT2(192.168.0.1:2000)[持续发送心跳包保活]
##### NAT2(192.168.0.1:2000)->CGNAT2(2.2.2.2:12000)->CGNAT4(4.4.4.4:14000)->NAT3(192.168.0.1:4000)[持续发送心跳包保活]