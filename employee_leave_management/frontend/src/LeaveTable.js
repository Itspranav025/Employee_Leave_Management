import React, { useState, useEffect } from "react";
import axios from "axios";

const BASE_URL = "http://localhost:8080"; // Backend URL

const LeaveTable = () => {
  const [leaveRecords, setLeaveRecords] = useState([]);

  useEffect(() => {
    fetchLeaveRecords();
  }, []);

  async function fetchLeaveRecords() {
    try {
      const response = await axios.get(`${BASE_URL}/api/leave-records`);
      if (response.status === 200) {
        setLeaveRecords(response.data);
      } else {
        console.error("Received unexpected response status:", response.status);
      }
    } catch (error) {
      console.error("Error fetching leave records:", error);
    }
  }
  

  return (
    <div>
      <h2>Leave Records</h2>
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Leave Type</th>
            <th>From</th>
            <th>To</th>
            <th>Team</th>
            <th>Reporter</th>
          </tr>
        </thead>
        <tbody>
          {leaveRecords.map((record) => (
            <tr key={record.id}>
              <td>{record.fullName}</td>
              <td>{record.leaveType}</td>
              <td>{record.fromDate}</td>
              <td>{record.toDate}</td>
              <td>{record.team}</td>
              <td>{record.reporter}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default LeaveTable;
