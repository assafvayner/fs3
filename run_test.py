import subprocess
import os

import matplotlib.pyplot as plt


run_command = "python3 python-client/measure.py"
output_dir = "test_data/"


def load_latencies():
    directory = output_dir + "load_latencies/"
    if not os.path.exists(directory):
        os.makedirs(directory)
    
    for load in range(100, 2000, 200):
        output_file_name = "load_" + str(load) + ".data"
        command = run_command.split(" ")
        command.extend(["-n", "10000", "-d", "30", "-l", str(load)])

        with open(directory + output_file_name, "w+") as file:
            subprocess.run(command, stdout=file)


def analyze_load_latencies():
    load_to_latency = {}
    load_to_uploads = {}

    for load in range(100, 2000, 200):
        output_file_name = "load_" + str(load) + ".data"
        file_name = output_dir + "load_latencies/" + output_file_name

        with open(file_name) as data_file:
            lines = data_file.readlines()
            load_to_uploads[load] = int(lines[0])
            index = 1
            total_time = 0
            while index < len(lines):
                line = lines[index]
                total_time += float(line.split(" ")[-1].strip(")\n"))
                index = index + 1
            load_to_latency[load] = total_time / int(lines[0])

    return load_to_latency, load_to_uploads


def file_size_latencies():
    directory = output_dir + "file_size_latencies/"
    if not os.path.exists(directory):
        os.makedirs(directory)

    for file_size in range(1000, 1000000, 10000):
        output_file_name = "file_size_" + str(file_size) + ".data"
        command = run_command.split(" ")
        command.extend(["-n", str(file_size), "-l", "50"])

        with open(directory + output_file_name, "w+") as file:
            subprocess.run(command, stdout=file)


def analyze_file_size_latencies():
    size_to_latency = {}
    size_to_uploads = {}

    for file_size in range(1000, 20000, 5000):
        output_file_name = "file_size_" + str(file_size) + ".data"
        file_name = output_dir + "file_size_latencies/" + output_file_name

        with open(file_name) as data_file:
            lines = data_file.readlines()
            size_to_uploads[file_size] = int(lines[0])
            index = 1
            total_time = 0
            while index < len(lines):
                line = lines[index]
                total_time += float(line.split(" ")[-1].strip(")\n"))
                index = index + 1
            size_to_latency[file_size] = total_time / int(lines[0])

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
    pass
    # file_size_latencies()
    # size_to_latencies, size_to_uploads = analyze_file_size_latencies()
    # plot_data(
    #     size_to_latencies, 
    #     "file_size (bytes)", 
    #     "latencies (ms)", 
    #     "file_size v.s. latencies",
    #     "test_data/file_size_latencies.png"
    # )
    # plot_data(
    #     size_to_uploads,
    #     "file_size (bytes)",
    #     "completed copy requests",
    #     "file_size v.s. completed_requests",
    #     "test_data/file_size_completed_requests.png"
    # )
    # print(size_to_latencies)
    # print(size_to_uploads)

    # load_latencies()
    # load_to_latencies, uploads_to_latencies = analyze_load_latencies()
    # print(load_to_latencies)
    # print(uploads_to_latencies)


if __name__ == '__main__':
    main()