import { Avatar, Badge, Button, Card, Flex, Group, Modal, Stack, Text, Title, useMantineTheme } from '@mantine/core'
import React from 'react'
import { CardProps } from '../utils/data'
import OfferCard from './OfferCard'
import { useDisclosure } from '@mantine/hooks';
import { useGeolocation } from '../hooks/useGeolocation';
import { IconWalk } from '@tabler/icons-react';

export default function OfferThumbnail({ offer }: { offer: CardProps }) {
  const theme = useMantineTheme();
  const [opened, { open, close }] = useDisclosure(false);
  console.log(offer);
  const { latitude, longitude, error } = useGeolocation();
  const toRadians = (degrees) => degrees * (Math.PI / 180);

  const distanceKm = (lat1, lon1, lat2, lon2) => {

    const distanceKm = Math.acos(
      Math.sin(toRadians(lat1)) * Math.sin(toRadians(lat2)) +
      Math.cos(toRadians(lat1)) * Math.cos(toRadians(lat2)) *
      Math.cos(toRadians(lon2) - toRadians(lon1))
    ) * 6371;
    let res = `${distanceKm.toFixed(2).toString()}km`
    if (distanceKm < 1) {
      res = `${(distanceKm * 1000).toFixed(2).toString()}m`;
    }
    return res
  }


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
          wrap="nowrap"
        >
          <Avatar my="xs" radius="sm" size="md" src={`src/assets/cards_thumbnails/dummy_logo.webp`} />
          <Stack mt="xs" align='start' gap={5} flex={1}>
            <Flex gap="sm"
              justify="space-between"
              align="baseline"
              w={"100%"}
              direction="row">
              <Title m="xs" size='md' order={2} ta="left">{offer.title}</Title>
              <Group gap={2} wrap='nowrap' align='center'>
                <Text m="xs" size='xs'>{distanceKm(latitude, longitude, offer?.companyData?.latitude, offer?.companyData?.longitude)}</Text>
                <IconWalk size={18}></IconWalk>
              </Group>

            </Flex>
            {/* <Text ta="left" size='xs'>{offer?.companyData?.name}, {offer?.companyData?.address}</Text>
            <Text size='xs'>24.05 10:30-12:30</Text> */}
            {/* <Text>{offer.title}</Text> */}

          </Stack>
        </Flex>
      </Card.Section>
      <Card.Section px="sm" pb="xs">
        <Text ta="left" size='xs'>{offer?.companyData?.name}, {offer?.companyData?.address}</Text>
        <Text ta="left" size='xs'>24.05 10:30-12:30</Text>
        <Flex justify="space-between"
          pt="xs"
          align="flex-start"
          direction="row"
          wrap={"nowrap"}
          gap={8}>
          <Group flex={1} gap={4}>
            {offer.categories?.map((category: string) =>
              <Badge key={offer.announcement_id + category} size="xs" variant="light" color={theme.colors.orange[9]} gradient={{ from: 'lightOrange', to: 'red', deg: 340 }}>
                {category}
              </Badge>
            )}
          </Group>
          <Button autoContrast size="xs" radius="sm" px={8} py={4} onClick={open} color={theme.primaryColor}>
            See more
          </Button>
        </Flex>
      </Card.Section>
      <Modal opened={opened} onClose={close} title={offer.title}>
        <OfferCard popUp={offer}></OfferCard>
      </Modal>
    </Card>
  )
}
