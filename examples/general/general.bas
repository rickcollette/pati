REM Example BASIC Program
LET A = 10
LET B = 5

REM Perform addition
LET C = A + B
PRINT "The sum of A and B is: "; C

REM Check if A is greater than B
IF A > B THEN
  PRINT "A is greater than B"
ELSE
  PRINT "A is not greater than B"
END IF

REM Simple loop to print numbers from 1 to 5
LET I = 1
WHILE I <= 5
  PRINT "Number: "; I
  LET I = I + 1
WEND

REM Call a simple procedure
CALL PrintMessage

PROC PrintMessage {
  PRINT "This is a message from a procedure!"
}
END
