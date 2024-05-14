import { Stack } from '@mantine/core'
import React from 'react'

import {CardProps, mock_cards} from '../utils/data';
import OfferCard from '../components/OfferCard';

export default function OffersListView() {

  // const =  mock_cards
  return (
    <Stack py={"2rem"} px={"1rem"}>
      {mock_cards.map((card:CardProps) => {
        // <div style={{height:"100vh", backgroundColor:"pink"}}>OffersListView</div>
        return <OfferCard popUp={card}></OfferCard>
      })}
    </Stack>
  )
}
