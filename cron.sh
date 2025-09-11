#!/bin/bash
# while true; do
echo "start: pprof for $1"
wget -O "routine_$(date).pprof" http://[$1]:6060/debug/pprof/goroutine
wget -O "mem_$(date).pprof" http://[$1]:6060/debug/pprof/heap
wget -O "cpu_$(date).pprof" http://[$1]:6060/debug/pprof/profile
# done

if [[ $2 == "trace" ]]; then
        curl -o trace.out http://[$1]:6060/debug/pprof/trace?seconds=5
fi
echo "done: pprof for $1"

# fdbd:dc06:9:d0e::48
# fdbd:dc06:9:d3a::48
# fdbd:dc06:9:e1a::48
