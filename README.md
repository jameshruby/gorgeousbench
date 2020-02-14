# gorgeousbench
A fork of https://github.com/cespare/prettybench
Reason for this fork is that I would like to go to different direction with the tool. 
(Eq. Original tool is printing time unit that makes most sense, while I would like to print different time units at once,
whih should be easier when comparing results.)

A tool for transforming `go test`'s benchmark output a bit to make it nicer for humans.

## Problem

Go benchmarks are great, particularly when used in concert with benchcmp. But the output can be a bit hard to
read:

![before](/screenshots/before.png)

Especially if you would like to compare 

## Solution

    $ go get github.com/jameshruby/gorgeousbench
    $ go test -bench=. | gorgeousbench


## Notes


## To Do (maybe)

* EVERYTHING
