from fs3.fs3_pb2_grpc import Fs3Stub
from fs3.fs3_pb2 import CopyRequest
from fs3.fs3_pb2 import GetRequest
from fs3.fs3_pb2 import RemoveRequest

import grpc
from time import time, sleep
import random
import string
import os
import threading


ip_address = "localhost:5000"

file_paths = []
copy_records = {}
get_records = {}
remove_records = {}

channel = grpc.insecure_channel(ip_address)
client = Fs3Stub(channel=channel)


def run_copy(size):
    path = gen_file_path()
    content = gen_file_content(size)
    copy_request = CopyRequest(file_path=path, file_content=content)

    try:
        before = time() * 1000
        result = client.Copy(copy_request)
        after = time() * 1000
        print(result.status)
        copy_records[(path, size, before)] = after - before
    except:
        raise Exception("Copy request failed")


def run_get(file_path):
    get_request = GetRequest(file_path=file_path)
    file_size = get_file_size(file_path=file_path)

    try:
        before = time() * 1000
        result = client.Get(get_request)
        after = time() * 1000
        print(result.status)
        get_records[(file_path, file_size, before)] = after - before
    except:
        raise Exception("Get request failed")


def run_remove(file_path):
    remove_request = RemoveRequest(file_path=file_path)
    file_size = get_file_size(file_path=file_path)

    try:
        before = time() * 1000
        result = client.Remove(remove_request)
        after = time() * 1000
        print(result.status)
        remove_records[(file_path, file_size, before)] = after - before
    except:
        raise Exception("Remove request failed")


def gen_file_content(size):
    content = os.urandom(size)
    return bytes([((x - 33) % (126-33) + 33) for x in content])


def gen_file_path():
    file_path = ''.join(random.choices(
        string.ascii_letters + string.digits, k=8))
    while file_path in file_paths:
        file_path = ''.join(random.choices(
            string.ascii_letters + string.digits, k=8))
    file_paths.append(file_path)
    return file_path


def get_file_size(file_path):
    for pair in copy_records.keys:
        if file_path in pair:
            return file_path[1]
    return -1


def run_copy_in_thread(file_size, rate_per_second, total_time):
    start_time = time()
    while True:
        loop_start = time()
        for _ in range(rate_per_second):
            run_copy(file_size)
        curr_time = time()
        if curr_time - start_time >= total_time:
            break
        sec_diff = curr_time - loop_start
        if sec_diff < 1:
            sleep(sec_diff)


# def run_get_in_thread(start, end):
#     for index in range(start, end):
#         file_info = copy_records.keys[index]
#         run_get(file_info[0])


# def run_remove_in_thread(start, end):
#     for index in range(start, end):
#         file_info = copy_records.keys[index]
#         run_remove(file_info[index])


def main():
    copy_records.clear()
    get_records.clear()
    remove_records.clear()

    total_time = 30
    file_size = 1000
    rate_per_sec = 20
    num_threads = 10

    rate_per_sec_per_thread = rate_per_sec / num_threads

    try:
        def make_threads_for_copy():
            threads = []
            thread_args = (file_size, rate_per_sec_per_thread, total_time)
            for num in range(num_threads):
                thread = threading.Thread(
                    target=run_copy_in_thread, args=thread_args)
                threads.insert(num, thread)
            return threads

        threads = make_threads_for_copy()
        for thread in threads:
            thread.start()
        sleep(total_time)
        for thread in threads:
            thread.join()
    except:
        raise Exception("Thread error")

    for item in copy_records.items():
        print(item)


if __name__ == '__main__':
    main()
