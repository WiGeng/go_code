package main

/* Golang的主要设计目标之一就是面向大规模后端服务程序，网络通信这块是服务端程序必不可少也是至关重要的一部分。
网络编程有两种：
	1) TCP socket编程，是网络编程的主流。之所以叫Tcp socket编程，是因为底层是基于cp/ip协议的．比如：QQ聊天
	2）b/s结构的http编程，我们使用浏览器去访问服务器时，使用的就是http协议，而http底层依旧是用tcp socket实现的。比如：京东商城
*/