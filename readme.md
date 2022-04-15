### Monkey language

`Monkey`程序就是由一系列`statements（语句）`组成的。

#### let statements

格式：`let <identifier> = <expression>;`。

```monkey
let x = 5;
let y = 10;
let foobar = add(5, 5);
let barfoo = 5 * 5 / 10 + 18 - add(5, 5) + multiply(124);
let anotherName = barfoo;

let add = fn(a, b) {
    return a + b;
};
```

#### return statements

格式：`return <expression>;`。

```monkey
return 5;
return 10;
return add(15);
```

#### expressions

```monkey
// prefix operators
-5
!true
!false

// infix operators
5 + 5
5 - 5
5 * 5
5 / 5

// basic arithmetic operators：comparison operators
foo == bar
foo != bar
foo < bar
foo > bar

// group expressions
5 * (5 + 5)
((5 + 5) * 5) * 5

// call expressions
add(2, 3)
add(add(2, 3), add(5, 10))
max(5, add(5, (5 * 5)))

// identifiers expressions
foo * bar / foobar
add(foo, bar)

// function literal expressions
fn(x, y) { return x + y }(5, 5)
(fn(x) { return x }(5) + 10) * 10

// if expressions
if (10 > 5) { true } else { false }
```