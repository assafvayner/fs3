import subprocess
import os
import numpy as np

import matplotlib.pyplot as plt


run_command = "python3 python-client/measure.py"
output_dir = "test_data/"


def load_latencies(config):
    directory = output_dir + "load_latencies/"
    if not os.path.exists(directory):
        os.makedirs(directory)
    
    for i in range(0, 4):
        for load in range(500, 5000, 500):
            output_file_name = config + "load_" + str(load) + "_" + str(i) + ".data"
            command = run_command.split(" ")
            command.extend(["-n", "10000", "-d", "30", "-l", str(load), "-s", "server0:5000"])

            with open(directory + output_file_name, "w+") as file:
                subprocess.run(command, stdout=file)


def analyze_load_latencies(config):
    load_to_latency = {}
    load_to_uploads = {}

    for load in range(500, 5000, 500):
        total_99_percentile = 0
        for i in range(0, 4):
            output_file_name = config + "load_" + str(load) + "_" + str(i) + ".data"
            file_name = output_dir + "load_latencies/" + output_file_name

            with open(file_name) as data_file:
                lines = data_file.readlines()
                load_to_uploads[load] = int(lines[0])
                index = 1
                total_time = []
                while index < len(lines):
                    line = lines[index]
                    total_time.append(float(line.split(" ")[-1].strip(")\n")))
                    index = index + 1
                percentile_99 = np.percentile(np.array(total_time), 99)
                total_99_percentile += percentile_99
        load_to_latency[load] = total_99_percentile / 4

    return load_to_latency, load_to_uploads


def file_size_latencies(config):
    directory = output_dir + "file_size_latencies/"
    if not os.path.exists(directory):
        os.makedirs(directory)
    
    for file_size in range(1000, 100000, 10000):
        output_file_name = config + "file_size_" + str(file_size) + ".data"
        command = run_command.split(" ")
        command.extend(["-n", str(file_size), "-l", "50", "-s", "server0:5000"])

        with open(directory + output_file_name, "w+") as file:
            subprocess.run(command, stdout=file)


def analyze_file_size_latencies(config):
    size_to_latency = {}
    size_to_uploads = {}
    
    for file_size in range(1000, 100000, 10000):
        output_file_name = config + "file_size_" + str(file_size) + "_" + ".data"
        file_name = output_dir + "file_size_latencies/" + output_file_name

        with open(file_name) as data_file:
            lines = data_file.readlines()
            size_to_uploads[file_size] = int(lines[0])
            index = 1
            total_time = []
            while index < len(lines):
                line = lines[index]
                total_time.append(float(line.split(" ")[-1].strip(")\n")))
                index = index + 1
            percentile_99 = np.percentile(np.array(total_time), 99)
            size_to_latency[file_size] = percentile_99

    return size_to_latency, size_to_uploads


def plot_data(data_dict, xlabel, ylabel, title, location):
    x = data_dict.keys()
    y = data_dict.values()
    plt.plot(x, y)

    plt.xlabel(xlabel)
    plt.ylabel(ylabel)
    plt.title(title)
    plt.savefig(location)


def main():
    file_size_latencies("bm_")
    size_to_latencies, size_to_uploads = analyze_file_size_latencies("bm_")
    #plot_data(
    #    size_to_latencies, 
    #    "file_size (bytes)", 
    #    "latencies (ms)", 
    #    "file_size v.s. latencies",
    #    "test_data/file_size_latencies_bm.png"
    #)
    # plot_data(
    #     size_to_uploads,
    #     "file_size (bytes)",
    #     "completed copy requests",
    #     "file_size v.s. completed_requests",
    #     "test_data/file_size_completed_requests.png"
    # )
    # print(size_to_latencies)
    # print(size_to_uploads)

    load_latencies("bm_")
    load_to_latencies, load_to_completed_requests = analyze_load_latencies("bm_")
    # print(load_to_latencies)
    # print(uploads_to_latencies)
    # plot_data(
    #     load_to_latencies, 
    #     "load (req/sec)", 
    #     "latencies (ms)", 
    #     "load v.s. latencies bm",
    #     "test_data/load_latencies_bm.png"
    # )
    #plot_data(
    #    load_to_completed_requests,
    #    "load (req/sec)",
    #    "completed_requests",
    #    "load v.s. completed_requests",
    #    "test_data/load_requests.png"
    #)

    
if __name__ == '__main__':
    main()
