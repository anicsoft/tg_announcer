import { Loader, Modal, Stack, UnstyledButton } from '@mantine/core'
import React from 'react'

import { CardProps } from './../utils/data';
import OfferThumbnail from '../components/OfferThumbnail';

export default function OffersListView({ offers }: { offers: CardProps[] }) {
  // const =  mock_cards
  return (
    <Stack py={"2rem"} px={"1rem"}>
      {offers && offers.length > 0 ? offers.map((card: CardProps, i: number) => {
        return <OfferThumbnail key={i + card.businessName + card.title} offer={card} ></OfferThumbnail>
      }) :
        <Loader></Loader>}
    </Stack>
  )
}
