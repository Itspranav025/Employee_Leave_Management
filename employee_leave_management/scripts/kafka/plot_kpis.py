import matplotlib.pyplot as plt
import pandas as pd
from delta.tables import DeltaTable

# Path to the Delta Lake directory
delta_table_path = "path_to_delta_table"

# Load the Delta table
delta_table = DeltaTable.forPath(spark, delta_table_path)

# Query the stored KPIs
kpi1_data = delta_table.toDF().where("kpi_id = 1").toPandas()
kpi2_data = delta_table.toDF().where("kpi_id = 2").toPandas()
kpi3_data = delta_table.toDF().where("kpi_id = 3").toPandas()

# Plot KPI 1: Average, Minimum and Maximum Time to Submit Form
plt.figure(figsize=(10, 6))
plt.plot(kpi1_data['hour'], kpi1_data['avg_time'], label='Average Time')
plt.plot(kpi1_data['hour'], kpi1_data['min_time'], label='Minimum Time')
plt.plot(kpi1_data['hour'], kpi1_data['max_time'], label='Maximum Time')
plt.xlabel('Hour')
plt.ylabel('Time (seconds)')
plt.title('KPI 1: Time to Submit Form')
plt.legend()
plt.grid()
plt.show()

# Plot KPI 2: Number of Active Sessions
plt.figure(figsize=(10, 6))
plt.plot(kpi2_data['timestamp'], kpi2_data['active_sessions'])
plt.xlabel('Timestamp')
plt.ylabel('Active Sessions')
plt.title('KPI 2: Number of Active Sessions')
plt.grid()
plt.show()

# Plot KPI 3: Sessions with Open and Close without Form Submit
plt.figure(figsize=(8, 5))
plt.bar(['Sessions'], [kpi3_data['count'][0]], color='blue', label='Sessions with Open and Close (No Submit)')
plt.xlabel('KPI 3')
plt.ylabel('Count')
plt.title('KPI 3: Sessions with Open and Close (No Submit)')
plt.legend()
plt.show()
