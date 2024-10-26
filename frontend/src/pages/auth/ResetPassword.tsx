import { useLocation } from "react-router-dom";
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
import { useState, useEffect } from "react";
import { HandleFormsInputChange, FormValue } from "@/lib/formUtils";
import { useMutation } from "@tanstack/react-query";
import { axiosInstance } from "@/services/axiosInstance";
import { useNavigate } from "react-router-dom";
const useQuery = () => {
  return new URLSearchParams(useLocation().search);
};

export function ResetPassword() {
  const query = useQuery();

  const [form, setForm] = useState<FormValue>({
    token: "",
    password: "",
  });
  const [error, setError] = useState<string>("");

  useEffect(() => {
    const token = query.get("token");
    if (token) {
      setForm({ ...form, token });
    }
  }, []);

  const navigation = useNavigate();

  const mutation = useMutation({
    mutationFn: () => axiosInstance.post("/api/reset-password", form),
    onSuccess: () => {
      navigation("/success-password");
    },
    onError: (error: any) => {
      console.log(error.response?.data);
      setError(
        error.response?.data?.error || "Um erro ocorreu, contate o suporte"
      );
    },
  });

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    mutation.mutate();
  };

  console.log("Form:", form);

  return (
    <>
      <div className="flex justify-center items-center">
        <Card className="w-full max-w-sm">
          <form onSubmit={(e) => handleSubmit(e)}>
            <CardHeader>
              <CardTitle className="text-2xl text-center">
                Esqueceu sua senha
              </CardTitle>
              <CardDescription className="text-center">
                Digite a nova senha
              </CardDescription>
            </CardHeader>
            <CardContent className="grid gap-4">
              <div className="grid gap-2">
                <Label htmlFor="email">Nova senha</Label>
                <Input
                  id="email"
                  type="password"
                  name="password"
                  onChange={(e) =>
                    HandleFormsInputChange(e, form as FormValue, setForm)
                  }
                  required
                />
              </div>
              <div className="grid gap-2">
                <Label htmlFor="password">Confirme a senha</Label>
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
            </CardContent>
            <CardFooter className="flex flex-col items-center gap-4">
              <Button disabled={mutation.isPending}>
                {mutation.isPending ? "Carregando..." : "Enviar"}
              </Button>
              <p className="text-center text-destructive">{error}</p>
            </CardFooter>
          </form>
        </Card>
      </div>
    </>
  );
}
