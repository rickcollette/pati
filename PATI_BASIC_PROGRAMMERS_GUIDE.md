
# PATI BASIC Programmer's Guide

Welcome to the updated PATI BASIC Programmer's Guide. This guide provides all the information you need to write programs in PATI BASIC, based on the current implementation. PATI BASIC is a simplified version of the BASIC programming language, designed to be easy to learn and use.

## Table of Contents

- Introduction
- Variables and Data Types
- Expressions and Operators
- Statements
  - LET Statement
  - IF...THEN Statement
  - PRINT Statement
  - INPUT Statement
  - RETURN Statement
  - END Statement
- Procedures
  - Defining Procedures
  - Calling Procedures
- Operators and Special Symbols
- Reserved Words
- Sample Program
- Limitations and Known Issues
- Conclusion

## Introduction

PATI BASIC is a procedural programming language inspired by traditional BASIC. It supports fundamental programming constructs such as variables, arithmetic operations, conditional statements, input/output operations, and user-defined procedures.

## Variables and Data Types

### Variables

Variables in PATI BASIC are used to store integer values. A variable name must be a single uppercase letter from A to Z.

### Declaration and Initialization

Variables are declared and initialized using the LET statement. There is no need for a separate declaration; assigning a value to a variable declares it.

**Syntax:**

```basic
LET <variable> = <expression>
```

**Example:**

```basic
LET A = 5
LET B = A + 10
```

### Data Types

Currently, PATI BASIC supports only integer data types. All variables and expressions are evaluated as integers.

## Expressions and Operators

Expressions are used to perform calculations and evaluate conditions. PATI BASIC supports the following arithmetic operators:

- Addition: `+`
- Subtraction: `-`
- Multiplication: `*`
- Division: `/`

**Note:** Operator precedence is limited. It's recommended to use parentheses to ensure the correct order of operations.

**Example:**

```basic
LET C = (A + B) * 2
```

## Statements

### LET Statement

Assigns the result of an expression to a variable.

**Syntax:**

```basic
LET <variable> = <expression>
```

**Example:**

```basic
LET X = 10
LET Y = X * 2
```

### IF...THEN Statement

Executes a statement if a condition is true.

**Syntax:**

```basic
IF <expression> <relational_operator> <expression> THEN <statement>
```

**Relational Operators:**

- Equal to: `=`
- Not equal to: `!=`
- Less than: `<`
- Less than or equal to: `<=`
- Greater than: `>`
- Greater than or equal to: `>=`

**Example:**

```basic
IF X > 5 THEN PRINT "X is greater than 5"
```

### PRINT Statement

Outputs a string or variable value to the console.

**Syntax:**

```basic
PRINT <string_or_expression> [; <string_or_expression> ...]
```

**Example:**

```basic
PRINT "Hello, World!"
PRINT "The value of X is "; X
```

**Note:** Use semicolons `;` to concatenate multiple strings or expressions in a single PRINT statement.

### INPUT Statement

Prompts the user for input and assigns it to a variable.

**Syntax:**

```basic
INPUT <variable>
```

**Example:**

```basic
INPUT A
```

### RETURN Statement

Returns from a procedure to the calling point.

**Syntax:**

```basic
RETURN
```

**Note:** The RETURN statement should be used inside a procedure to return control to the calling code.

### END Statement

Ends the program execution.

**Syntax:**

```basic
END
```

## Procedures

Procedures are blocks of code that perform specific tasks and can be called from other parts of the program by using their name.

### Defining Procedures

Procedures are defined using the `PROC` keyword, followed by the procedure name and a block of code enclosed in `{` and `}`.

**Syntax:**

```basic
PROC <procedure_name> {
    <statements>
}
```

**Example:**

```basic
PROC Greet {
    PRINT "Hello from the procedure!"
    RETURN
}
```

### Calling Procedures

Procedures are called by simply writing their name as a statement in your code.

**Syntax:**

```basic
<procedure_name>
```

**Example:**

```basic
Greet
```

**Note:** There is no CALL keyword used to invoke procedures in the current implementation. Simply write the procedure name to call it.

**Parameters and Arguments:**

_Current Limitation_: Procedures do not support parameters or arguments in the current implementation.

- All variables are global, and procedures can access and modify them directly.

## Operators and Special Symbols

### Arithmetic Operators

- Addition: `+`
- Subtraction: `-`
- Multiplication: `*`
- Division: `/`

### Relational Operators

- Equal to: `=`
- Not equal to: `!=`
- Less than: `<`
- Less than or equal to: `<=`
- Greater than: `>`
- Greater than or equal to: `>=`

### Parentheses

- Left Parenthesis: `(`
- Right Parenthesis: `)`
Use parentheses to group expressions and control the order of evaluation.

### Braces

- Left Brace: `{`
- Right Brace: `}`
Used to enclose the body of a procedure.

### Semicolon

`;`
Used in PRINT statements to concatenate strings and expressions.

### Comma

`,`
Currently, commas are not used in the language syntax.

## Reserved Words

The following keywords are reserved in PATI BASIC and cannot be used as variable names:

- LET
- IF
- THEN
- PRINT
- INPUT
- PROC
- RETURN
- END

## Sample Program

Here's a sample PATI BASIC program demonstrating the use of variables, expressions, procedures, and control flow statements.

```basic
LET A = 5
LET B = 10

PRINT "Initial values:"
PRINT "A = "; A
PRINT "B = "; B

AddNumbers

END

PROC AddNumbers {
    LET C = A + B
    PRINT "The sum of A and B is "; C
    RETURN
}
```

**Explanation:**

- The program initializes two variables `A` and `B`.
- It prints their initial values.
- It calls the procedure `AddNumbers` by simply writing its name.
- The `AddNumbers` procedure calculates the sum of `A` and `B` and prints the result.
- The program ends after the `END` statement.

## Limitations and Known Issues

- **Expression Parsing**: The expression parser is limited and does not fully support operator precedence. Use parentheses to ensure expressions are evaluated correctly.
- **Variable Names**: Variables must be single uppercase letters from A to Z. There is no support for multi-character variable names.
- **Data Types**: Only integer variables are supported. There is no support for strings or other data types in variables.
- **Procedures**:
  - No Parameters: Procedures do not accept parameters or arguments.
  - Variable Scope: All variables are global. Procedures access and modify global variables.
  - Return Values: Procedures do not support return values. They perform actions but do not return data to the caller.
- **Error Handling**: Error messages may be generic and provide limited information. Syntax errors may not be reported accurately.
- **Input Validation**: The INPUT statement assumes that the user will enter integer values. Non-integer input may cause unexpected behavior.
- **Comments**: The language does not support comments. Any attempt to include comments may result in a syntax error.
- **Loops and Arrays**: There is no support for loops (FOR, WHILE) or arrays.
- **Unsupported Tokens**:
  - `TOKEN_REM`: Comments are not supported.
  - `TOKEN_COMMA`: Commas are not used in the current syntax.
  - `TOKEN_SEMICOLON`: Only used in PRINT statements for concatenation.

## Conclusion

This guide covers the basics of writing programs in PATI BASIC, including variables, expressions, statements, and procedures. While PATI BASIC is limited in its current implementation, it provides a foundation for learning fundamental programming concepts.

As the language evolves, additional features and enhancements may be added to improve its capabilities. Keep an eye out for updates to this guide that will include new language features and improvements.

Happy coding!

**Note**: This guide is based on the current state of the PATI BASIC interpreter and may not reflect future updates or changes. Always refer to the latest documentation for the most accurate information.
