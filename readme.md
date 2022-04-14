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
