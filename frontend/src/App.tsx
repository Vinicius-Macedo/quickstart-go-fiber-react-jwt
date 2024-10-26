import { Login } from "./pages/auth/Login";
import { Ping } from "./pages/Ping";
import { Register } from "./pages/auth/Register";
// import { Payment } from "./pages/Payment";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { axiosInstance } from "@/services/axiosInstance";
// @ts-ignore

import "./index.css";
import { User } from "./pages/User";
import { ForgotPassword } from "./pages/auth/ForgotPassword";
import { ResetPassword } from "./pages/auth/ResetPassword";
import { SuccessEmail } from "./pages/auth/SuccessEmail";
import { SuccessResetPassword } from "./pages/auth/SuccessResetPassword";

export function App() {
  const { isLoading, isError } = useQuery({
    queryKey: ["pings"],
    queryFn: async () => {
      const { data } = await axiosInstance.get("/api");
      return data;
    },
  });

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (isError) {
    return <div>Error...</div>;
  }

  const router = createBrowserRouter([
    {
      path: "/",
      element: <Login />,
    },
    {
      path: "/ping",
      element: <Ping />,
    },
    {
      path: "/login",
      element: <Login />,
    },
    {
      path: "/user",
      element: <User />,
    },
    {
      path: "/register",
      element: <Register />,
    },
    {
      path: "/forgot-password",
      element: <ForgotPassword />,
    },
    {
      path: "/reset-password",
      element: <ResetPassword />,
    },
    {
      path: "/success-email",
      element: <SuccessEmail />,
    },
    {
      path: "/success-password",
      element: <SuccessResetPassword />,
    },
  ]);

  return (
    <>
      <RouterProvider router={router} />
      <ReactQueryDevtools initialIsOpen={false} />
    </>
  );
}
