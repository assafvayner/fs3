from protos.fs3.fs3_pb2_grpc import fs3Stub
from protos.fs3.fs3_pb2 import CopyRequest
from protos.fs3.fs3_pb2 import CopyReply

import grpc
import time
import random
import string
import os

ip_address = "localhost:5000"
file_names = []
copy_records = {}


def run_copy(size):
    channel = grpc.insecure_channel(ip_address)
    client = fs3Stub(channel=channel)
    path = gen_file_path()
    content = gen_file_content(size)
    copy_request = CopyRequest(file_path=path, file_content=content)

    try:
        before = round(time.time() * 1000)
        client.Copy(copy_request)
        after = round(time.time() * 1000)
        copy_records[size] = after - before
    except:
        raise Exception("Copy request failed")


def gen_file_content(size):
    content = list(os.urandom(size))
    return content


def gen_file_path():
    file_name = ''.join(random.choices(string.ascii_letters + string.digits, k=8))
    while file_name in file_names:
        file_name = ''.join(random.choices(string.ascii_letters + string.digits, k=8))
    file_names.append(file_name)
    return file_name

def main():
    pass