import { Avatar, Flex, Text } from '@mantine/core'
import React from 'react'
import { CardProps } from '../utils/data'

export default function OfferThumbnail({offer}: {offer: CardProps}) {
  return (
    <Flex
      mih={50}
      bg="rgba(0, 0, 0, .3)"
      gap="md"
      justify="flex-start"
      align="flex-start"
      direction="row"
      wrap="wrap"
    >
      <Avatar variant="filled" radius="sm" size="xl" src={offer.logo} />;
      <Text>{offer.title}</Text>
      <Text>{offer.businessName}, {offer.address}</Text>
      <Text>24.05 10:30-12:30</Text>
      <Text>{offer.title}</Text>
    </Flex>
  )
}
