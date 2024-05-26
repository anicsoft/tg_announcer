import { Modal, Stack, UnstyledButton } from '@mantine/core'
import React from 'react'

import {CardProps, mock_cards} from './../utils/data';
import OfferThumbnail from '../components/OfferThumbnail';

export default function OffersListView() {
  // const =  mock_cards
  return (
    <Stack py={"2rem"} px={"1rem"}>
      {mock_cards.map((card:CardProps) => {
        return <OfferThumbnail key={card.businessName +card.title} offer={card} ></OfferThumbnail>
      })}
    </Stack>
  )
}
