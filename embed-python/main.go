package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	// This initializes gpython for runtime execution and is essential.
	// It defines forward-declared symbols and registers native built-in modules, such as sys and time.
	_ "github.com/go-python/gpython/stdlib"

	// Commonly consumed gpython
	"github.com/go-python/gpython/py"
	"github.com/go-python/gpython/repl"
	"github.com/go-python/gpython/repl/cli"
)

func main() {
	flag.Parse()
	if err := envPython3(src); err != nil {
		log.Println("envPython3 err", err)
		os.Exit(2)
	}
	// runWithFile(flag.Arg(0))
	// if err := runWithSrc(src); err != nil {
	// 	log.Println(err)
	// }
}

var src = `
print("hh")
a = "abc"
print(a)
print("hh")
print("hh")
print("hh")
print("hh")
`

var config = `
import os
import sys
from enum import Enum
import datetime

class NetType(Enum):
    UNKNOWN = 0
    WIFI    = 1
    MOBILE  = 2
    NT3G    = 3
    NT4G    = 4
    NT5G    = 5
    WEB     = 6

class ClientInfo:
    def __init__(self, punch_times, succ_times, last_punch_success, device_score, net_type, nat_type):
        self.punch_times = punch_times 
        self.succ_times = succ_times 
        self.last_punch_success = last_punch_success
        self.device_score = device_score
        self.net_type = net_type 
        self.nat_type = nat_type

class StreamInfo:
    def __init__(self, user_id, bit_rate, concurrent_users, is_prev_open, group_id):
        self.user_id = user_id
        self.bit_rate = bit_rate
        self.concurrent_users = concurrent_users
        self.is_prev_open = is_prev_open
        self.group_id = group_id

def time_in_range(start, end, x):
    """Return true if x is in the range [start, end]"""
    if start <= end:
        return start <= x <= end
    else:
        return start <= x or x <= end

def getRuntimeGroupID(client_info, stream_info):
    if client_info.net_type == NetType.NT4G and hash(client_info.user_id) % ((sys.maxsize + 1) * 2) % 100 < 100 :
        return 2002  
    if client_info.nat_type == 3 and hash(client_info.user_id) % ((sys.maxsize + 1) * 2) % 100 < 100:
        return 2002 
    # TODO: device score?

    if client_info.punch_times > 100 and client_info.succ_times < 30 and stream.concurrent_users < 1000:
        return 2002
    # start_time = "2:13:57"
    # end_time = "11:46:38"
    t1 = datetime.time(23, 0, 0)
    t2 = datetime.time(1, 0, 0)
    # convert time string to datetime
    # t1 = datetime.datetime.strptime(start_time, "%H:%M:%S")
    # print('Start time:', t1.time())
    # t2 = datetime.datetime.strptime(end_time, "%H:%M:%S")
    # print('End time:', t2.time())

    today = datetime.datetime.today()
    now = datetime.datetime.now()
    # [t1, t2] enable phone p2p 
    print(t1, t2, now.time())
    if time_in_range(t1, t2, now.time()):
        return 2002
    return 2003

'''
argv: 
    user_id, bit_rate, concurrent_users, is_prev_open, group_id,
    punch_times, succ_times, last_punch_success, device_score, net_type, nat_type,
'''
if __name__ == "__main__":
    user_id, bit_rate, concurrent_users, is_prev_open, group_id = sys.argv[2], int(sys.argv[3]), int(sys.argv[4]), sys.argv[5] == 1, int(sys.argv[6])

    punch_times, succ_times, last_punch_success, device_score, net_type, nat_type = int(sys.argv[7]), int(sys.argv[8]), int(sys.argv[9]), int(sys.argv[10]), int(sys.argv[11]), int(sys.argv[12])
    client = ClientInfo(punch_times, succ_times, last_punch_success, device_score, NetType(net_type), nat_type)
    stream = StreamInfo(user_id, bit_rate, concurrent_users, is_prev_open, group_id)
    print(getRuntimeGroupID(client, stream))
`

func envPython3(src string) error {
	start := time.Now()
	defer func() {
		log.Printf("envPython3 cost: %v", time.Since(start))
	}()
	// fmt.Printf("src: %q", src)
	// d1 := []byte(src)
	// if err := os.WriteFile("./dat1.py", d1, 0644); err != nil {
	// 	panic(err)
	// }
	// python3 := exec.Command("python3", "./config.py", "1")
	python3 := exec.Command("python3", "-c", config, "user1", "4390", "1000", "100", "0", "2003", "0", "0", "0", "0", "5", "1")
	// stdout, err := python3.StdoutPipe()
	// CombinedOutput
	python3.Stdout = os.Stdout
	python3.Stderr = os.Stderr
	if err := python3.Run(); err != nil {
		panic(err)
	}
	return nil

	// r := bufio.NewReader(stdout)
	// if err := python3.Start(); err != nil {
	// 	return err
	// }
	// if s, err := r.ReadString('\n'); err != nil {
	// 	return err
	// } else {
	// 	fmt.Printf("%q", s)
	// }
	// if err := python3.Wait(); err != nil {
	// 	return err
	// }

	// return nil
}

func runWithSrc(src string) error {
	// See type Context interface and related docs
	ctx := py.NewContext(py.DefaultContextOpts())
	defer ctx.Close()

	var err error
	_, err = py.RunSrc(ctx, src, "a simple source", nil)
	return err
}

func runWithFile(pyFile string) error {

	// See type Context interface and related docs
	opts := py.DefaultContextOpts()
	opts.SysPaths = append(opts.SysPaths, "/opt/homebrew/Cellar/python@3.10/3.10.8/Frameworks/Python.framework/Versions/3.10/lib/python3.10")
	fmt.Println(opts)
	ctx := py.NewContext(opts)

	// This drives modules being able to perform cleanup and release resources
	defer ctx.Close()

	var err error
	if len(pyFile) == 0 {
		replCtx := repl.New(ctx)

		fmt.Print("\n=======  Entering REPL mode, press Ctrl+D to exit  =======\n")

		_, err = py.RunFile(ctx, "lib/REPL-startup.py", py.CompileOpts{}, replCtx.Module)
		if err == nil {
			cli.RunREPL(replCtx)
		}

	} else {
		_, err = py.RunFile(ctx, pyFile, py.CompileOpts{}, nil)
	}

	if err != nil {
		py.TracebackDump(err)
	}

	return err
}
