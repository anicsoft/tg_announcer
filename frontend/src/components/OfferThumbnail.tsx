import { Avatar, Card, Flex, Text } from '@mantine/core'
import React from 'react'
import { CardProps } from '../utils/data'

export default function OfferThumbnail({ offer }: { offer: CardProps }) {
  return (
    <Card withBorder radius="sm">
      <Card.Section inheritPadding>

        <Flex
          mih={50}
          // bg="rgba(0, 0, 0, .3)"
          gap="md"
          justify="flex-start"
          align="flex-start"
          direction="row"
          wrap="wrap"
        >
          <Avatar m="xs" radius="sm" size="xl" src={`src/assets/cards_thumbnails/${offer.logo}`} />
          <Text m="xs" >{offer.title}</Text>
        </Flex>
      </Card.Section>
      <Card.Section>

        <Text>{offer.businessName}, {offer.address}</Text>
        <Text>24.05 10:30-12:30</Text>
        <Text>{offer.title}</Text>
      </Card.Section>
    </Card>
  )
}
