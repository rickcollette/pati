# **PATI and PATI-Linter**

**P**arse, **A**ssemble, **T**ranslate, and **I**nterpret  
(Also, Patricia is my grandmothers’ name, and that’s really why I am calling this PATI)

---

## **Table of Contents**

1. **Introduction**  
2. **PATI BASIC Language Overview**  
3. **Using PATI Interpreter**  
4. **Using PATI-Linter**  
5. **VSCode Integration**  
6. **Common Issues and Troubleshooting**  
7. **Best Practices for PATI BASIC Programming**

---

## **1\. Introduction**

PATI is a toolset designed to facilitate writing and running BASIC programs in the PATI BASIC language. It comes with a robust linter, PATI-Linter, which helps you identify and fix common issues in your code before execution. Additionally, a Visual Studio Code (VSCode) extension is available to integrate PATI features directly into your development environment.

---

## **2\. PATI BASIC Language Overview**

PATI BASIC is a simple, structured language with support for common programming constructs. Here’s a basic rundown of the syntax:

* **Variables**: Single-letter variable names like `A`, `B`, etc.  
* **Control Structures**: `IF ... THEN`, `RETURN`, `END`, `PROC` for procedure definition.  
* **Basic Arithmetic**: `+`, `-`, `*`, `/` for mathematical operations.  
* **I/O Operations**: `PRINT` to display output and `INPUT` to read user input.  
* **Comments**: Use `REM` to add comments to your code.

**Example Code:**

basic  
Copy code  
`REM Simple PATI BASIC Program`  
`LET A = 5`  
`LET B = 10`  
`IF A < B THEN PRINT "A is less than B"`  
`END`

---

## **3\. Using PATI Interpreter**

### **Running a PATI BASIC Program**

**Prepare Your Program**: Write your PATI BASIC code and save it in a file with the `.bas` extension.  
**Run Using PATI Interpreter**:  

* Open your terminal.

Use the following command:  

```bash  
pati <file.bas>
```

* Replace `<file.bas>` with the name of your file.  

**Output**: The interpreter will execute your program and display any output or errors in the terminal.

### **Features of the Interpreter**

* **Line-by-Line Execution**: PATI interprets your code line by line, helping you debug easily.  
* **Error Reporting**: Detailed error messages are provided for easy identification and correction of issues.

---

## **4\. Using PATI-Linter**

PATI-Linter is a command-line tool designed to check your PATI BASIC code for syntax and logical errors before you run it.

### **Running PATI-Linter**

* **Prepare Your Code**: Save your PATI BASIC code in a `.bas` file.  
* **Run the Linter**:  
* Open your terminal.

Use the command:  

```bash  
`pati-linter <file.bas>`
```

* Replace `<file.bas>` with the path to your file.  

* **Review Warnings and Errors**: The linter will analyze your code and provide warnings or errors with line numbers to help you correct issues.

### **Key Features of PATI-Linter**

* **Syntax Checking**: Detects missing or mismatched syntax elements.  
* **Variable Declaration Check**: Ensures all variables are declared before use.  
* **Unmatched Braces Check**: Identifies any unmatched `{` or `}` in your code.  
* **Unused Variables**: Warns if you have declared variables that are not used in your code.  
* **Type Mismatch Detection**: Flags variables used inconsistently across different data types.

---

## **5. VSCode Integration**

The PATI VSCode Extension provides an integrated development experience with syntax highlighting, linting, and easy access to run your programs.

### **Installing the Extension**

1. **Open VSCode** and go to the **Extensions** view.  
2. Search for **PATI BASIC Extension**.  
3. Click **Install** to add the extension to your editor.

### **Features in VSCode**

* **Syntax Highlighting**: Provides color coding for keywords, strings, numbers, comments, and variables.  
* **Automatic Linting**: Runs PATI-Linter automatically when you save a `.bas` file.  
* **Manual Linting**: Use the `Run PATI Linter` command from the command palette.  
* **Check for Updates**: Easily check for updates to the linter from within VSCode.

### **Running the Linter in VSCode**

1. **Open Your File**: Open a `.bas` file in VSCode.  
2. **Save or Use Command**: The linter runs automatically on save or use the command:  
   * Open the command palette (`Ctrl+Shift+P` or `Cmd+Shift+P` on Mac).  
   * Type `Run PATI Linter` and select it.

---

## **6\. Common Issues and Troubleshooting**

1. **Linter Not Found**: Ensure the `pati-linter` binary is correctly installed and the path is set in the extension settings.

**Permission Errors**: On Unix-based systems, you may need to give the linter executable permissions using:  

```bash  
chmod +x pati-linter
```

2. **Update Issues**: If checking for updates fails, ensure you have an active internet connection and the repository URL is correct.

---

## **7\. Best Practices for PATI BASIC Programming**

1. **Use Comments Wisely**: Use `REM` to annotate your code but keep it concise.  
2. **Declare Variables**: Always declare your variables before use to avoid warnings.  
3. **Consistent Formatting**: Use consistent spacing and indentation for better readability.  
4. **Test Incrementally**: Run your code in small chunks to catch errors early.

Enjoy coding with PATI and make sure to lint your code frequently for a smooth programming experience\!
