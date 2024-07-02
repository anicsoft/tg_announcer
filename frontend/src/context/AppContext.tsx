import { useDisclosure } from '@mantine/hooks';
import React, { ReactNode, createContext, useState } from 'react';

import { retrieveLaunchParams } from '@tma.js/sdk';
import { User } from '@tma.js/sdk-react';
import { parseLaunchParams, launchParamsParser } from '@tma.js/sdk';

interface BaseState {
  userData: User | undefined,
  viewType: string;
  filterDrawerOpened: boolean;
  latitude: number;
  longitude: number;
  setViewType: (x: string) => void;
  filterDrawerHandlers: {
    open: () => void;
    close: () => void;
    toggle: () => void;
  };
  setLatitude: (x: number) => void
  setLongitude: (x: number) => void
  initDataRaw: string;
  
}

export const AppContext = createContext<BaseState>({
  userData: undefined,
  viewType: 'map',
  filterDrawerOpened: false,
  latitude: 0,
  longitude: 0,
  setViewType: () => { },
  filterDrawerHandlers: {
    open: () => { },
    close: () => { },
    toggle: () => { }
  },
  setLatitude: () => { },
  setLongitude: () => { },
  initDataRaw: "",
});

function AppContextProvider({ children }: { children: ReactNode }) {
  let userData;
  try {
    const { initData } = retrieveLaunchParams();
    userData = initData?.user;
    
  } catch (error) {
    console.log('Error');
    userData = { firstName: "Jane", lastName: "Doe" };
  }

  const [latitude, setLatitude] = useState<number>(0);
  const [longitude, setLongitude] = useState<number>(0);
  const [viewType, setViewType] = useState<string>('map');
  const { initDataRaw, initData } = retrieveLaunchParams();

  const [filterDrawerOpened, filterDrawerHandlers] = useDisclosure(false, {
    onOpen: () => console.log('Opened'),
    onClose: () => console.log('Closed'),
  });

  const values = {
    userData,
    viewType,
    filterDrawerOpened,
    setViewType,
    filterDrawerHandlers,
    setLatitude,
    setLongitude,
    initData,
    initDataRaw
  }

  return <AppContext.Provider value={values}>{children}</AppContext.Provider>;
}

export default AppContextProvider;