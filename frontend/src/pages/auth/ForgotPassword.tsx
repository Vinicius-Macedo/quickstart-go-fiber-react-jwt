import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useState } from "react";
import { useMutation } from "@tanstack/react-query";
import { axiosInstance } from "@/services/axiosInstance";
import { useNavigate } from "react-router-dom";

export function ForgotPassword() {
  const [form, setForm] = useState({
    email: "",
  });

  const [error, setError] = useState<string>("");
  const navigate = useNavigate();

  const mutation = useMutation({
    mutationFn: () => axiosInstance.post("/api/forgot-password", form),
    onSuccess: () => {
      navigate("/success-email");
    },
    onError: (error: any) => {
      console.log(error.response?.data);
      setError(
        error.response?.data?.error || "Um erro ocorreu, contate o suporte"
      );
    },
  });

  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    mutation.mutate();
  };

  return (
    <>
      <div className="w-screen min-h-screen flex justify-center items-center">
        <Card className="w-full max-w-sm">
          <form onSubmit={(e) => handleSubmit(e)}>
            <CardHeader>
              <CardTitle className="text-2xl text-center">
                Esqueceu sua senha
              </CardTitle>
              <CardDescription className="text-center">
                Digite seu email
              </CardDescription>
            </CardHeader>
            <CardContent className="grid gap-4">
              <div className="grid gap-2">
                <Label htmlFor="email">Email</Label>
                <Input
                  id="email"
                  type="email"
                  name="email"
                  placeholder="m@example.com"
                  value={form.email}
                  onChange={(e) => setForm({ ...form, email: e.target.value })}
                  required
                />
              </div>
            </CardContent>
            <CardFooter className="flex flex-col items-center gap-4">
              <Button disabled={mutation.isPending}>
                {mutation.isPending ? "Carregando..." : "Enviar"}
              </Button>
              <p
                className="text-sm underline cursor-pointer"
                onClick={() => navigate("/login")}
              >
                voltar
              </p>
              <p className="text-center text-destructive">{error}</p>
            </CardFooter>
          </form>
        </Card>
      </div>
    </>
  );
}
