import React, { useEffect, useState } from 'react';
import { Navigate } from 'react-router-dom';
import { getCookie } from '../auth/cookieUtil';

interface User {
  id: number;
  first_name: string;
  last_name: string;
  email: string;
  role: number;
}

interface ProtectedRouteProps {
  element: React.ReactNode | any;
}

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ element }) => {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const token = getCookie('jwt')

    if (token) {
      fetch(`${import.meta.env.VITE_API_BASE_URL}/user`, {
        headers: {
          "Authorization": `Bearer ${token}`,
        },
      })
        .then(res => res.json())
        .then((user) => {
          setUser(user);
        })
        .finally(() => setIsLoading(false));
    }
    else {
      setUser(null);
      setIsLoading(false);
    }
  }, []);

  useEffect(() => {
    if (user != null && user.role !== 1) {
      setUser(null);
    }
  }, [user]);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (user == null) {
    return <Navigate to="/login" />;
  }

  return element;
}

export default ProtectedRoute;
