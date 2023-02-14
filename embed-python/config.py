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
    start_time = time.time()
    user_id, bit_rate, concurrent_users, is_prev_open, group_id = sys.argv[1], int(sys.argv[2]), int(sys.argv[3]), sys.argv[4] == 1, int(sys.argv[5])

    punch_times, succ_times, last_punch_success, device_score, net_type, nat_type = int(sys.argv[6]), int(sys.argv[7]), int(sys.argv[8]), int(sys.argv[9]), int(sys.argv[10]), int(sys.argv[11])
    client = ClientInfo(punch_times, succ_times, last_punch_success, device_score, NetType(net_type), nat_type)
    stream = StreamInfo(user_id, bit_rate, concurrent_users, is_prev_open, group_id)
    print(getRuntimeGroupID(client, stream))
