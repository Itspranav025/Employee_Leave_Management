--KPI 1 View - Total Employees on Leave in August 2022:
CREATE OR REPLACE VIEW kpi1_view AS
SELECT COUNT(DISTINCT full_name) AS employees_on_leave
FROM leave_records
WHERE from_date >= '2022-08-01' AND from_date < '2022-09-01';

-- KPI2 View - Month with the Most Sick Leaves in 2022:
CREATE OR REPLACE VIEW kpi2_view AS
SELECT EXTRACT(MONTH FROM from_date) AS month, COUNT(DISTINCT full_name) AS employees_on_sick_leave
FROM leave_records
WHERE EXTRACT(YEAR FROM from_date) = 2022 AND leave_type = 'Sick Leave'
GROUP BY month
ORDER BY employees_on_sick_leave DESC
LIMIT 1;

-- KPI3 View - Top 5 Employees with the Most Leave Days in 2023:
CREATE OR REPLACE VIEW kpi3_view AS
SELECT full_name, SUM((to_date - from_date)::int) AS total_leave_days
FROM leave_records
WHERE EXTRACT(YEAR FROM from_date) = 2023
GROUP BY full_name
ORDER BY total_leave_days DESC
LIMIT 5;

-- KPI4 View - Employees on Leave Under Each Manager in Q1 2023:
CREATE OR REPLACE VIEW kpi4_view AS
SELECT reporter AS manager_name,
       COUNT(DISTINCT full_name) AS employees_on_leave
FROM leave_records
WHERE EXTRACT(YEAR FROM from_date) = 2023
      AND EXTRACT(MONTH FROM from_date) BETWEEN 1 AND 3
GROUP BY manager_name
ORDER BY manager_name;

-- KPI5 View - Total Leave Days by Team in 2022:
CREATE OR REPLACE VIEW kpi5_view AS
WITH LeaveData AS (
    SELECT team,SUM((to_date - from_date) + 1) AS total_leave_days
    FROM public.leave_records
    WHERE EXTRACT(YEAR FROM from_date) = 2022
    GROUP BY team
)
SELECT team,total_leave_days,
RANK() OVER (ORDER BY total_leave_days DESC) AS leave_rank
FROM LeaveData;

-- KPI6 View - Leave Type Distribution for Top 2 Teams in 2022:
CREATE OR REPLACE VIEW kpi6_view AS
WITH Top2Teams AS (
    SELECT team
    FROM leave_records
    WHERE EXTRACT(YEAR FROM from_date) = 2022
    GROUP BY team
    ORDER BY SUM((to_date - from_date) + 1) DESC
    LIMIT 2
)
SELECT team,
       leave_type,
       COUNT(*) AS leave_count
FROM leave_records
WHERE EXTRACT(YEAR FROM from_date) = 2022
      AND team IN (SELECT team FROM Top2Teams)
GROUP BY team, leave_type
ORDER BY team, leave_type;
