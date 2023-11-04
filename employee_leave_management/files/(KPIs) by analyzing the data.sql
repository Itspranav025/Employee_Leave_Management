--KPI 1 - Total Employees on Leave in August 2022:
SELECT employees_on_leave FROM kpi1_view;

--KPI 2 - Month with the Most Sick Leaves in 2022:
SELECT month, employees_on_sick_leave FROM kpi2_view;

--KPI 3 - Top 5 Employees with the Most Leave Days in 2023:
SELECT full_name, total_leave_days FROM kpi3_view;

--KPI 4 - Employees on Leave Under Each Manager in Q1 2023:
SELECT manager_name, employees_on_leave FROM kpi4_view;

--KPI 5 - Total Leave Days by Team in 2022:
SELECT team, total_leave_days FROM kpi5_view;

--KPI 6 - Leave Type Distribution for Top 2 Teams in 2022:
SELECT team, leave_type, leave_count FROM kpi6_view;