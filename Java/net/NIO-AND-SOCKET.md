# NIO

## **1.** SockerChannel

SocketChannel 是一个连接到TCP网络套接字（Socket）的通道。

### **1.1** 创建SocketChannel的两种方式
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


### **1.2** 从SocketChannel中读取数据

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

### **1.3** 写入SocketChannel
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

### **1.4** 非阻塞模式

> 可以设置SocketChannel为非阻塞模式（non-blocking mode），设置之后，就可以在异步模式下调用connect(),write(),read()

```java

SocketChannel socketChannel = SocketChannel.open();

socketChannel.configureBlocking(false);
socketChannel.connect(new InetSocketAddress("127.0.0.1",8080));

```
#### 1.4.1 connect()

> 如果SocketChannel在非阻塞模式下，此时调用connect(), 该方法可能在链接建立之前就返回了。为了确定是否建立了链接，可以调用finishConnect()的方法来确定。

```java
SocketChannel socketChannel = SocketChannel.open();
socketChannel.configureBlocking(false);
socketChannel.connect(new InetSocketAddress("127.0.0.1",8080));

while(!socketChannel.finishConnect()){
    // wait, or do something else..
}
```
#### 1.4.2 write()

> 非阻塞模式下，write()方法在尚未写出任何内容时可能就返回了。所以需要在循环中调用write()。

```java
String data = "Hello World!";
SocketChannel socketChannel = SocketChannel.open();
ByteBuffer buf = ByteBuffer.allocate(100);
buf.clear();
buf.put(data.getBytes());
buf.flip();

while(buf.hasRemaining()){
    socketChannel.write(buf);
}

```

#### 1.4.3 read()

> 非阻塞模式下，read()方法在尚未读取到任何数据时可能就返回了。所以要关注它的int返回值，它会告诉你当前读取了多少字节。


总结：非阻塞模式与选择器搭配会工作的更好，通过将一个或多个SocketChannel注册到Selector，可以询问Selector哪个通道准备好了读取，写入等。