# Linear_Inverse_Modeling


## 1. 什么是逆线性模型

> 逆线性模型(linear inverse model, LIM)是根据物种之间的质量平衡关系, 依托营养动力学原理构建的生态系统能量流动模型。逆线性问题最早出现在地球物理科学中, 直到20 世纪80 年代末, 才应用于食物网能量流动分析。

### 1.1 什么是线性模型

我们通常会使用线性模型来进行预测未知事件，例如通过统计发病人的一些年龄，身体因素来关联癌症的发病率等，在二维下呈现一条直线，所以称为线性。

1）线性回归
https://zhuanlan.zhihu.com/p/72513104


## 2. 逆线性模型有什么用途

### 2.1 在海洋生态的研究
[基于LIM-MCMC模型研究海州湾食物网能量流动特征.pdf](./基于LIM-MCMC模型研究海州湾食物网能量流动特征.pdf)


## 3. 举个栗子

### 3.1 LIM生态系统应用

假设我们有一个简单的生态系统，其中包含三种物种：A、B和C。我们知道这些物种之间存在捕食关系，即物种A捕食物种B，物种B捕食物种C。我们的目标是使用LIM来估计这些捕食关系的强度。

首先，我们需要收集关于这些物种数量的数据。假设我们已经有了一段时间内物种A、B和C的种群数量记录。

接下来，我们构建LIM的线性方程组。在这个例子中，LIM方程可以表示为：

A_t = a * A_{t-1} - b * B_t + ...
B_t = c * B_{t-1} - d * A_t - e * C_t + ...
C_t = f * C_{t-1} - g * B_t + ...

其中，A_t、B_t和C_t分别表示物种A、B和C在时刻t的种群数量，a、b、c、d、e、f和g是待求解的参数，它们代表了物种间相互作用的强度。

然后，我们使用数学方法（如最小二乘法）来求解这个方程组，得到参数a、b、c、d、e、f和g的估计值。这些估计值就代表了物种间捕食关系的强度。

最后，我们可以利用这些估计值来预测生态系统在未来一段时间内的动态变化。例如，我们可以预测物种A、B和C在未来某个时刻的种群数量，以及它们之间的相互作用如何影响这些数量的变化。

需要注意的是，这个例子是一个简化的LIM应用，实际的生态系统往往更加复杂，包含更多的物种和相互作用关系。因此，在实际应用中，LIM的构建和求解过程可能会更加复杂，需要借助专业的数学和统计工具来完成。

此外，LIM的应用并不仅限于生态系统中物种数量的预测，它还可以用于研究生态系统的稳定性、物种多样性的变化等问题。通过分析和预测这些生态指标的变化，我们可以更好地理解生态系统的结构和功能，为生态保护和资源管理提供科学依据。

## 4. 总结


## 5. 参考

[【Github】美国国家气象站-LIM](https://github.com/NOAA-PSL/Linear_Inverse_Modeling)