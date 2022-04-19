# Kafka架构

## 1. Kafka Producer


## 2. Kafka Broker


## 3. Kafka Consumer


## 4. Zookeeper and Raft

## 5. Topic and Partition

### 5.1 Topic
* 逻辑概念，同一个Topic的消息可分布在一个或多个节点（Broker）上
* 一个Topic包含一个或者多个Partition
* 每条消息都属于且仅属于一个topic
* Producer发布数据时，必须指定将该消息发送到哪一个Topic
* Consumer订阅消息时，也必须指定订阅哪个Topic的消息

### 5.2 Partition
* 物理概念，一个Partition只分布于一个Broker上（不考虑备份）
* 一个Partition物理上对应一个文件夹
* 一个Partition包含多个Segment（Segment对用户透明）
* 一个Segment对应一个文件
* Segment由一个个不可变记录组成的
* 记录只会被append到Segment中，不会被单独删除或修改。
* 清除过期日志时，直接删除一个或多个Segment