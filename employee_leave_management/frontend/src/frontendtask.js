import React, { useState } from "react";
import { BrowserRouter as Router, Link, Route, Routes } from "react-router-dom";
import LeaveTable from "./LeaveTable"; // Import the LeaveTable component
import axios from "axios";

import "./frontendstyle.css";

const BASE_URL = "http://localhost:8080"; // Backend URL

const FrontendTask = () => {
  const [fullName, setFullName] = useState("");
  const [leaveType, setLeaveType] = useState("Casual Leave");
  const [fromDate, setFromDate] = useState("");
  const [toDate, setToDate] = useState("");
  const [team, setTeam] = useState("Team A");
  const [medicalCertificate, setMedicalCertificate] = useState(null);
  const [reporter, setReporter] = useState("");
  
  
  const [showForm, setShowForm] = useState(true); // Added state to control form visibility

  const handleViewLeaveTable = () => {
    setShowForm(false); // Hide the form when clicking on the link
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    // trackEvent("FormSubmit");
    let medicalCertificateUrl = "";
    if (leaveType === "Sick Leave" && medicalCertificate) {
      medicalCertificateUrl = await uploadMedicalCertificate(
        medicalCertificate
      );
    }

    const leaveData = {
      fullName,
      leaveType,
      fromDate,
      toDate,
      team,
      medicalCertificateUrl,
      reporter,
    };

    try {
      await applyLeave(leaveData);
      // Reset form fields
      setFullName("");
      setLeaveType("Casual Leave");
      setFromDate("");
      setToDate("");
      setTeam("Team A");
      setMedicalCertificate(null);
      setReporter("");
    } catch (error) {
      console.error("Error applying leave:", error);
    }
  };

  const uploadMedicalCertificate = async (file) => {
    const formData = new FormData();
    formData.append("medicalCertificate", file);

    try {
      const response = await axios.post(
        `${BASE_URL}/api/upload-medical-certificate`,
        formData,
        {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        }
      );
      return response.data.medicalCertificateUrl;
    } catch (error) {
      console.error("Error uploading medical certificate:", error);
      throw error;
    }
  };

  const applyLeave = async (leaveData) => {
    try {
      await axios.post(`${BASE_URL}/api/apply-leave`, leaveData);
    } catch (error) {
      console.error("Error applying leave:", error);
      throw error;
    }
  };

  return (
    <Router>
      <div>
        {showForm &&(
        <form onSubmit={handleSubmit}>
          <div>
            <h1>Leave Application Form</h1>
            <div>
              <label>Name:</label>
              <input
                type="text"
                value={fullName}
                onChange={(e) => setFullName(e.target.value)}
                required
              />
            </div>

            <div>
              <label>Leave Type:</label>
              <select
                value={leaveType}
                onChange={(e) => setLeaveType(e.target.value)}
                required
              >
                <option value="Casual Leave">Casual Leave</option>
                <option value="Earned leave">Earned leave</option>
                <option value="Sick Leave">Sick Leave</option>
              </select>
            </div>
            <div>
              <label>From:</label>
              <input
                type="date"
                value={fromDate}
                onChange={(e) => setFromDate(e.target.value)}
                required
              />
            </div>
            <div>
              <label>To:</label>
              <input
                type="date"
                value={toDate}
                onChange={(e) => setToDate(e.target.value)}
                required
              />
            </div>
            <div>
              <label>Team:</label>
              <select
                value={team}
                onChange={(e) => setTeam(e.target.value)}
                required
              >
                <option value="AI">AI</option>
                <option value="Consulting">Consulting</option>
                <option value="Data Engineering">Data Engineering</option>
                <option value="dev ops">dev ops</option>
                <option value="Finance">Finance</option>
                <option value="HR">HR</option>
                <option value="IT">IT</option>
                <option value="Platform team">Platform team</option>
                <option value="Sales">Sales</option>
                <option value="Security">Security</option>
              </select>
            </div>
            {leaveType === "Sick Leave" && (
              <div>
                <label>Medical Certificate:</label>
                <input
                  type="file"
                  accept=".pdf,.png"
                  onChange={(e) => setMedicalCertificate(e.target.files[0])}
                  required
                />
              </div>
            )}
            <div>
              <label>Reporter:</label>
              <select
                value={reporter}
                onChange={(e) => setReporter(e.target.value)}
                required
              >
                <option value="David Davis">David Davis</option>
                <option value="Curtis Gardner">Curtis Gardner</option>
                <option value="Brenda Jefferson">Brenda Jefferson</option>
                <option value="Sabrina Martin">Sabrina Martin</option>
                <option value="Alex Gonzalez">Alex Gonzalez</option>
                <option value="Rebecca Davis">Rebecca Davis</option>
                <option value="Barbara Olson">Barbara Olson</option>
                <option value="Juan Myers DVM">Juan Myers DVM</option>
                <option value="Barbara Olson">Barbara Olson</option>
                <option value="Juan Myers DVM">Juan Myers DVM</option>
                <option value="Tammy Carey">Tammy Carey</option>
                

              </select>
            </div>
            <div>
              <button type="submit">Submit</button>
            </div>
          </div>
        </form>
        )}

        <Link to="/leave-records" onClick={handleViewLeaveTable}>View Leave Records</Link>

        <Routes>
          <Route path="/leave-records" element={<LeaveTable />} />
        </Routes>
      </div>
    </Router>
  );
};

export default FrontendTask;