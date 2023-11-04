import requests

# Define the API endpoint
api_endpoint = 'http://localhost:8080/api/merge-and-append-data'  # Replace with your actual API endpoint

# Load the JSON data from the file
with open('./Pythonscript/test_data.json', 'r') as file:
    test_data = file.read()

# Send a POST request with the JSON data
response = requests.post(api_endpoint, json=test_data)

if response.status_code == 200:
    print("Data loaded successfully.")
else:
    print("Failed to load data:", response.text)
