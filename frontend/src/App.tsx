
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
  let theme_color: "light" | "dark" | undefined = "light";
  if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
    theme_color = "light";
  }

  return (
    <>
      <MantineProvider theme={theme} defaultColorScheme={theme_color} forceColorScheme={theme_color}>
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
