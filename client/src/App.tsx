import { QueryClientProvider, QueryClient } from "@tanstack/react-query";
import { Header } from "./components/Header";
import { Layout } from "./components/Layout";
import { Todos } from "./components/Todos";
import "./style.css";
import { AlertsContextProvider } from "./context/AlertsContext";
import { AlertsContainer } from "./components/AlertsContainer";

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <AlertsContextProvider>
        <div className="App">
          <Layout>
            <AlertsContainer />
            <Header />
            <Todos />
          </Layout>
        </div>
      </AlertsContextProvider>
    </QueryClientProvider>
  );
}

export default App;
