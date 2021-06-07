### 递归下降 parser 伪代码

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
        
        if (statement != null) {
            program.Statements.push(statement) 
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

虽然这段伪代码有很多省略，但是里面包含了递归下降解析的基本思想。`parseProgram`是入口，初始化了 AST 的根节点（`newProgramASTNode()`），然后开始通过调用其它可以根据当前 token 来判断应该创建什么 AST 节点的函数来创建子节点，也就是语句。其它函数又会递归地互相调用。

使用"Pratt parsing"。