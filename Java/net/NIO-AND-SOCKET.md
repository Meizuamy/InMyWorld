## NIO

### **1.** SockerChannel

SocketChannel 是一个连接到TCP网络套接字（Socket）的通道。

#### **1.1** 创建SocketChannel的两种方式
1. 打开一个SocketChannel并连接到互联网上的某台服务器。

> 打开SocketChannel

```java

SocketChannel socketChannel = SocketChannel.open();
socketChannel.connect(new InetSocketAddress("127.0.0.1",8080));

```

> 关闭SocketChannel

```java
socketChannel.close();
```

2. 一个新连接到达ServerSocketChannel时，回创建一个SocketChannel。

```java
ServerSocketChannel serverSocketChannel = ServerSocketChannel.open();
serverSocketChannel.socket().bind(new InetSocketAddress(8080));

while(true){
    SocketChannel socketChannel = serverSocketChannel.accept();

    // 处理socketChannel
}
```


#### **1.2** 从SocketChannel中读取数据

> 使用ByteBuffer申请内存，使用read()读取返回的数据。

```java
// 直接读取读取当前ByteBuffer空间大小的数据。
ByteBuffer buf = ByteBuffer.allocate(1024);
// 函数的返回值是读取了多少字节，如果返回了-1代表你已经读取完了所有的数据
int size = socketChannel.read(buf);

ByteBuffer[] buffers = new ByteBuffer[]{ buf };

// 读取指定长度的数据        
// ByteBuffer[] dsts - ByteBuffer数组
// int offset - 偏移地址，从当前字节开始读取
// int length - 读取长度
socketChannel.read(buf, 0, 1024);

//读取多个ByteBuffer空间大小的数据。
socketChannel.read(buffers);    
```

#### **1.3** 写入SocketChannel
> 使用write方法写入ByteBuffer
```java

SocketChannel socketChannel = SocketChannel.open();

String data = "Hello World!";
ByteBuffer buf = ByteBuffer.allocate(100);
buf.put(data.getBytes());

// flip重置当前buffer的标记位置（mark = -1)
// flip重置当前buffer的limit为当前buffer写入的字节数（limit=position)，并将当前写入指针position设置为0
buf.flip();

// hasRemaining方法是判断当前buffer的数据是否读取完毕。
while(buf.hasRemaining()){
    socketChannel.write(buf);
}

```

#### **1.4** 非阻塞模式
