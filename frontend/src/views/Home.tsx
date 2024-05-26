import { AppShell } from '@mantine/core'
import React, { useContext } from 'react'
import FiltersDrawer from '../components/FiltersDrawer'
import { AppContext } from '../context/AppContext';
import Menu from '../components/Menu';
import BasicMap from './BasicMap';
import OffersListView from './OffersListView';
import { useQuery } from '@tanstack/react-query';

export default function Home() {
  const { viewType } = useContext(AppContext);

  const { isPending, error, data } = useQuery({
    queryKey: ['repoData'],
    queryFn: () =>
      fetch('http://localhost:8080/categories/business').then((res) =>
        res.json(),
      ),
  })

  if (isPending) return 'Loading...'

  if (error) return 'An error has occurred: ' + error.message

  return (
    <AppShell
      header={{ height: 80, offset: true }}
    >
      <AppShell.Header
        style={{ alignContent: "center" }}
      >
        <Menu></Menu>

      </AppShell.Header>

      <AppShell.Main>
        {viewType === 'map' &&
          <BasicMap data={data}></BasicMap>
        }
        {viewType === 'list' &&
          <OffersListView data={data}></OffersListView>
        }
        <FiltersDrawer></FiltersDrawer>
      </AppShell.Main>
    </AppShell>
  )
}
