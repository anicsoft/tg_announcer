import { Modal, Stack, UnstyledButton } from '@mantine/core'
import React from 'react'

import { CardProps, Company, mock_cards } from './../utils/data';
import OfferThumbnail from '../components/OfferThumbnail';

export default function OffersListView({ offers }: { offers: CardProps[] }) {
  // const =  mock_cards
  return (
    <Stack py={"2rem"} px={"1rem"}>
      {offers.map((card: CardProps, i: number) => {
        return <OfferThumbnail key={i + card.businessName + card.title} offer={card} ></OfferThumbnail>
      })}
    </Stack>
  )
}
