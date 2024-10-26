import { useRef } from "react";

import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { Camera } from "lucide-react";
import { axiosInstance } from "@/services/axiosInstance";

export function User() {
  const { isLoading, isError, data } = useQuery({
    queryKey: ["repoData"],
    queryFn: () => axiosInstance.get("/api/user").then((res) => res.data),
  });

  const queryClient = useQueryClient();

  // quero limpar o cookie jwt e redirecionar para a página de login
  async function handleLogout() {
    await axiosInstance.post("/api/logout");
    window.location.href = "/login";
  }

  const fileInputRef = useRef<HTMLInputElement>(null);

  const handleClick = () => {
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  const mutation = useMutation({
    mutationFn: (formData: FormData) =>
      axiosInstance.post("/api/upload-image", formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["repoData"] });
    },
    onError: (error: any) => {
      console.error("Error uploading image:", error);
    },
  });

  const handleFileChange = async (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const file = event.target.files?.[0];
    if (file) {
      const formData = new FormData();
      formData.append("image", file);

      mutation.mutate(formData);
    }
  };

  if (isLoading) return null;
  if (isError) return <div>Erro ao carregar dados</div>;

  return (
    <>
      <div className="flex w-full flex-col">
        <main className="flex min-h-[calc(100vh_-_theme(spacing.16))] flex-1 flex-col gap-4 bg-muted/40 p-4 md:gap-8 md:p-10">
          <div className="mx-auto grid w-full max-w-6xl gap-2">
            <h1 className="text-3xl font-semibold">Minha conta</h1>
          </div>
          <div className="mx-auto grid w-full max-w-6xl items-start gap-6 md:grid-cols-[180px_1fr] lg:grid-cols-[250px_1fr]">
            <nav
              className="grid gap-4 text-sm text-muted-foreground"
              x-chunk="dashboard-04-chunk-0"
            >
              <a href="#" className="font-semibold text-primary">
                Geral
              </a>
            </nav>
            <div className="grid gap-6">
              <Card x-chunk="dashboard-04-chunk-1">
                <CardHeader>
                  {data?.image ? (
                    <img
                      src={data?.image}
                      alt="profile"
                      className="w-32 h-32 rounded-full object-cover"
                    />
                  ) : (
                    <div>
                      <div
                        className="w-32 h-32 bg-gray-200 rounded-full flex justify-center items-center flex-col gap-4"
                        onClick={handleClick}
                        style={{ cursor: "pointer" }}
                      >
                        <Camera size={40} color="#aaa" />
                        <p
                          className="text-center text-sm text-[#aaa] font-bold"
                          style={{ lineHeight: "12px" }}
                        >
                          Add. imagem
                        </p>
                      </div>
                      <input
                        type="file"
                        ref={fileInputRef}
                        style={{ display: "none" }}
                        onChange={handleFileChange}
                      />
                    </div>
                  )}
                  <CardTitle>Meus dados</CardTitle>

                  <CardDescription>
                    Nome: {data.name}
                    <br />
                    Email: {data.email}
                    <br />
                    Usuário: {data.username}
                  </CardDescription>
                </CardHeader>
                <CardContent></CardContent>
                <CardFooter className="border-t px-6 py-4">
                  <Button
                    onClick={() => handleLogout()}
                    variant={"destructive"}
                  >
                    Sair
                  </Button>
                </CardFooter>
              </Card>
            </div>
          </div>
        </main>
      </div>
    </>
  );
}
