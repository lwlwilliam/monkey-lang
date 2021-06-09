### Monkey 语言的表达式

在 Monkey 语言中，除了 let 和 return 语句，其它的一切都是表达式。这些表达式以不同的形式出现。

> 前缀操作符表达式

```
-5
!true
!false
```

> 中缀操作符表达式(或者叫"二元操作符")

``` 
5 + 5
5 - 5
5 / 5
5 * 5
```

> 比较操作符表达式

```
foo == bar
foo != bar
foo < bar
foo > bar
```

> 小括号组成的表达式(影响计算顺序)

```
5 * (5 + 5)
((5 + 5) * 5) * 5
```

> 调用表达式

```
add(2, 3)
add(add(2, 3), add(5, 10))
max(5, add(5, (5 * 5)))
```

> 标识符也是表达式

```
foo * bar / foobar
add(foo, bar)
```

> 函数字面量表达式(在 Monkey 中，函数是第一公民)

```
let add = fn(x, y) { return x + y };
```

```
fn(x, y) { return x + y }(5, 5)
(fn(x) { return x }(5) + 10 ) * 10
```

> if 表达式(跟很多流行的语言不一样，Monkey 中的 if 是表达式)

```
let result = if (10 > 5) { true } else { false };
result // => true
```


### 术语

> 前缀操作符(prefix operator)

```
--5
```

> 后缀操作符(postfix operator)

Monkey 解释器没有 postfix operator。不是因为技术限制，而是纯粹篇幅所限。

```
foobar++
```

> 中缀表达式(infix operator)

infix operator 出在在二元表达式中，操作符有两个操作数。

```
5 * 8
```

> 操作符优先级(operator precedence, 也叫"运算次序"(order of operations))

```
5 + 5 * 10
```

以上表达式的结果是 55 而不是 100，因为`*`操作符的优先级更高，它比`+`操作符"更重要"。我有时会将操作符优先级想象成"操作符粘性"：操作符旁边的操作数对其粘性有多大。