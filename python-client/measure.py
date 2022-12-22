import threading
import os
import string
import random
import argparse
import grpc
from time import time, sleep

import sys
sys.path.append(".")
sys.path.append("./protos/")
from protos.fs3.fs3_pb2_grpc import Fs3Stub
from protos.fs3.fs3_pb2 import CopyRequest

file_paths = set()
copy_records = {}
get_records = {}
remove_records = {}

TEST_START = time() * 1000


def run_copy(client, size):
    path = gen_file_name()
    content = gen_file_content(size)
    copy_request = CopyRequest(file_path=path, file_content=content)

    try:
        before = time() * 1000
        client.Copy(copy_request)
        after = time() * 1000
        copy_records[(path, size, before - TEST_START)] = after - before
    except:
        raise Exception("Copy request failed")
    file_paths.add(path)


def gen_file_content(size):
    """
    returns randomly generated file contents within visible ascii range such
    that we may be able to view the generated content on the server
    """
    content = os.urandom(size)
    return bytes([((x - 33) % (126-33) + 33) for x in content])


def gen_file_name(length=12):
    """
    generates random file name
    """
    return ''.join(random.choices(string.ascii_letters + string.digits, k=length))


def run_copy_in_thread(server_addr, file_size, rate_per_second, total_time):
    channel = grpc.insecure_channel(server_addr)
    client = Fs3Stub(channel=channel)
    start_time = time()
    while True:
        loop_start = time()
        for _ in range(int(rate_per_second)):
            run_copy(client, file_size)
        curr_time = time()
        if curr_time - start_time >= total_time:
            break
        sec_diff = curr_time - loop_start
        if sec_diff < 1:
            sleep(sec_diff)
    return


def main():
    copy_records.clear()
    get_records.clear()
    remove_records.clear()

    parser = argparse.ArgumentParser(description="Running fs3 tests")
    parser.add_argument("-s", "--server", type=str, default="localhost:5000",
                        help="server address from which to make requests")
    parser.add_argument("-n", "--file_size", type=int, default=1000,
                        help="size of the file to copy")
    parser.add_argument("-d", "--total_time", type=int, default=10,
                        help="total amount of time the test runs (in second)")
    parser.add_argument("-l", "--rate_per_second", type=int, default=100,
                        help="number of copy operations per second")
    parser.add_argument("-t", "--num_threads", type=int, default=10,
                        help="number of threads to spawn")

    args = parser.parse_args()
    server_addr = args.server
    file_size = args.file_size
    total_time = args.total_time
    rate_per_sec = args.rate_per_second
    num_threads = args.num_threads

    rate_per_sec_per_thread = rate_per_sec / num_threads

    try:
        def make_threads_for_copy():
            threads = []
            thread_args = (server_addr, file_size,
                           rate_per_sec_per_thread, total_time)
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

    items = copy_records.items()
    print(len(items))
    for item in items:
        print(item)


if __name__ == '__main__':
    main()
