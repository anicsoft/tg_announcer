import { AppShell } from '@mantine/core'
import React, { useContext } from 'react'
import FiltersDrawer from '../components/FiltersDrawer'
import {AppContext} from '../context/AppContext';
import Menu from '../components/Menu';
import BasicMap from './BasicMap';
import OffersListView from './OffersListView';

export default function Home() {
  const { viewType } = useContext(AppContext);
  
  return (
    <AppShell
    header={{ height: 80, offset: true }}
    >
      <AppShell.Header
        style={{alignContent: "center"}}
      >
        <Menu></Menu>

      </AppShell.Header>

      <AppShell.Main>
        {viewType === 'map' &&
          <BasicMap></BasicMap>
        }
        {viewType === 'list' &&
        <OffersListView></OffersListView>
        }
        <FiltersDrawer></FiltersDrawer>
      </AppShell.Main>
    </AppShell>
  )
}
