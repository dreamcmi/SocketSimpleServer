# SocketSimpleServer

SocketSimpleServer是一个使用Golang编写的简单的Socket服务器，可以用于测试Socket通信的场景。它支持TCP和UDP协议，可以同时处理多个客户端的连接和数据传输。它还提供了一个config.toml文件，可以方便地调整服务器的IP,端口号和缓冲区大小。

## 安装

要安装SocketSimpleServer，你需要先安装Golang环境，可以参考[这里]的教程。然后，你可以使用以下命令来获取和编译项目：

```shell
git clone https://github.com/dreamcmi/SocketSimpleServer.git
cd SocketSimpleServer
go build main.go -o SocketSimpleServer
```

这样，你就可以在当前目录下找到一个名为`SocketSimpleServer`的可执行文件。

## 使用

要运行SocketSimpleServer，你需要在项目的根目录下找到一个名为config.toml的文件，然后按照以下格式来填写你想要的参数：

```toml
[tcp]
ip = "0.0.0.0" # ip
port = 23333   # 端口号
maxSize = 1000 # 接收最大包长

[udp]
ip = "0.0.0.0" # ip
port = 23334   # 端口号
maxSize = 1400 # 接收最大包长

```

然后，你就可以直接执行`./SocketSimpleServer`命令，它会自动读取config.toml文件中的参数。如果你想修改参数，只需要修改config.toml文件，然后重新启动服务器即可。

当服务器启动后，它会显示一些基本信息，如下所示：

```shell
2023-08-05 14:21:14 [ INFO  ] Welcome to SocketSimpleServer.
2023-08-05 14:21:14 [ INFO  ] Copyright 2023 dreamcmi. All rights reserved. 
2023-08-05 14:21:14 [ INFO  ] TCP Run: IP(0.0.0.0) PORT(23333) MAXSIZE(1000)
2023-08-05 14:21:14 [ INFO  ] UDP Run: IP(0.0.0.0) PORT(23334) MAXSIZE(1400)
```

然后，你就可以使用任何支持Socket通信的客户端来连接服务器，并发送和接收数据。服务器会在控制台上显示每个客户端的连接状态和数据内容，如下所示：

```shell
2023-08-05 14:27:19 [ INFO  ] TCP client connect(127.0.0.1:4294)
2023-08-05 14:27:20 [ INFO  ] TCP(127.0.0.1:4294) Receive(12):hello socket
2023-08-05 14:27:20 [ INFO  ] TCP(127.0.0.1:4294) Receive Hex:68656c6c6f20736f636b6574
2023-08-05 14:27:22 [ INFO  ] TCP client disconnect(127.0.0.1:4294)
```

## 开源协议

SocketSimpleServer是基于[Apache2.0]协议开源的。如果你觉得这个项目对你有帮助，请给它一个星星🌟，或者给我提出一些改进意见和建议。
