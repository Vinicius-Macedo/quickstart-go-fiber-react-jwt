import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { HandleFormsInputChange, FormValue } from "@/lib/formUtils";
import { useState } from "react";
import { axiosInstance } from "@/services/axiosInstance";
import { useMutation } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";

export function Register() {
  const [form, setForm] = useState<FormValue>({
    name: "",
    username: "",
    email: "",
    password: "",
    repeat_password: "",
  });
  const [error, setError] = useState<string>("");

  const navigate = useNavigate();

  const mutation = useMutation({
    mutationFn: () =>
      axiosInstance.post("/api/register", {
        name: form.name,
        username: form.username,
        email: form.email,
        password: form.password,
      }),
    onSuccess: () => {
      axiosInstance.post("/api/login", {
        email: form.email,
        password: form.password,
      });

      navigate("/user");
    },
    onError: (error: any) => {
      console.log("Failed", error);
      setError(error.response?.data?.error || "An error occurred");
    },
  });

  function handleSubmit(event: React.FormEvent) {
    event.preventDefault();

    if (form.password !== form.repeat_password) {
      alert("Passwords do not match");
      return;
    }

    mutation.mutate();
    console.log(form);
  }

  return (
    <div
      className="flex min-h-screen justify-center items-center
    "
    >
      <Card className="mx-auto max-w-sm">
        <CardHeader>
          <CardTitle className="text-xl">Criar conta</CardTitle>
          <CardDescription>
            {error ? error : "Preencha os campos para criar uma conta"}
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={(e) => handleSubmit(e)}>
            <div className="grid gap-4">
              <div className="grid gap-2">
                <Label htmlFor="email">Nome completo</Label>
                <Input
                  id="name"
                  type="name"
                  name="name"
                  placeholder="John Galt"
                  onChange={(e) =>
                    HandleFormsInputChange(e, form as FormValue, setForm)
                  }
                  required
                />
              </div>
              <div className="grid gap-2">
                <Label htmlFor="username">Usuário</Label>
                <Input
                  id="username"
                  type="text"
                  name="username"
                  placeholder="john_galt"
                  onChange={(e) =>
                    HandleFormsInputChange(e, form as FormValue, setForm)
                  }
                  required
                />
              </div>
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
                <Label htmlFor="password">Senha</Label>
                <Input
                  id="password"
                  type="password"
                  name="password"
                  onChange={(e) =>
                    HandleFormsInputChange(e, form as FormValue, setForm)
                  }
                />
              </div>
              <div className="grid gap-2">
                <Label htmlFor="repeat-password">Repita a senha</Label>
                <Input
                  id="repeat-password"
                  type="password"
                  name="repeat_password"
                  onChange={(e) =>
                    HandleFormsInputChange(e, form as FormValue, setForm)
                  }
                />
              </div>
              <Button
                disabled={mutation.isPending}
                type="submit"
                className="w-full"
              >
                {mutation.isPending ? "Carregando..." : "Criar conta"}
              </Button>
            </div>
            <p className="mt-4 text-center text-sm">
              Já tem uma conta?{" "}
              <span
                onClick={() => navigate("/login")}
                className="underline cursor-pointer"
              >
                Entre
              </span>
            </p>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
