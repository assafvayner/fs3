from fs3.fs3_pb2_grpc import Fs3Stub
from fs3.fs3_pb2 import CopyRequest
from fs3.fs3_pb2 import GetRequest
from fs3.fs3_pb2 import RemoveRequest

import grpc
import time
import random
import string
import os


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
        before = time.time() * 1000
        result = client.Copy(copy_request)
        after = time.time() * 1000
        print(result.status)
        copy_records[(path, size)] = after - before
    except:
        raise Exception("Copy request failed")


def run_get(file_path):
    get_request = GetRequest(file_path=file_path)
    file_size = get_file_size(file_path=file_path)

    try:
        before = time.time() * 1000
        result = client.Get(get_request)
        after = time.time() * 1000
        print(result.status)
        get_records[(file_path, file_size)] = after - before
    except:
        raise Exception("Get request failed")


def run_remove(file_path):
    remove_request = RemoveRequest(file_path=file_path)
    file_size = get_file_size(file_path=file_path)

    try:
        before = time.time() * 1000
        result = client.Remove(remove_request)
        after = time.time() * 1000
        print(result.status)
        remove_records[(file_path, file_size)] = after - before
    except:
        raise Exception("Remove request failed")


def gen_file_content(size):
    content = os.urandom(size)
    return bytes([((x - 33) % (126-33) + 33) for x in content])


def gen_file_path():
    file_path = ''.join(random.choices(string.ascii_letters + string.digits, k=8))
    while file_path in file_paths:
        file_path = ''.join(random.choices(string.ascii_letters + string.digits, k=8))
    file_paths.append(file_path)
    return file_path


def get_file_size(file_path):
    for pair in copy_records.keys:
        if file_path in pair:
            return file_path[1]
    return -1


def main():
    copy_records = {}
    get_records = {}
    remove_records = {}

    for i in range(1000, 10000, 500):
        run_copy(i)
    for item in copy_records.items():
        print(item)
    
    for file_info in copy_records.keys:
        run_get(file_info[0])
    for item in get_records.items():
        print(item)

    for file_info in copy_records.keys:
        run_remove(file_info[0])
    for item in remove_records.items():
        print(item)


if __name__ == '__main__':
    main()