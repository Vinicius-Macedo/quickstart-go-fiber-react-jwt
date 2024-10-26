import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useState } from "react";
import { HandleFormsInputChange, FormValue } from "@/lib/formUtils";
import { useMutation } from "@tanstack/react-query";
import { axiosInstance } from "@/services/axiosInstance";
import { useNavigate } from "react-router-dom";

export function Login() {
  const [form, setForm] = useState<FormValue>({
    email: "",
    password: "",
  });
  const [error, setError] = useState<string>("");

  const navigate = useNavigate();

  const mutation = useMutation({
    mutationFn: () => axiosInstance.post("/api/login", form),
    onSuccess: () => {
      navigate("/user");
    },
    onError: (error: any) => {
      console.log("Failed", error);
      setError(error.response?.data?.message || "An error occurred");
    },
  });

  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    mutation.mutate();
  };

  return (
    <>
      <div className="w-full h-full min-h-screen lg:grid lg:grid-cols-2">
        <div className="flex items-center justify-center py-12 h-full">
          <form
            onSubmit={(e) => handleSubmit(e)}
            className="mx-auto grid w-[350px] gap-6"
          >
            <div className="grid gap-2 text-center">
              <h1 className="text-3xl font-bold">Login</h1>
              <p className="text-balance text-muted-foreground">
                {error ? (
                  <span className="text-red-500">{error}</span>
                ) : (
                  "Preencha os campos abaixo para acessar"
                )}
              </p>
            </div>
            <div className="grid gap-4">
              <div className="grid gap-2">
                <Label htmlFor="email">Email</Label>
                <Input
                  id="email"
                  type="email"
                  name="email"
                  placeholder="m@example.com"
                  onChange={(e) =>
                    HandleFormsInputChange(e, form as FormValue, setForm)
                  }
                  required
                />
              </div>
              <div className="grid gap-2">
                <div className="flex items-center">
                  <Label htmlFor="password">Senha</Label>
                  <p
                    onClick={() => navigate("/forgot-password")}
                    className="ml-auto inline-block text-sm underline cursor-pointer"
                  >
                    Esqueceu sua senha?
                  </p>
                </div>
                <Input
                  id="password"
                  type="password"
                  name="password"
                  onChange={(e) =>
                    HandleFormsInputChange(e, form as FormValue, setForm)
                  }
                  required
                />
              </div>
              <Button type="submit" className="w-full">
                Entrar
              </Button>
            </div>
            <p className="mt-4 text-center text-sm">
              Não tem uma conta?{" "}
              <span
                onClick={() => navigate("/register")}
                className="underline cursor-pointer"
              >
                Registre-se
              </span>
            </p>
          </form>
        </div>
        <div className="hidden bg-muted lg:flex justify-center items-center">
          <img
            src="/vite.svg"
            alt="Image"
            className="w-40 dark:brightness-[0.2] dark:grayscale"
          />
        </div>
      </div>
    </>
  );
}
