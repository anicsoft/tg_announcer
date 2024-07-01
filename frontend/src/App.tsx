
import './App.css'
import "leaflet/dist/leaflet.css";

import '@mantine/core/styles/global.css';
import '@mantine/core/styles.css';
import '@mantine/tiptap/styles.css';
import '@mantine/dates/styles.css';

// import '@mantine/core/styles/AppShell.css';
// import '@mantine/core/styles/ScrollArea.css';
// import '@mantine/core/styles/UnstyledButton.css';
// import '@mantine/core/styles/VisuallyHidden.css';
// import '@mantine/core/styles/Paper.css';
// import '@mantine/core/styles/Popover.css';
// import '@mantine/core/styles/CloseButton.css';
// import '@mantine/core/styles/Group.css';
// import '@mantine/core/styles/Loader.css';
import '@mantine/core/styles/Overlay.css';
import '@mantine/core/styles/ModalBase.css';
import '@mantine/core/styles/Drawer.css';
// import '@mantine/core/styles/Input.css';
// import '@mantine/core/styles/Flex.css';
// import '@mantine/core/styles/Switch.css';
// import '@mantine/core/styles/SegmentedControl.css';
// import '@mantine/core/styles/Button.css';

import { DEFAULT_THEME, MantineProvider, mergeMantineTheme } from '@mantine/core';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import AppContextProvider from './context/AppContext';
import Home from './views/Home';
import { AnicLightTheme } from './utils/mantineTheme';
import { useColorScheme } from '@mantine/hooks';
import { retrieveLaunchParams } from "@tma.js/sdk";

const queryclient = new QueryClient

function App() {
  let initDataRaw, initData;

  // Check if retrieveLaunchParams is defined and call it if available
  if (typeof retrieveLaunchParams === 'function') {
    try {
      const params = retrieveLaunchParams();
      initDataRaw = params.initDataRaw;
      initData = params.initData;
      console.log('init data', initDataRaw);
    } catch (error) {
      console.error('Error retrieving launch params:', error);
    }
  } else {
    console.warn('retrieveLaunchParams function is not available');
  }
  const theme = mergeMantineTheme(DEFAULT_THEME, AnicLightTheme);

  const colorScheme = useColorScheme();
  const theme_color = colorScheme == 'dark' ? 'light' : 'light';
  // if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
  //   theme_color = "light";
  // }

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
