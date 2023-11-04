import React from "react";
import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import FrontendTask from "./frontendtask";
import axios from "axios";

jest.mock("axios");

describe("FrontendTask", () => {
  beforeEach(() => {
    axios.post.mockResolvedValue({}); // Mock axios post request
    axios.post.mockClear(); // Reset the mock function before each test
  });

  it("renders the form initially", () => {
    render(<FrontendTask />);
    const formElement = screen.getByText("Leave Application Form");
    expect(formElement).toBeInTheDocument();
  });

  it("submits leave application form", async () => {
    // Mock the axios.post function
    axios.post.mockResolvedValue({});

    render(<FrontendTask />);

    const nameInput = screen.getByLabelText(/Name:/i);
    fireEvent.change(nameInput, {
      target: { value: "John Doe" },
    });

    const leaveTypeInput = screen.getByLabelText(/Leave Type:/i);
    fireEvent.change(leaveTypeInput, {
      target: { value: "Sick Leave" },
    });

    const fromDateInput = screen.getByLabelText(/From:/i);
    fireEvent.change(fromDateInput, {
      target: { value: "2023-09-10" },
    });

    const toDateInput = screen.getByLabelText(/To:/i);
    fireEvent.change(toDateInput, {
      target: { value: "2023-09-12" },
    });

    const teamInput = screen.getByLabelText(/Team:/i);
    fireEvent.change(teamInput, {
      target: { value: "AI" },
    });

    // Simulate uploading a file
    const file = new File(["(binary data)"], "medical_certificate.pdf", {
      type: "application/pdf",
    });
    const medicalCertificateInput = screen.getByLabelText("Medical Certificate:");
    fireEvent.change(medicalCertificateInput, {
      target: {
        files: [file],
      },
    });

    const reporterInput = screen.getByLabelText(/Reporter:/i);
    fireEvent.change(reporterInput, {
      target: { value: "David Davis" },
    });

    fireEvent.click(screen.getByText("Submit"));

    await waitFor(() => {
      expect(axios.post).toHaveBeenCalledTimes(1);
      expect(axios.post).toHaveBeenCalledWith(
        "http://localhost:8080/api/apply-leave",
        {
          fullName: "John Doe",
          leaveType: "Sick Leave",
          fromDate: "2023-09-10",
          toDate: "2023-09-12",
          team: "AI",
          medicalCertificateUrl: expect.any(String),
          reporter: "David Davis",
        }
      );

      // Ensure that the form fields are reset after submission
      expect(nameInput).toHaveValue("");
      expect(leaveTypeInput).toHaveValue("Casual Leave"); // Adjusted to match your component's initial value
      expect(fromDateInput).toHaveValue("");
      expect(toDateInput).toHaveValue("");
      expect(teamInput).toHaveValue("AI"); // Adjusted to match your component's initial value
      expect(reporterInput).toHaveValue("");
    });
  });

  it("displays leave records table after clicking 'View Leave Records'", async () => {
    render(<FrontendTask />);
    fireEvent.click(screen.getByText("View Leave Records"));
    const leaveRecordsTable = await screen.findByText("Leave Records");
    expect(leaveRecordsTable).toBeInTheDocument();
  });
});
