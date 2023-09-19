# TraversalShare
## 介绍
打洞型点对点互传工具
## 实现目标
使用NAT穿透技术，通过STUN服务器检测NAT类型，筛选最优方案进行文字、文件等点对点互传通信
## 穿透方案
从上至下优先度递减
### IPv6
无防火墙IPv6-任意IPv6=直连
有防火墙IPv6-有防火墙IPv6=双方均进行防火墙打洞
### IPv4
NAT1-任意NAT=NAT1方打洞
NAT2-NAT3=双方均打洞
NAT2-NAT4=双方均打洞