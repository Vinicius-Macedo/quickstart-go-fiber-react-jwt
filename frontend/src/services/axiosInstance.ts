import axios from "axios";

// Função para obter o valor de um cookie pelo nome
function getCookie(name: string): string | undefined {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) return parts.pop()?.split(';').shift();
}

export const axiosInstance = axios.create({
  baseURL: "/proxy",
  headers: {
    "Content-type": "application/json"
  },
  withCredentials: true
});

// Adiciona um interceptor para incluir o CSRF token em todas as requisições
axiosInstance.interceptors.request.use((config) => {
  const csrfToken = getCookie('csrf_'); // Nome do cookie que armazena o CSRF token
  if (csrfToken) {
    config.headers['X-CSRF-Token'] = csrfToken;
  }
  return config;
});
