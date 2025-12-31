# Spreadsheet Functions Spec Delta

## ADDED Requirements

### Requirement: Function Registry

The system SHALL provide a function registry for built-in Excel-compatible functions.

#### Scenario: Function interface definition
- GIVEN the Function interface
- THEN the interface has method Name() string
- AND the interface has method MinArgs() int
- AND the interface has method MaxArgs() int (returns -1 for unlimited)
- AND the interface has method Call(args []Value, ctx *EvalContext) (Value, error)

#### Scenario: Register function
- GIVEN a FunctionRegistry
- AND a function implementation SumFunction
- WHEN Register(SumFunction) is called
- THEN function is registered with name "SUM"
- AND Get("SUM") returns the function
- AND Get("sum") returns the function (case-insensitive)

#### Scenario: Default function registry
- GIVEN DefaultFunctionRegistry() is called
- THEN registry contains 60+ built-in functions
- AND all math, logical, text, lookup, datetime, stats functions are registered

### Requirement: SUM Function

The system SHALL implement the SUM function for summing numeric values.

#### Scenario: SUM with number arguments
- GIVEN SUM(1, 2, 3)
- WHEN function is evaluated
- THEN result is 6

#### Scenario: SUM with range argument
- GIVEN SUM(A1:A5)
- AND A1:A5 contains [10, 20, 30, 40, 50]
- WHEN function is evaluated
- THEN result is 150

#### Scenario: SUM ignores text in ranges
- GIVEN SUM(A1:A3)
- AND A1:A3 contains [10, "text", 20]
- WHEN function is evaluated
- THEN result is 30 (text ignored)

#### Scenario: SUM with mixed arguments
- GIVEN SUM(5, A1:A3, 10)
- AND A1:A3 contains [1, 2, 3]
- WHEN function is evaluated
- THEN result is 21 (5+1+2+3+10)

#### Scenario: SUM propagates errors
- GIVEN SUM(A1:A3)
- AND A1 contains #DIV/0!
- WHEN function is evaluated
- THEN result is #DIV/0!

### Requirement: AVERAGE Function

The system SHALL implement the AVERAGE function for calculating arithmetic mean.

#### Scenario: AVERAGE with numbers
- GIVEN AVERAGE(10, 20, 30)
- WHEN function is evaluated
- THEN result is 20

#### Scenario: AVERAGE with range
- GIVEN AVERAGE(A1:A4)
- AND A1:A4 contains [10, 20, 30, 40]
- WHEN function is evaluated
- THEN result is 25

#### Scenario: AVERAGE ignores text
- GIVEN AVERAGE(A1:A4)
- AND A1:A4 contains [10, "text", 20, 30]
- WHEN function is evaluated
- THEN result is 20 (average of 10, 20, 30)

#### Scenario: AVERAGE with no numbers
- GIVEN AVERAGE(A1:A3)
- AND A1:A3 all contain text
- WHEN function is evaluated
- THEN result is #DIV/0! (no numbers to average)

### Requirement: COUNT and COUNTA Functions

The system SHALL implement COUNT (count numbers) and COUNTA (count non-empty) functions.

#### Scenario: COUNT counts numbers only
- GIVEN COUNT(1, "text", 2, TRUE, 3)
- WHEN function is evaluated
- THEN result is 3 (1, 2, 3 are numbers)

#### Scenario: COUNTA counts all non-empty
- GIVEN COUNTA(1, "text", 2, TRUE, 3)
- WHEN function is evaluated
- THEN result is 5 (all values)

#### Scenario: COUNT with range
- GIVEN COUNT(A1:A5)
- AND A1:A5 contains [10, "text", 20, TRUE, 30]
- WHEN function is evaluated
- THEN result is 3

#### Scenario: COUNTA with range
- GIVEN COUNTA(A1:A5)
- AND A1:A5 contains [10, "text", 20, "", 30]
- WHEN function is evaluated
- THEN result is 4 (empty string not counted)

### Requirement: MIN and MAX Functions

The system SHALL implement MIN (minimum) and MAX (maximum) functions.

#### Scenario: MIN with numbers
- GIVEN MIN(5, 2, 8, 1, 9)
- WHEN function is evaluated
- THEN result is 1

#### Scenario: MAX with numbers
- GIVEN MAX(5, 2, 8, 1, 9)
- WHEN function is evaluated
- THEN result is 9

#### Scenario: MIN with range
- GIVEN MIN(A1:A5)
- AND A1:A5 contains [50, 20, 80, 10, 90]
- WHEN function is evaluated
- THEN result is 10

#### Scenario: MIN ignores text
- GIVEN MIN(A1:A4)
- AND A1:A4 contains [50, "text", 20, 10]
- WHEN function is evaluated
- THEN result is 10

### Requirement: ROUND Functions

The system SHALL implement ROUND, ROUNDUP, and ROUNDDOWN functions.

#### Scenario: ROUND to integer
- GIVEN ROUND(2.5, 0)
- WHEN function is evaluated
- THEN result is 2 (banker's rounding)

#### Scenario: ROUND with positive digits
- GIVEN ROUND(3.14159, 2)
- WHEN function is evaluated
- THEN result is 3.14

#### Scenario: ROUND with negative digits
- GIVEN ROUND(1234, -2)
- WHEN function is evaluated
- THEN result is 1200

#### Scenario: ROUNDUP always rounds up
- GIVEN ROUNDUP(2.1, 0)
- WHEN function is evaluated
- THEN result is 3

#### Scenario: ROUNDDOWN always rounds down
- GIVEN ROUNDDOWN(2.9, 0)
- WHEN function is evaluated
- THEN result is 2

### Requirement: Mathematical Functions

The system SHALL implement ABS, SQRT, POWER, and MOD functions.

#### Scenario: ABS with negative number
- GIVEN ABS(-5.5)
- WHEN function is evaluated
- THEN result is 5.5

#### Scenario: SQRT with positive number
- GIVEN SQRT(16)
- WHEN function is evaluated
- THEN result is 4

#### Scenario: SQRT with negative number
- GIVEN SQRT(-1)
- WHEN function is evaluated
- THEN result is #NUM! (cannot take square root of negative)

#### Scenario: POWER function
- GIVEN POWER(2, 3)
- WHEN function is evaluated
- THEN result is 8

#### Scenario: MOD function
- GIVEN MOD(10, 3)
- WHEN function is evaluated
- THEN result is 1

### Requirement: PRODUCT Function

The system SHALL implement the PRODUCT function for multiplying values.

#### Scenario: PRODUCT with numbers
- GIVEN PRODUCT(2, 3, 4)
- WHEN function is evaluated
- THEN result is 24

#### Scenario: PRODUCT with range
- GIVEN PRODUCT(A1:A3)
- AND A1:A3 contains [2, 5, 3]
- WHEN function is evaluated
- THEN result is 30

### Requirement: Conditional Aggregation Functions

The system SHALL implement SUMIF, COUNTIF, and AVERAGEIF functions.

#### Scenario: SUMIF with numeric criteria
- GIVEN SUMIF(A1:A5, ">5", B1:B5)
- AND A1:A5 contains [3, 7, 5, 9, 2]
- AND B1:B5 contains [10, 20, 30, 40, 50]
- WHEN function is evaluated
- THEN result is 60 (20+40, where A>5)

#### Scenario: SUMIF with text criteria
- GIVEN SUMIF(A1:A4, "Yes", B1:B4)
- AND A1:A4 contains ["Yes", "No", "Yes", "No"]
- AND B1:B4 contains [10, 20, 30, 40]
- WHEN function is evaluated
- THEN result is 40 (10+30)

#### Scenario: COUNTIF function
- GIVEN COUNTIF(A1:A5, ">5")
- AND A1:A5 contains [3, 7, 5, 9, 2]
- WHEN function is evaluated
- THEN result is 2 (7 and 9)

#### Scenario: AVERAGEIF function
- GIVEN AVERAGEIF(A1:A4, ">5", B1:B4)
- AND A1:A4 contains [3, 7, 5, 9]
- AND B1:B4 contains [10, 20, 30, 40]
- WHEN function is evaluated
- THEN result is 30 (average of 20, 40)

### Requirement: IF Function

The system SHALL implement the IF logical function.

#### Scenario: IF with true condition
- GIVEN IF(5>3, "Yes", "No")
- WHEN function is evaluated
- THEN result is "Yes"

#### Scenario: IF with false condition
- GIVEN IF(2>5, "Yes", "No")
- WHEN function is evaluated
- THEN result is "No"

#### Scenario: IF with cell reference condition
- GIVEN IF(A1>10, "High", "Low")
- AND A1 contains 15
- WHEN function is evaluated
- THEN result is "High"

#### Scenario: IF with two arguments (omit else)
- GIVEN IF(TRUE, 1)
- WHEN function is evaluated
- THEN result is 1
- WHEN IF(FALSE, 1) is evaluated
- THEN result is FALSE

#### Scenario: Nested IF
- GIVEN IF(A1>10, "High", IF(A1>5, "Medium", "Low"))
- AND A1 contains 7
- WHEN function is evaluated
- THEN result is "Medium"

### Requirement: Logical Functions

The system SHALL implement AND, OR, and NOT functions.

#### Scenario: AND with all true
- GIVEN AND(TRUE, TRUE, TRUE)
- WHEN function is evaluated
- THEN result is TRUE

#### Scenario: AND with one false
- GIVEN AND(TRUE, FALSE, TRUE)
- WHEN function is evaluated
- THEN result is FALSE

#### Scenario: OR with one true
- GIVEN OR(FALSE, TRUE, FALSE)
- WHEN function is evaluated
- THEN result is TRUE

#### Scenario: OR with all false
- GIVEN OR(FALSE, FALSE, FALSE)
- WHEN function is evaluated
- THEN result is FALSE

#### Scenario: NOT function
- GIVEN NOT(TRUE)
- WHEN function is evaluated
- THEN result is FALSE

### Requirement: Error Handling Functions

The system SHALL implement IFERROR and IFNA functions.

#### Scenario: IFERROR with no error
- GIVEN IFERROR(5/2, "Error")
- WHEN function is evaluated
- THEN result is 2.5

#### Scenario: IFERROR with error
- GIVEN IFERROR(5/0, "Error")
- WHEN function is evaluated
- THEN result is "Error"

#### Scenario: IFNA with non-NA error
- GIVEN IFNA(#DIV/0!, "NA value")
- WHEN function is evaluated
- THEN result is #DIV/0! (not #N/A, so not handled)

#### Scenario: IFNA with NA error
- GIVEN IFNA(#N/A, "Not found")
- WHEN function is evaluated
- THEN result is "Not found"

### Requirement: IFS Function

The system SHALL implement the IFS multi-condition function.

#### Scenario: IFS with first condition true
- GIVEN IFS(A1>10, "High", A1>5, "Medium", TRUE, "Low")
- AND A1 contains 12
- WHEN function is evaluated
- THEN result is "High"

#### Scenario: IFS with second condition true
- GIVEN IFS(A1>10, "High", A1>5, "Medium", TRUE, "Low")
- AND A1 contains 7
- WHEN function is evaluated
- THEN result is "Medium"

#### Scenario: IFS with default case
- GIVEN IFS(A1>10, "High", A1>5, "Medium", TRUE, "Low")
- AND A1 contains 3
- WHEN function is evaluated
- THEN result is "Low"

### Requirement: Text Functions

The system SHALL implement CONCATENATE, LEFT, RIGHT, MID, and LEN functions.

#### Scenario: CONCATENATE function
- GIVEN CONCATENATE("Hello", " ", "World")
- WHEN function is evaluated
- THEN result is "Hello World"

#### Scenario: LEFT function
- GIVEN LEFT("Hello", 2)
- WHEN function is evaluated
- THEN result is "He"

#### Scenario: RIGHT function
- GIVEN RIGHT("Hello", 2)
- WHEN function is evaluated
- THEN result is "lo"

#### Scenario: MID function
- GIVEN MID("Hello", 2, 3)
- WHEN function is evaluated
- THEN result is "ell" (starts at position 2, length 3)

#### Scenario: LEN function
- GIVEN LEN("Hello")
- WHEN function is evaluated
- THEN result is 5

#### Scenario: LEFT with default num_chars
- GIVEN LEFT("Hello")
- WHEN function is evaluated
- THEN result is "H" (defaults to 1)

### Requirement: Case Conversion and Trimming Functions

The system SHALL implement UPPER, LOWER, and TRIM functions.

#### Scenario: UPPER function
- GIVEN UPPER("hello")
- WHEN function is evaluated
- THEN result is "HELLO"

#### Scenario: LOWER function
- GIVEN LOWER("HELLO")
- WHEN function is evaluated
- THEN result is "hello"

#### Scenario: TRIM function
- GIVEN TRIM("  Hello  World  ")
- WHEN function is evaluated
- THEN result is "Hello World" (leading/trailing/extra spaces removed)

### Requirement: Text Search and Replace Functions

The system SHALL implement FIND and SUBSTITUTE functions.

#### Scenario: FIND function
- GIVEN FIND("l", "Hello")
- WHEN function is evaluated
- THEN result is 3 (first occurrence)

#### Scenario: FIND function case-sensitive
- GIVEN FIND("h", "Hello")
- WHEN function is evaluated
- THEN result is #VALUE! (case-sensitive, not found)

#### Scenario: FIND with start position
- GIVEN FIND("l", "Hello", 4)
- WHEN function is evaluated
- THEN result is 4 (second "l")

#### Scenario: SUBSTITUTE function
- GIVEN SUBSTITUTE("Hello", "l", "L")
- WHEN function is evaluated
- THEN result is "HeLLo" (all occurrences)

#### Scenario: SUBSTITUTE specific instance
- GIVEN SUBSTITUTE("Hello", "l", "L", 1)
- WHEN function is evaluated
- THEN result is "HeLlo" (first occurrence only)

### Requirement: TEXT Function

The system SHALL implement the TEXT function for formatting numbers.

#### Scenario: TEXT with decimal format
- GIVEN TEXT(1234.5, "0.00")
- WHEN function is evaluated
- THEN result is "1234.50"

#### Scenario: TEXT with thousands separator
- GIVEN TEXT(1234.5, "#,##0.00")
- WHEN function is evaluated
- THEN result is "1,234.50"

#### Scenario: TEXT with date format
- GIVEN TEXT(DATE(2000,1,1), "mm/dd/yyyy")
- WHEN function is evaluated
- THEN result is "01/01/2000"

### Requirement: VLOOKUP Function

The system SHALL implement the VLOOKUP vertical lookup function.

#### Scenario: VLOOKUP exact match
- GIVEN VLOOKUP(5, A1:B10, 2, FALSE)
- AND A1:A10 contains [1, 3, 5, 7, 9, ...]
- AND B1:B10 contains [10, 30, 50, 70, 90, ...]
- WHEN function is evaluated
- THEN result is 50 (found 5 in column 1, returned column 2)

#### Scenario: VLOOKUP approximate match
- GIVEN VLOOKUP(6, A1:B10, 2, TRUE)
- AND A1:A10 contains sorted [1, 3, 5, 7, 9, ...]
- AND B1:B10 contains [10, 30, 50, 70, 90, ...]
- WHEN function is evaluated
- THEN result is 50 (largest value <= 6 is 5)

#### Scenario: VLOOKUP not found
- GIVEN VLOOKUP(100, A1:B10, 2, FALSE)
- AND 100 is not in A1:A10
- WHEN function is evaluated
- THEN result is #N/A

#### Scenario: VLOOKUP invalid column index
- GIVEN VLOOKUP(5, A1:B10, 5, FALSE)
- WHEN function is evaluated
- THEN result is #REF! (column 5 doesn't exist in range)

### Requirement: HLOOKUP Function

The system SHALL implement the HLOOKUP horizontal lookup function.

#### Scenario: HLOOKUP exact match
- GIVEN HLOOKUP("Q2", A1:D5, 3, FALSE)
- AND A1:D1 contains ["Q1", "Q2", "Q3", "Q4"]
- AND A3:D3 contains [100, 200, 300, 400]
- WHEN function is evaluated
- THEN result is 200

### Requirement: INDEX and MATCH Functions

The system SHALL implement INDEX (array lookup) and MATCH (index search) functions.

#### Scenario: INDEX function
- GIVEN INDEX(A1:C10, 5, 2)
- AND cell B5 contains 42
- WHEN function is evaluated
- THEN result is 42

#### Scenario: INDEX with row only
- GIVEN INDEX(A1:A10, 5)
- AND A5 contains 100
- WHEN function is evaluated
- THEN result is 100

#### Scenario: MATCH exact match
- GIVEN MATCH(5, A1:A10, 0)
- AND A1:A10 contains [1, 3, 5, 7, 9, ...]
- WHEN function is evaluated
- THEN result is 3 (5 found at index 3)

#### Scenario: MATCH approximate match
- GIVEN MATCH(6, A1:A10, 1)
- AND A1:A10 contains sorted [1, 3, 5, 7, 9, ...]
- WHEN function is evaluated
- THEN result is 3 (largest value <= 6 is 5 at index 3)

#### Scenario: INDEX-MATCH combination
- GIVEN INDEX(B1:B10, MATCH(5, A1:A10, 0))
- AND A1:A10 contains [1, 3, 5, 7, 9, ...]
- AND B1:B10 contains [10, 30, 50, 70, 90, ...]
- WHEN function is evaluated
- THEN result is 50

### Requirement: XLOOKUP Function

The system SHALL implement the XLOOKUP modern lookup function.

#### Scenario: XLOOKUP exact match
- GIVEN XLOOKUP(5, A1:A10, B1:B10)
- AND A1:A10 contains [1, 3, 5, 7, 9, ...]
- AND B1:B10 contains [10, 30, 50, 70, 90, ...]
- WHEN function is evaluated
- THEN result is 50

#### Scenario: XLOOKUP not found
- GIVEN XLOOKUP(100, A1:A10, B1:B10)
- AND 100 not in A1:A10
- WHEN function is evaluated
- THEN result is #N/A

### Requirement: CHOOSE Function

The system SHALL implement the CHOOSE function for selecting from a list.

#### Scenario: CHOOSE function
- GIVEN CHOOSE(2, "A", "B", "C")
- WHEN function is evaluated
- THEN result is "B"

#### Scenario: CHOOSE with cell reference index
- GIVEN CHOOSE(A1, "Low", "Medium", "High")
- AND A1 contains 3
- WHEN function is evaluated
- THEN result is "High"

#### Scenario: CHOOSE out of range
- GIVEN CHOOSE(5, "A", "B", "C")
- WHEN function is evaluated
- THEN result is #VALUE!

### Requirement: Date/Time Functions

The system SHALL implement TODAY, NOW, DATE, TIME, and date component extraction functions.

#### Scenario: TODAY function
- GIVEN TODAY()
- WHEN function is evaluated on 2025-01-15
- THEN result is Excel serial number for 2025-01-15 (45677)

#### Scenario: NOW function
- GIVEN NOW()
- WHEN function is evaluated on 2025-01-15 at 12:00:00
- THEN result is Excel serial number 45677.5 (45677 + 0.5 for noon)

#### Scenario: DATE function
- GIVEN DATE(2000, 1, 1)
- WHEN function is evaluated
- THEN result is 36526 (Excel serial for Jan 1, 2000)

#### Scenario: TIME function
- GIVEN TIME(12, 0, 0)
- WHEN function is evaluated
- THEN result is 0.5 (noon is half a day)

#### Scenario: YEAR function
- GIVEN YEAR(DATE(2000, 1, 1))
- WHEN function is evaluated
- THEN result is 2000

#### Scenario: MONTH function
- GIVEN MONTH(DATE(2000, 7, 15))
- WHEN function is evaluated
- THEN result is 7

#### Scenario: DAY function
- GIVEN DAY(DATE(2000, 7, 15))
- WHEN function is evaluated
- THEN result is 15

#### Scenario: HOUR function
- GIVEN HOUR(TIME(14, 30, 0))
- WHEN function is evaluated
- THEN result is 14

#### Scenario: MINUTE function
- GIVEN MINUTE(TIME(14, 30, 0))
- WHEN function is evaluated
- THEN result is 30

#### Scenario: SECOND function
- GIVEN SECOND(TIME(14, 30, 45))
- WHEN function is evaluated
- THEN result is 45

### Requirement: DATEDIF Function

The system SHALL implement the DATEDIF function for date differences.

#### Scenario: DATEDIF years
- GIVEN DATEDIF(DATE(2000,1,1), DATE(2005,1,1), "Y")
- WHEN function is evaluated
- THEN result is 5

#### Scenario: DATEDIF months
- GIVEN DATEDIF(DATE(2000,1,1), DATE(2000,7,1), "M")
- WHEN function is evaluated
- THEN result is 6

#### Scenario: DATEDIF days
- GIVEN DATEDIF(DATE(2000,1,1), DATE(2000,1,11), "D")
- WHEN function is evaluated
- THEN result is 10

### Requirement: Statistical Functions

The system SHALL implement MEDIAN, MODE, STDEV, VAR, PERCENTILE, and QUARTILE functions.

#### Scenario: MEDIAN odd count
- GIVEN MEDIAN(1, 2, 3, 4, 5)
- WHEN function is evaluated
- THEN result is 3

#### Scenario: MEDIAN even count
- GIVEN MEDIAN(1, 2, 3, 4)
- WHEN function is evaluated
- THEN result is 2.5

#### Scenario: MODE function
- GIVEN MODE(1, 2, 2, 3, 3, 3, 4)
- WHEN function is evaluated
- THEN result is 3 (most common value)

#### Scenario: STDEV function
- GIVEN STDEV(1, 2, 3, 4, 5)
- WHEN function is evaluated
- THEN result is approximately 1.58114 (sample standard deviation)

#### Scenario: VAR function
- GIVEN VAR(1, 2, 3, 4, 5)
- WHEN function is evaluated
- THEN result is 2.5 (sample variance)

#### Scenario: PERCENTILE function
- GIVEN PERCENTILE(A1:A10, 0.75)
- AND A1:A10 contains [10, 20, 30, 40, 50, 60, 70, 80, 90, 100]
- WHEN function is evaluated
- THEN result is 75 (75th percentile)

#### Scenario: QUARTILE function
- GIVEN QUARTILE(A1:A10, 1)
- AND A1:A10 contains [10, 20, 30, 40, 50, 60, 70, 80, 90, 100]
- WHEN function is evaluated
- THEN result is 25 (first quartile)

### Requirement: Function Argument Validation

The system SHALL validate function argument counts and types.

#### Scenario: Too few arguments
- GIVEN SUM() with no arguments
- WHEN function is evaluated
- THEN result is #VALUE! (minimum 1 argument required)

#### Scenario: Too many arguments with fixed count
- GIVEN IF(A1, B1, C1, D1) with 4 arguments
- WHEN function is evaluated
- THEN result is #VALUE! (maximum 3 arguments)

#### Scenario: Unlimited arguments supported
- GIVEN SUM(A1, A2, A3, ..., A100) with 100 arguments
- WHEN function is evaluated
- THEN result is sum of all 100 values (no limit)

### Requirement: Function Error Handling

The system SHALL handle errors in function arguments appropriately.

#### Scenario: Function propagates error argument
- GIVEN SUM(1, #DIV/0!, 3)
- WHEN function is evaluated
- THEN result is #DIV/0!

#### Scenario: IFERROR suppresses error
- GIVEN IFERROR(SUM(1, #DIV/0!, 3), 0)
- WHEN function is evaluated
- THEN result is 0

#### Scenario: Function with type error
- GIVEN SQRT("text")
- WHEN function is evaluated
- THEN result is #VALUE!

### Requirement: Excel Compatibility

The system SHALL match Excel behavior for all implemented functions.

#### Scenario: Empty cell handling
- GIVEN SUM(A1:A3)
- AND A2 is empty
- WHEN function is evaluated
- THEN empty cell treated as 0 (same as Excel)

#### Scenario: Boolean in formula context
- GIVEN SUM(TRUE, FALSE, 1)
- WHEN function is evaluated
- THEN result is 2 (TRUE=1, FALSE=0 in formula context)

#### Scenario: String number coercion
- GIVEN SUM("5", "10")
- WHEN function is evaluated
- THEN result is #VALUE! (text strings not coerced in SUM)

#### Scenario: Case-insensitive string comparison
- GIVEN IF("hello"="HELLO", "Match", "No match")
- WHEN function is evaluated
- THEN result is "Match" (Excel is case-insensitive)
