# NIO


## **1.** Channel
> Java 的NIO通道和流类似，但又不同
> 
> * 既可以读取通道中的数据，又可以写数据到通道中，但流是单向的。
> * 通道可以异步的读写。
> * 通道中的数据总是要读取一个Buffer，或者从Buffer写入。

### **1.1** Channel的实现
* FileChannel - 从文件中读写数据
* DatagramChannel - 通过UDP读写网络中的数据
* SocketChannel - 通过TCP读写网络中的数据
* ServerSocketChannel - 监听新进来的TCP链接，对每一个新进来的链接，都会创建一个SocketChannel


## **2.** SockerChannel

SocketChannel 是一个连接到TCP网络套接字（Socket）的通道。

### **2.1** 创建SocketChannel的两种方式
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


### **2.2** 从SocketChannel中读取数据

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

### **2.3** 写入SocketChannel
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

### **2.4** 非阻塞模式

> 可以设置SocketChannel为非阻塞模式（non-blocking mode），设置之后，就可以在异步模式下调用connect(),write(),read()

```java

SocketChannel socketChannel = SocketChannel.open();

socketChannel.configureBlocking(false);
socketChannel.connect(new InetSocketAddress("127.0.0.1",8080));

```
#### 2.4.1 connect()

> 如果SocketChannel在非阻塞模式下，此时调用connect(), 该方法可能在链接建立之前就返回了。为了确定是否建立了链接，可以调用finishConnect()的方法来确定。

```java
SocketChannel socketChannel = SocketChannel.open();
socketChannel.configureBlocking(false);
socketChannel.connect(new InetSocketAddress("127.0.0.1",8080));

while(!socketChannel.finishConnect()){
    // wait, or do something else..
}
```
#### 2.4.2 write()

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

#### 2.4.3 read()

> 非阻塞模式下，read()方法在尚未读取到任何数据时可能就返回了。所以要关注它的int返回值，它会告诉你当前读取了多少字节。


总结：非阻塞模式与选择器搭配会工作的更好，通过将一个或多个SocketChannel注册到Selector，可以询问Selector哪个通道准备好了读取，写入等。


## **3.** ServerSocketChannel




## **4.** FileChannel

> 基本的FileChannel实例

```java
RandomAccessFile file = new RandomAccessFile("data/nio-data.txt","rw");
FileChannel inChannel = file.getChannel();
ByteBuffer buf = ByteBuffer.allocate(48);

while(inChannel.read(buf) != -1){
    System.out.println("Read" + byteRead);
    buf.flip();

    while(buf.hasRemaining()){
        System.out.print((char) buf.get());
    }

    buf.clear();
}

file.close();

```

## **1.5** DatagramChannel
> Java NIO 中的DatagramChannel 是一个能收发UDP包的通道。因为UDP是无连接的网络协议，所以不能像其他通道那样读取和写入。它发送的是数据包。

### **1.5.1** 打开DatagramChannel
> 下面是一个打开DatagramChannel的简单例子

```java
DatagramChannel channel = DatagramChannel.open();
channel.socket().bind(new InetSocketAddress(9999));
```

### **1.5.2** 接收数据

> 可以通过 receive() 方法接收数据

```java
DatagramChannel channel = DatagramChannel.open();
channel.socket().bind(new InetSocketAddress(9999));
ByteBuffer buf = ByteBuffer.allocate(48);
buf.clear();
channel.receive(buf);
```
### **1.5.2** 发送数据

> 通过send()方法从DatagramChannel发送数据。

```java
DatagramChannel channel = DatagramChannel.open();
String data = "Hello World!";
ByteBuffer buf = ByteBuffer.allocate(48);
buf.clear();
buf.put(data.getBytes());
buf.flip();
channel.send(buf, new InetSocketAddress("127.0.0.1",9999));
```

### **1.5.3** 链接到特定的地址
> 可以将DatagramChannel "链接" 到网络的特定地址。由于UDP是无连接的，链接到特定地址并不会像TCP那样创建一个真正的链接。而是锁住DatagramChannel，让其只能从特定的地址收发数据。

```java
DatagramChannel channel = DatagramChannel.open();
channel.connect(new InetSocketAddress("127.0.0.1",9999));
ByteBuffer buf = ByteBuffer.allocate(48);
int bytesRead = channel.read(buf);
int bytesWritten = channel.write(buf);
```

## **2** Selector

> Selector 可以检测一到多个NIO通道，并能够知晓通道是否为读写时间准备好的组件。这样，一个单独的西安城可以管理多个Channel，从而管理多个网络连接。

### **2.1** 为什么使用Selctor？


### **2.2** 创建Selctor

```java
Selector selector = Selector.open();
```

### **2.3** 向Selector注册通道

> 可以通过SelectableChannel.register()方法来进行注册到selector。SelectableChannel是一个抽象类，所有的通道都实现自这个类，所以所有的通道都可以使用此方法。

```java
SocketChannel channel = SocketChannel.open();
Selector selector = Selector.open();
channel.configureBlocking(false);
SelectionKey key = channel.register(selector, SelectionKey.OP_READ);

```
> 向Selector注册到channel时，必须将channel置于非阻塞模式下。由于FileChannel不能设置非阻塞模式，所以不能和Selector一起使用。
> register()方法的第二个参数，是一个”interest集合“，意思是在通过Selector监听Channel时对什么时间感兴趣。可以监听四种事件：
1. Connect
2. Accept
3. Read
4. Write
通道触发了一个事件意思是该事件已经就绪。某个Channel成功连接到另外一个服务器称为”连接就绪“。一个ServerSocketChannel准备好接收新的进入的连接成为”接收就绪“。一个可读的通道可以说是”读就绪”。等待写数据的通道可以说是“写就绪”。

这四种事件用SelectionKey的四个常量来表示：
1. SelectionKey.OP_CONNECT
2. SelectionKey.OP_ACCEPT
3. SelectionKey.OP_READ
4. SelectionKey.OP_WRITE

如果不止对一种事件感兴趣，可以用“位或”操作符将常量连接起来：
```java
int interestSet = SelectionKey.OP_READ | SelectionKey.OP_WRITE;
```

### **2.4** SelectionKey
向Selector注册Channel的时候，register() 方法会返回一个 SelectionKey 对象。这个对象包含一些属性：
* interest集合
* ready集合
* Channel
* Selctor
* 附加的对象（可选）

#### **2.4.1** interest集合
interest集合是你所选择的感兴趣的事件集合。可以通过SelctionKey读写interest集合。
```java 
int interestSet = selectionKey.interestOps();

boolean isInterestedInAccept = (interestSet & SelectionKey.OP_ACCEPT) == SelectionKey.OP_ACCEPT;
boolean isInterestedInConnect = interestSet & SelectionKey.OP_CONNECT;
boolean isInterestedInRead = interestSet & SelectionKey.OP_READ;
boolean isInterestedInWrite = interestSet & SelectionKey.OP_WRITE;
```

#### **2.4.2** ready集合
ready集合是通道已经准备就绪的操作集合。在一次选择（selection)之后，你会首先访问这个readySet。

```java

int readySet = selectionKey.readyOps();
```

可以用类似检测interest集合的方法来检测channel中什么事件或者操作已经准备就绪。也可以使用以下方法来检测，他们都会返回一个bool类型。

```java
selectionKey.isAcceptable();
selectionKey.isConnectable();
selectionKey.isReadable();
selectionKey.isWriteable();
```

#### **2.4.5** 使用SelectionKey 访问 Channel 和 Selector
从SelectionKey 访问 Channel 和 Selector：
```java
Channel channel = selectionKey.channel();
Selector selector = selectionKey.selector();
```


### **2.5** 通过Selector选择通道
一旦向Selector注册了一个或多个通道，就可以调用几个重载的select()方法。这些方法返回你所感兴趣的事件（如连接，接受，读和写）已经准备就绪的那些通道。例如：如果你对“读就绪”的通道感兴趣，select()方法会返回读事件已经就绪的那些通道。
下面是select()方法和它的一些重载：
```java
int select();
int select(long timeout);
int selectNow();
```
select() 阻塞到至少有一个通道在你注册的事件上就绪了。
select(long timeout) 和select()一样，除了最长会阻塞timeout毫秒（参数）;
selectNow() 不会阻塞，不管什么通道就绪都立即返回。

select() 方法返回的int值表示有多少通道已经就绪。即上次调用select()方法后有多少通道变成就绪状态。如果调用select()方法，因为有一个通道变成就绪状态，所以返回1，若再次调用select()方法，如果零一个通道就绪了，它会再次返回1。如果对第一个就绪的channel没有做任何操作，现在就有两个就绪的通道，但在每次select()方法调用之前，只有一个通道就绪了。

#### **2.5.1** slectedKeys()
一旦调用了select()方法，并且返回值表明有一个或更多个通道就绪了，然后可以通过调用selector的selectedKeys()方法，访问“已选择键集（selected key set）”中的就绪通道。

```java
Set selectedkeys = selector.selectedKeys();
```

当向Selector注册Channel时，Channel.register()方法会返回一个SelectionKey对象。这个对象代表了注册到该Selector的通道。可以通过SelectionKey的selectedKeySet()方法访问这些对象。
可以遍历这个已选择的键集合来访问就绪的通道：

```java
Set selectedKeys = selector.selectedKeys();
Iterator keyIterator = selectedKeys.iterator();
while(keyIterator.hasNext()){
    SelectionKey key = keyIterator.next();
    if(key.isAcceptable()){
        // a connection was accepted by a ServerSocketChannel.
    } else if (key.isConnectable()){
        // a connection was established with a remote server.
    } else if (key.isReadable()){
        // a channel is ready for reading.
    } else if (key.isWritable()){
        // a channel is ready for writing.
    }
    keyIterator.remove();
}
```
这个循环遍历已选择键集合中的每个键，并检测各个键所对应的通道的就绪事件。

注意每次迭代末尾的keyIterator.remove()调用。Selector不会自己从已选择键集中移除SelectionKey实例。必须在处理玩通道时自己移出。下次该通道编程就绪时，Selector会再次将其放入已选择键集中。

SelctionKey.channel()方法返回的通道需要转型成你要处理的类型，如ServerSocketChannel或SocketChannel等。

#### **2.5.2** wakeUp()
某个线程调用select()方法后阻塞了，即使没有通道已经就绪，也有办法将其从select()方法返回。只要让其它线程在第一个线程调用的select()方法的哪个对象上调用Selector.wakeup()方法即可。阻塞在select() 方法上的线程会立马返回。

如果有其他线程调用wakeup()方法，但当前没有线程阻塞在select()方法上，下个调用select()方法的线程会立即”wake up“

#### **2.5.3** close()

用完Selector后调用其close()方法会关闭该Selector，并使注册到该Selector上的所有SelectionKey实例无效。通道本身并不会关闭。

### **** 完整的示例

```java

Selector selector = Selector.open();

channel.configureBlocking(false);

SelectionKey key = channel.register(channel,SelectionKey.OP_READ);

while(true){
    int readyChannels = selector.select();
    int (readyChannels = 0) continue;
    Set selectedKeys = selector.selectedKeys();
    Iterator iterator = selectedKeys.iterator();
    while(iterator.hasNext()){
        SelectionKey key = iterator.next();
        if(key.isAcceptable()){

        } else if (key.isConnectable()){

        } else if (key.isReadable()){

        } else if (key.isWritable()){

        }
        iterator.remove();
    }
}
```

### Pipe

Java NIO 管道是一个线程之间的单向数据连接。Pipe有一个source通道和一个sink通道。数据会被写到sink通道，从source通道读取。

#### 创建管道
通过Pipe.open()方法打开管道。例如：
```java
Pipe pipe = Pipe.open();
```

#### 向管道写数据
要向管道写数据，需要访问sink通道。像这样：
```java
Pipe.SinkChannel sinkCHannel = pipe.sink();
```

通过调用SinkChannel的write()方法，将数据写入SinkChannel，像这样:
```java
String newData = "New String to write to file..." + System.currentTimeMillis();
ByteBuffer buf = ByteBuffer.allocate(48);
buf.clear();
buf.put(newData.getBytes());

buf.flip();

while(buf.hasRemaining()){
    sinkChannel.write(buf);
}
```

#### 从管道读取数据

通过source通道读取管道中的数据
```java

Pipe.SourceChannel sourceChannel = pipe.source();

```

调用source通道的read()方法来读取数据
```java
ByteBuffer buf = ByteBuffer.allocate(48);
int bytesRead = sourceChannel.read(buf);
```
read()方法返回的int值会告诉我们多少字节被读进了缓冲区。


### Path

Java的Path接口时Java NIO2的一部分，是对Java6和Java7的NIO的更新。Java的Path接口在Java7中被添加到Java NIO，位于java.nio.file包中，其全路径是java.nio.file.Path。

一个Path实例代表一个文件系统中的路径。一个路径可以指向一个文件或者一个文件夹。一个路径可以是绝对路径或相对路径。绝对路径是从根路径开始的全路径，相对路径是一个相对于其他路径的文件或者文件夹路径。

在很多地方java.nio.file.Path接口和java.io.File类是相似的，但是他们还是有很多不同的。在很多类中，你可以使用很多Path接口替代File类。

#### 创建一个Path实例
想要使用一个Path实例，你必须要先创建一个Path实例，可以使用Paths类中的静态方法Paths.get()创建。注意，使用get()方法是默认会解析`.`和`..`
```java
import java.nio.file.Path;
import java.nio.file.Paths;

public class FileExample {

    public static void main(String[] args){
        Path path = Paths.get("c:\\data\\myfile.txt");
    }
    
}

```

通常创建相对路径的Path我们可以使用Paths.get(basePath, relativePath)方法创建一个相对路径的实例。

#### Path.normalize()方法
Path接口中的normalize()可以标准化一个路径。标准化意思是解析路径中的`.`和`..`。
```java

String originalPath = "d:\\data\\projects\\a-project\\..\\another-project";

Path path = Paths.get(origianlPath);
System.out.println("path = " + path);

Path normalizePath = path.normalize();

System.out.println("normalizePath = " + normalizePath);

```

## 通道间的数据传输

在java nio中如果两个通道有一个是FileChannel，那你可以直接将数据从一个channel传输到另外一个channel。


### transferFrom()

FileChannel的transferFrom()方法可以将数据从原通道传输到FileChannel中，如下面的示例：
```java
RandomAccessFile fromFile = new RandomAccessFile("fromFile.txt","rw");
FileChannel fromChannel = fromFile.getChennel();

RandomAccessFile toFile = new RandomAccessFile("toFile.txt","rw");
FileChannel toChannel = toFile.getChannel();

long position = 0;
long count = fromChannel.size();

toChannel.transferFrom(position, count, fromChannel);
```

方法的参数：

position，表示从position处开始向目标文件写入数据。
count，表示最多传输的字节数。如果源通道的剩余空间小于count个字节，则所传输的字节数要小于请求的字节数。
fromChannel，字节源

> **注意**
> 在SocketChannel的实现中，SocketChannel只会传输此刻准备好的数据（可能不足count字节）。因此，SocketChannel可能不会将请求的所有数据（count个字节）全部传输到FileChannel中。


### transferTo()

transferTo方法将数据从一个FileChannel传输到其他的channel中。

一个简答的例子：

```java
RandomAccessFile fromFile = new RandomAccessFile("fromFile.txt", "rw");
FileChannel fromChannel = fromFile.getChannel();

RandomAccessFile toFile = new RandomAccessFile("toFile.txt", "rw");
FileChannel toChannel = toFile.getChannel();

long position = 0;
long count = fromChannel.size();

fromChannel.transferTo(position, count, toChannel);

```

如上所说的关于SocketChannel的问题，在transferTo()方法中同样存在。SocketChannel会一致传输到数据直到目标buffer被填满。