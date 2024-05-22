
import './App.css'
import "leaflet/dist/leaflet.css";
import '@mantine/core/styles.css';
import { MantineProvider } from '@mantine/core';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import AppContextProvider from './context/AppContext';
import Home from './views/Home';
import { MantineLightTheme } from './ui/MantineLightTheme';

// import { retrieveLaunchParams } from '@tma.js/sdk';

const queryclient = new QueryClient

// if (initDataRaw) {
  //   const data = JSON.parse(initDataRaw)
//   console.log(data);
  
// }
function App() {
  // try {
  //   const { initDataRaw, initData } = retrieveLaunchParams();
  //   console.log('init data ', initData)
    
  // } catch (error) {
  //   console.log(error);
    
  // }

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
