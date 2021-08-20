# Timed Quiz

A cli tool to take timed quizzes read from a csv file.

# Usage

```
go build quiz
./quiz -csv filename -limit time_limit
```

Define command line flag `-csv` as the file path to the csv file and `-limit`, the time limit in seconds, default set to 30 seconds.
Quiz csv should have rows of the form: `question, ans`, eg, `1+1, 2`.
Answer as many questions as possible before the time runs out.
