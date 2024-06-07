import { AppShell } from '@mantine/core'
import React, { useContext } from 'react'
import FiltersDrawer from '../components/FiltersDrawer'
import { AppContext } from '../context/AppContext';
import Menu from '../components/Menu';
import BasicMap from './BasicMap';
import OffersListView from './OffersListView';
import { useQueries, useQuery } from '@tanstack/react-query';
import { mock_cards } from '../utils/data';
import AdminConsole from './AdminConsole';

export default function Home() {
  const { viewType } = useContext(AppContext);

  const { data } = useQuery({
    queryKey: ['offers'],
    queryFn: () =>
      fetch('http://localhost:8888/announcements').then((res) =>
        res.json(),
      )
  })

  console.log('DATA');
  console.log(data);

  const queries = useQueries({
    queries: data?.announcements
      ? data.announcements.map(offer => {
        return {
          queryKey: ['companies', offer.company_id],
          queryFn: () => fetch(`http://localhost:8888/companies/${offer.company_id}`).then((res) =>
            res.json(),
          ),
        };
      })
      : [], // if data?.announcements is undefined, an empty array will be returned
    combine: (results) => {
      const companies = results.reduce((acc, item) => {
        acc[item.data?.company_id] = item.data;
        return acc;
      }, {})

      return {
        res: data?.announcements.map(offer => { return { ...offer, companyData: companies[offer.company_id] } }),
        // res: companies
      }
    },
  },
  );

  // console.log(data);
  // console.log(queries);


  // if (isPending) return 'Loading...'

  // if (error) return 'An error has occurred: ' + error.message

  // const data = mock_cards

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
          <><BasicMap offers={queries?.res}></BasicMap><FiltersDrawer></FiltersDrawer></>
        }
        {viewType === 'list' &&
          <><OffersListView offers={queries?.res}></OffersListView><FiltersDrawer></FiltersDrawer></>
        }
        {viewType === 'admin' &&
          <AdminConsole></AdminConsole>
        }
      </AppShell.Main>
    </AppShell>
  )
}
