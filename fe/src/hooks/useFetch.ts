import { useState, useEffect } from 'react';

function useFetch(endpoint: string) {
  const apiUrl = `${import.meta.env.VITE_API_BASE_URL}/${endpoint}`;
  const [data, setData] = useState<any | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetch(apiUrl)
      .then((res) => {
        if (!res.ok) {
          throw new Error('Network response was not ok');
        }
        return res.json();
      })
      .then((data) => setData(data))
      .catch((error) => setError(error.message));
  }, [apiUrl]);

  return [data, error];
}

export default useFetch;