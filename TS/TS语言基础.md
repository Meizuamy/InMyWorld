## 1. Class

* 面向对象三大特征：封装、继承、多态

* 类

* 对象

```ts
class Animal {
    constructor(name){
        this.name = name
    }

    run(){
        return `${this.name} is running`
    }
}

const snake = new Animal('lily')
console.log(snake.run())

```

### 1.1 修饰符

* Public 修饰的属性或方法是共有的

* Private 修饰的属性或方法是私有的

* Protected 修饰的属性或方法是受保护的


```ts
class Animal {
    constructor(name){
        this.name = name
    }

    // 私有方法
    private run(){
        return `${this.name} is running`
    }
}

const snake = new Animal('lily')
console.log(snake.run())

```

## 2. 接口

TypeScript 使用 `interface` 关键字声明接口。

```ts
// 声明接口
interface Radio {
    switchRadio(trigger: boolean): void;
}

```