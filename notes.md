### Lexing

在源码被`evaluate`之前，需要通过两个步骤改变其表示形式，如下：

```
Source Code --> Tokens --> Abstract Syntax Tree
```

第一步从`Source Code`到`Tokens`，被称为`lexical analysis`，或者简称`lexing`。这一步由`lexer`来处理（也叫`tokenizer`或`scanner`）。

`tokens`本身是提供给`parser`易于分类的小型数据结构，`parser`则进行第二步处理，将`tokens`转换为`Abstract Syntax Tree`。

给`lexer`输入代码`"let x = 5 + 5;"`，输出结果形式大概如下：

```
[
    LET,
    IDENTIFIER("x"),
    EQUAL_SIGN,
    INTEGER(5),
    PLUS_SIGN,
    INTEGER(5),
    SEMICOLON
]
```

在`Monkey`语言中，空白符会被忽略，因为它们在语言中没有意义。但在其它语言中，如`Python`中，空白符是有意义的，这意味着需要将空白符当作`tokens`输出到`parser`中。

生产级别的`lexer`可能会加上`tokens`的行号、列号以及文件名，以输出更多有用的信息。

#### Defining Our Tokens

首先要定义`lexer`要输出的`tokens`。先从一些`token`的定义开始，然后再扩展`lexer`。我们将要`lex`的`Monkey`语言的子集看起来如下所示：

```
let five = 5;
let ten = 10;

let add = fn(x, y) {
    x + y;
};

let result = add(five, ten);
```

上例包含了几种`tokens`：像`5`和`10`这样的数字；变量名`x`、`y`、`add`和`result`；不是数字，也不是变量名，属于语言一部分的单词(`keyword`)`let`和`fn`；当然还有很多特别的字符`(`、`)`
、`{`、`}`、`=`、`,`和`;`。

正如上面所说的。我们定义的`Token`数据结构中需要一个`type`属性，还需要一个字段保存`token`对应的值。

```go
package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
```

#### The Lexer

`lexer`以源码作为输入，并输出可以表示源码的`tokens`。它不需要缓冲或保存`tokens`。

也就是说，我们只需要用源码初始化`lexer`然后不断调用`NextToken()`获取下一个`token`即可。为了简化，可以使用字符串表示代码。

`lexer`的工作不是告诉我们代码是否有意义或者是否工作以及包含错误，它的工作应该仅仅是将输入转换为`tokens`。

##### REPL

`Monkey`语言需要一个`REPL`。`REPL`表示"Read Eval Print Loop"。`Python`、`Ruby`以及每个`JavaScript`运行时都有`REPL`。有时候`REPL`又叫`console`或`interactive mode`。它们的概念是一样的：`REPL`读取输入，将其发送给`interpreter`进行`evaluation`，将结果/输出打印出来，然后重复以上步骤。
