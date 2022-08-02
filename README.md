# Golang Concurrency Patterns
Common and useful golang concurrency patterns I implemented from Rob Pike's famous 2012 Google I/O talk.

[Google I/O Talk](https://www.youtube.com/watch?v=f6kdp27TYZs&t=1021s)

### Common Patterns
- Generator: function that runs goroutine and returns channel
- Multiplexing (fan-in): function that takes multiple channels and pipes to one channel, so that the returned channel receives both outputs
- Daisychaining: functions whose I/O are daisy-chained with channels together
# go-concurrency-patterns
# go-concurrency-patterns
