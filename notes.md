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

#### REPL

`Monkey`语言需要一个`REPL`。`REPL`表示"Read Eval Print Loop"。`Python`、`Ruby`以及每个`JavaScript`运行时都有`REPL`。有时候`REPL`又叫`console`或`interactive mode`。它们的概念是一样的：`REPL`读取输入，将其发送给`interpreter`进行`evaluation`，将结果/输出打印出来，然后重复以上步骤。

### Parsing

#### Parsers

`parser`作为一个软件组件，接收输入数据（通常是文本）并构建数据结构——常常是一些解析树、抽象语法树或者其它层级结构，使用这些结构来表示输入内容并检查语法。`parser`通常会有一个前置的独立词法分析器用于从输入字符序列中创建`token`。

这看起来有点抽象，以下用示例来说明。这是`JavaScript`代码：

```bash
> var input = '{"name": "Thorsten", "age": 28}';
> var output = JSON.parse(input)
> output
{ name: 'Thorsten', age: 28 }
> output.name
'Thorsten'
> output.age
28
>
```

以上的输入就是一些文本字符串。我们将其传到隐藏在`JSON.parse`函数背后的`parser`中就可以得到一个输出。输出的数据结构表示了输入：一个有两个字段`name`和`age`的JavaScript`对象。

`JSON`的`parser`跟编程语言的`parser`并没有什么不同。当然，`JSON`所要解析的数据可以一眼就可以看出应该使用什么数据结构来表示，而编程语言的输入文本则要复杂得多。

在大部分`interpreters`以及`compilers`中，用于表示源码的数据结构称为`syntax tree`或者`abstract syntax tree(AST)`。之所以是`abstract`，是因为源码中的细节在`AST`中会被省略。像`;`、`\n`、`\r`、` `、`(`、`)`、`{`、`}`和注释等可能不会在`AST`中表示出来。

并不存在一种通用的、正确的`AST`格式。但是所有`parser`的`AST`实现都十分类似，概念也是类似的，只是细节不一样。`AST`的具体实现取决于语言本向。

我们有以下`JavaScript`源码：

```js
if (3 * 5 > 10) {
    return "hello";
} else {
    return "goodbye";
}
```

假设我们有`MagicLexer`和`MagicParser`，构建出`JavaScript`对象形式的`AST`，那么在解析时可能会生成如下内容：

```bash
> var input = 'if (3 * 5 > 10) { return "hello"; } else { return "goodbye"; }';
> var tokens = MagicLexer.parse(input);
> MagicParser.parse(tokens);
{
  type: "if-statement",
  condition: {
    type: "operator-expression",
    operator: ">",
    left: {
      type: "operator-expression",
      operator: "*",
      left: { type: "integer-literal", value: 3 },
      right: { type: "integer-literal", value: 5 }
    },
    right: { type: "integer-literal", value: 10 }
  },
  consequence: {
    type: "return-statement",
    returnValue: { type: "string-literal", value: "hello" }
  },
  alternative: {
    type: "return-statement",
    returnValue: { type: "string-literal", value: "goodbye" }
  }
}
```

虽然`parser`输出的`AST`很抽象，但确实精确地表示了源码。这就是`parsers`干的活。它们以源码或作为输入（文本或者`tokens`形式），输出可以表示源码的数据结构。在构建数据结构时，会不可避免地分析输入数据，检查它们是否符合期望的结构。因此，解析的过程也叫`syntactic analysis（语法分析）`。

#### Why not a parser generator?

`parser`生成器是将语言的规范描述作为输入，然后生成`parser`的工具。该生成的`parser`又可以将源码作为输入来生成语法树。大部分生成器会使用`context-free grammar(CFG)`作为它们的输入。`CFG`是一套用于描述如何生成正确的（根据语法来判断）语言语句的规则。最常用的`CFG`标记格式是`Backus-Naur Form(BNF)`或者`Extended Backus-Naur Form(EBNF)`。

学习写自己的`parser`并不是在浪费时间，这其实是极为可贵的。只有在写过自己的`parser`之后，或者至少尝试过，才会发现生成器的优缺点，以及了解它们解决了哪方面的问题。

想要理解`parser`是如何工作的，最好的办法就是自己动手写一个，这也是十分有趣的。

#### Writing a parser for the Monkey programming language

解析一门编程语言有两种主要的策略：自顶向下解析和自底向上解析。`recursive descent parsing（递归下降解析）`，`Early parsing`或`predictive parsing`都是自顶向下解析的变体。

#### Parser's first steps: parsing let statements

```
let <identifier> = <expression>;
```

`statements`和`expressions`的区别：`expression`会产生值，`statements`不会。如下代码：`let x = 10;`不会产生值，而`5`会；`return 5;`说一句不会产生值，而`add(5, 5)`会。

```monkey
let x = 10;
let y = 15;
let foobar = add(5, 5);

let add = fn(a, b) {
    return a + b;
};
```

其实`expression`或`statement`是什么，哪个会产生值哪个不会，取决于编程语言本身。在某些语言中，函数字面量如`fn(x, y) { return x + y; }`是表达式，可以在其它任何可以使用表达式的地方使用。在某些语言中，函数字面量只能作为程序顶级函数声明语句的一部分。在某些语言中有`if expression`，其中的`conditionals`是`expressions`，会产生值。

`parser`的伪代码：

```
function parseProgram() {
    program = newProgramASTNode()
    
    advanceTokens()
    
    for (currentToken() != EOF_TOKEN) {
        statement = null
        
        if (currentToken() == LET_TOKEN) {
            statement = parseLetStatement()
        } else if (currentToken() == RETURN_TOKEN) {
            statement = parseReturnStatement()
        } else if (currentToken() == IF_TOKEN) {
            statement = parseIfStatement()
        }
        
        if (statement != null {
            program.Statments.push(statement) 
        }
        
        advanceTokens()
    }
    
    return program
}

function parseLetStatement() {
    advanceTokens()
    identifier = parseIdentifier()
    advanceTokens()
    
    if currentToken() != EQUAL_TOKEN {
        parseError("no equal sign!")
        return null
    }
    
    advanceTokens()
    value = parseExpression()
    variableStatement = newVariableStatementASTNode()
    variableStatement.identifier = identifier
    variableStatement.value = value
    return variableStatement
}

function parseIdentifier() {
    identifier = newIdentifierASTNode()
    identifier.token = currentToken()
    return identifier
}

function parseExpression() {
    if (currentToken() == INTEGER_TOKEN) {
        if (nextToken() == PLUS_TOKEN) {
            return parseOperatorExpression()
        } else if (nextToken() == SEMICOLON_TOKEN) {
            return parseIntegerLiteral()
        }
    } else if (currentToken() == LEFT_PAREN) {
        return parseGroupedExpression()
    }
    // [...]
}

function parseOperatorExpression() {
    operatorExpression = newOperatorExpression()
    operatorExpression.left = parseIntegerLiteral()
    operatorExpression.operator = currentToken()
    operatorExpression.right = parseExpression()
    return operatorExpression()
}

// [...]
```

递归下降解析的基本思想就如上述伪代码。入口就是`parseProgram`，初始化`AST`的根节点(`newProgramASTNode()`)。接着就可以基于当前`token`构建子结点——`statements`了，如此递归。

解析`expression`是写`parser`中最有趣的部分了。解析`statements`则相当直接，从左到右处理`tokens`，预测下一个`token`或者拒绝它，如果一切都符合就返回一个`AST`节点。

解析`expression`则会有很多挑战，操作符优先级可能就是其中最先被想到的。假设我们想要解析如下算术表达式：

```
5 * 5 + 10
```

实际我们想要一个能表达如下表达式的`AST`：

```
((5 * 5) + 10)
```

也就是说，`5 * 5`需要在`AST`的更深处，比加法更早计算。为了处理像这种`AST`，`parser`要知道操作符`*`的优先级比`+`高。再考虑如下表达式：

```
5 * (5 + 10)
```

这里的小括号将`5 + 10`组合在一起了，现在加法要比乘法先计算。因为小括号的优先级比`*`更高。

另一个大的挑战就是`expressions`的`tokens`可以出现在多处。与此相反，`let`的只会在`let`语句的开始位置出现一次。下面看看这个表达式：

```
-5 - 10
```

这里在表达式开始位置的`-`操作符是作为`prefix`操作符，在中间位置的则作为`infix`操作符。再看看下面出现的该挑战的变体：

```
5 * (add(2, 3) + 10)
```

外面的小括号是作为组表达式存在的，里面的小括号是一个`call expression`。`token`的正确性取决于上下文——`tokens`是出现在前面、后面以及它们的优先级。

