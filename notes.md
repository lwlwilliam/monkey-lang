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

首先要定义`lexer`要输出的`tokens`。