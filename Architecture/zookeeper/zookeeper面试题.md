## zookeeper 用来做什么的，有几种类型的节点

### zookeeper 提供了什么
1. 文件系统
zookeeper 维护了一个类似文件系统的数据结构。每个子目录项都被成为znode,和文件系统类似，我们能够自由的增加、删除、znode,在znode下增加、删除子node,唯一的不同在于znode是可以存储数据的。