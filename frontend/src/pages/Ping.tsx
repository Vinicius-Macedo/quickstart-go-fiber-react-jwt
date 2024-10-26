import { useQuery } from '@tanstack/react-query';
import { axiosInstance } from "@/services/axiosInstance";

// Função para fazer a requisição
const fetchPing = async () => {
  const response = await axiosInstance.get('http://localhost:3000');
  return response.data;
};

export function Ping() {
  // Use useQuery to make the request
  const { data, error, isLoading } = useQuery({
    queryKey: ['ping'],
    queryFn: fetchPing,
  });

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error: {error.message}</div>;

  return (
    <div>
      <h1>Ping Response</h1>
      <pre>{JSON.stringify(data, null, 2)}</pre>
    </div>
  );
}