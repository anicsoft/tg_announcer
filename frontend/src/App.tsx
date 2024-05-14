
import './App.css'
import "leaflet/dist/leaflet.css";
import '@mantine/core/styles.css';
import { MantineProvider } from '@mantine/core';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import AppContextProvider from './context/AppContext';
import Home from './views/Home';
import { MantineLightTheme } from './ui/MantineLightTheme';


const queryclient = new QueryClient

function App() {

  return (
    <>
      <MantineProvider theme={MantineLightTheme}>
        <AppContextProvider>
          <QueryClientProvider client={queryclient}>
          <Home></Home>
          </QueryClientProvider>
        </AppContextProvider>
      </MantineProvider>

    </>
  )
}

export default App
