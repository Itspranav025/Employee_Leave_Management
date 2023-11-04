import React from "react";
import { render, screen } from "@testing-library/react";
import App from "./App";

test("renders App component with header and footer", () => {
  render(<App />);
  
  // Check if the header text is present
  const headerElement = screen.getByText("Leave Form App");
  expect(headerElement).toBeInTheDocument();
  
  // Check if the footer text containing the current year is present
  const currentYear = new Date().getFullYear();
  const footerElement = screen.getByText(`© ${currentYear} Your Company. All rights reserved.`);
  expect(footerElement).toBeInTheDocument();
});
