import React, { useState } from 'react';
import { setCookie } from './cookieUtil';

type LoginFormData = {
  email: string;
  password: string;
};

function Login() {
  const [formData, setFormData] = useState<LoginFormData>({
    email: '',
    password: '',
  });
  const [loginResult, setLoginResult] = useState("")

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      const response = await fetch(`${import.meta.env.VITE_API_BASE_URL}/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });
      if (response.ok) {
        setLoginResult("Logged in successfully")
      }

      if (!response.ok) {
        throw new Error('Login failed');
      }

      const data = await response.json();
      // Set the cookie with the token
      setCookie('jwt', data.token, 1);


      // Redirect to the user profile page
    } catch (err) {
      setLoginResult("Login error: " + err)
      console.error(err);
    }
  };

  return (
    <>
      <h1>Login</h1>
      {loginResult}<br />
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="email">Email:</label>
          <input
            type="email"
            id="email"
            name="email"
            value={formData.email}
            onChange={handleChange}
            required
          />
        </div>
        <div>
          <label htmlFor="password">Password:</label>
          <input
            type="password"
            id="password"
            name="password"
            value={formData.password}
            onChange={handleChange}
            required
          />
        </div>
        <button className="btn btn-primary" type="submit">Login</button>
      </form>

    </>
  );
}

export default Login;