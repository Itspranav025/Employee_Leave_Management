import requests
import pandas as pd

# Define the API endpoint
api_endpoint = 'http://localhost:8080/api/merge-and-append-data'  # Replace with your actual API endpoint

# Load employee_data from CSV
employee_data = pd.read_csv('./files/employee_data.csv')

# Load manager_data from CSV
manager_data = pd.read_csv('./files/manager_data.csv')

# Merge the two dataframes on the 'employee_name' column
merged_data = pd.merge(employee_data, manager_data, on='employee_name', how='outer')

# Convert 'leave_dates' to the 'yyyy-mm-dd' format
merged_data['leave_dates'] = pd.to_datetime(merged_data['leave_dates'], format='%m/%d/%Y').dt.strftime('%Y-%m-%d')

# Calculate the 'ToDate' based on 'FromDate' and 'leave_duration'
merged_data['todate'] = pd.to_datetime(merged_data['leave_dates']) + pd.to_timedelta(merged_data['leave_duration'], unit='D')

# Convert 'ToDate' to string
merged_data['todate'] = merged_data['todate'].dt.strftime('%Y-%m-%d')

# Add an 'id' column
merged_data['id'] = range(1, len(merged_data) + 1)

# Transform data to match the LeaveRecord structure
transformed_data = []
for index, row in merged_data.iterrows():
    leave_records = {
        'id': row['id'],
        'fullName': row['employee_name'],
        'leaveType': row['leave_type'],
        'fromDate': row['leave_dates'],
        'todate': row['todate'],  # Use the calculated 'ToDate'
        'team': row['team_name'],  # Use the 'team_name' from manager_data
        'medicalCertificateUrl': '',  # If needed
        'reporter': row['manager_name'],  # Fill Reporter with 'manager_name'
    }
    transformed_data.append(leave_records)

# Send the combined and transformed data to the Go API
response = requests.post(api_endpoint, json=transformed_data)

if response.status_code == 200:
    print("Data loaded successfully.")
else:
    print("Failed to load data:", response.text)
