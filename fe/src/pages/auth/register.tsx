import React, { useState } from "react";

interface RegisterFormData {
  first_name: string;
  last_name: string;
  email: string;
  password: string;
}

function Register() {
  const [formData, setFormData] = useState<RegisterFormData>({
    first_name: '',
    last_name: '',
    email: '',
    password: '',
  });

  const [registraitonResult, setRegistrationResult] = useState("")

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleFormSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    try {
      const response = await fetch(`${import.meta.env.VITE_API_BASE_URL}/register`, {
        method: 'POST',
        body: new FormData(event.currentTarget),
      });

      if (response.ok) {
        //registration complete
        setRegistrationResult("Successfully registered, you can now login.");
      } else {
        setRegistrationResult("Registration failed");
      }
    } catch (error) {
      setRegistrationResult("Registration failed" + error);
    }
  };

  return (
    <>
      <h1>Register</h1>
      {formData.first_name}<br />
      {formData.last_name}<br />
      {formData.email}<br />
      {formData.password}
      {
        registraitonResult != '' ? <blockquote>{registraitonResult}</blockquote> : ''
      }
      <form onSubmit={handleFormSubmit}>
        <div>
          <label>First name:</label>
          <input type="text" name="first_name" value={formData.first_name} onChange={handleInputChange} />
        </div>
        <div>
          <label>Last name:</label>
          <input type="text" name="last_name" value={formData.last_name} onChange={handleInputChange} />
        </div>
        <div>
          <label>Email:</label>
          <input type="text" name="email" value={formData.email} onChange={handleInputChange} />
        </div>
        <div>
          <label>Password:</label>
          <input type="password" name="password" value={formData.password} onChange={handleInputChange} />
        </div>
        <button className="btn btn-primary">Register</button>
      </form>
    </>
  );
}

export default Register;