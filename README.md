# prefixMatch

*GO ENV DETAILS:*
  1. Version: go version go1.16.2 windows/amd64
  2. GOHOSTOS=windows
  3. GOHOSTARCH=amd64

To Execute Code:
- Sample Input Command:
  ```
  go run main.go --prefix-source-file ./source/sample_prefixes.txt --input-string QE9tIPvasd
  ```
- Sample Output:
  ```
  2022/04/24 15:43:48 Prefix source contains: 264986 Input string: QE9tIPvasd
  2022/04/24 15:43:48 Result: QE9tIPv is the longest prefix found!
  ```
To Excute Test Cases:
- Under FileReader Directory
  - Sample input Command
    ```
    C:\Users\shivgosw\go\src\github.com\ShivangGoswami\prefixesMatch\filereader>go test -v
    ```
  - Sample Output
    ```
    === RUN   TestReadFile
    === RUN   TestReadFile/invalid
    2022/04/24 15:46:35 Error in reading prefix file: open : The system cannot find the file specified.
    === RUN   TestReadFile/valid
    --- PASS: TestReadFile (0.13s)
        --- PASS: TestReadFile/invalid (0.00s)
        --- PASS: TestReadFile/valid (0.13s)
    PASS
    ok      github.com/ShivangGoswami/prefixesMatch/filereader      0.267s
    ```
- Under Matcher Directory
  - Sample input Command
    ```
    C:\Users\shivgosw\go\src\github.com\ShivangGoswami\prefixesMatch\matcher>go test -v 
    ```
  - Sample Output
    ```
    === RUN   TestMatcher
    === RUN   TestMatcher/valid
    === RUN   TestMatcher/valid#01
    === RUN   TestMatcher/invalid
    --- PASS: TestMatcher (0.00s)
        --- PASS: TestMatcher/valid (0.00s)
        --- PASS: TestMatcher/valid#01 (0.00s)
        --- PASS: TestMatcher/invalid (0.00s)
    PASS
    ok      github.com/ShivangGoswami/prefixesMatch/matcher 0.125s
    ```
