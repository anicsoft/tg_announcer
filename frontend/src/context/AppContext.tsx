import { useDisclosure } from '@mantine/hooks';
import React, { ReactNode, createContext, useState } from 'react';

import { retrieveLaunchParams } from '@tma.js/sdk';
import { User } from '@tma.js/sdk-react';


interface BaseState {
  userData: User | undefined,
  viewType: string;
  filterDrawerOpened: boolean;
  setViewType: (x: string) => void;
  filterDrawerHandlers: {
    open: () => void;
    close: () => void;
    toggle: () => void;
  };
}

export const AppContext = createContext<BaseState>({
  userData: undefined,
  viewType: 'map',
  filterDrawerOpened: false,
  setViewType: () => {},
  filterDrawerHandlers: {
    open: () => {},
    close: () => {},
    toggle: () => {}
  }
});

function AppContextProvider({ children }: { children: ReactNode }) {
  let userData; 
  try {
    const { initData } = retrieveLaunchParams();
    userData = initData?.user;
  } catch (error) {
    console.log('Error');
    
  }
  const [viewType, setViewType] = useState<string>('map');
  const [filterDrawerOpened, filterDrawerHandlers] = useDisclosure(false, {
    onOpen: () => console.log('Opened'),
    onClose: () => console.log('Closed'),
  });

  const values = {
    userData,
    viewType,
    filterDrawerOpened,
    setViewType,
    filterDrawerHandlers
  }

  return <AppContext.Provider value={values}>{children}</AppContext.Provider>;
}

export default AppContextProvider;