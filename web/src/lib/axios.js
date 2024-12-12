import axios from 'axios';

// Axios configuration
const api = axios.create({
  baseURL: 'http://localhost:8080' /* 'http://localhost:8080', 'http://172.27.176.1:8080', 'http://192.168.1.143:8080 */,
  headers: {
    'Content-Type': 'application/json'
  }
});

// Add interceptors for authentication
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('authToken');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

export default api;