import pandas as pd
import matplotlib.pyplot as plt


def parse_time(time_str):
    if 'µs' in time_str:
        return float(time_str.replace('µs', '')) / 1000000  # Convert milliseconds to seconds
    elif 'ms' in time_str:
        return float(time_str.replace('ms', '')) / 1000  # Convert milliseconds to seconds
    elif 's' in time_str:
        return float(time_str.replace('s', ''))
    else:
        return float(time_str)  # Assuming it's already in seconds if no unit is specified



# Read data from CSV
file_path = 'results.csv'  # Replace with your CSV file path
data = pd.read_csv(file_path)

# Extracting data
N_values = data['N']
sequential_times = data['sequential']  # Assuming times are in milliseconds
concurrent2_times = data['concurrent2']  # Assuming times are in milliseconds
concurrent5_times = data['concurrent5']  # Assuming times are in milliseconds

# Convert times to seconds for plotting
sequential_times_s = [parse_time(t) / 1000 for t in sequential_times]
concurrent2_times_s = [parse_time(t) / 1000 for t in concurrent2_times]
concurrent5_times_s = [parse_time(t) / 1000 for t in concurrent5_times]

# Plotting
plt.figure(figsize=(10, 6))
plt.plot(N_values, sequential_times_s, label='Sequential', marker='o')
plt.plot(N_values, concurrent2_times_s, label='Concurrent 2', marker='o')
plt.plot(N_values, concurrent5_times_s, label='Concurrent 5', marker='o')

plt.xscale('log')
plt.yscale('log')
plt.xlabel('N (scale of task)')
plt.ylabel('Time (seconds)')
plt.title('Performance Comparison: Sequential vs Concurrent Processing')
plt.legend()
plt.grid(True)
plt.show()


