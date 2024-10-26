import { Footer } from "./Footer";
import { Header } from "./Header";

interface LayoutProps {
  children: React.ReactNode;
}
export function Layout(props : LayoutProps){
    return(
        <>
        <Header/>
          <main className="min-h-screen">
            {props.children}
          </main>
        <Footer/>
        </>
    )
}