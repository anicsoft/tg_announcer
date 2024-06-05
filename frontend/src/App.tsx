
import './App.css'
import "leaflet/dist/leaflet.css";
import '@mantine/core/styles.css';
import '@mantine/tiptap/styles.css';
import '@mantine/dates/styles.css';


import { DEFAULT_THEME, MantineProvider, mergeMantineTheme } from '@mantine/core';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import AppContextProvider from './context/AppContext';
import Home from './views/Home';
import { AnicLightTheme } from './utils/mantineTheme';


const queryclient = new QueryClient

function App() {

  const theme = mergeMantineTheme(DEFAULT_THEME, AnicLightTheme);
  return (
    <>
      <MantineProvider theme={theme} defaultColorScheme='dark' forceColorScheme='dark'>
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
