### monkey 语言

```
// 整型、布尔型和字符串
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);

// 数组、哈希
let myArray = [1, 2, 3, 4, 5];
let thorsten = {"name": "Thorsten", "age": 28};

// 通过索引表达式访问数组和哈希
myArray[0]       // => 1
thorsten["name"] // => "Thorsten"

// 变量绑定函数
let add = fn(a, b) { return a + b; };

// 隐式返回值
let add = fn(a, b) { a + b; };

// 调用函数
add(1, 2);

// 复杂点的函数
let fibonacci = fn(x) {
    if (x == 0) {
        0
    } else {
        if (x == 1) {
            1 
        } else {
            fibonacci(x - 1) + fibonacci(x - 2); 
        }
    }
};

// 高阶函数：将其它函数作为该函数的参数
let twice = fn(f, x) {
    return f(f(x));
};

let addTwo = fn(x) {
    return x + 2;
};

twice(addTwo, 2); // => 6


```