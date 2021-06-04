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
```