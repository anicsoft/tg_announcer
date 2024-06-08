import React, { useContext, useEffect } from 'react'
import { Avatar, Flex, SegmentedControl, Text, useMantineTheme } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { AppContext } from '../context/AppContext';
import { useGeolocation } from "./../hooks/useGeolocation";
import { IconCurrentLocation } from '@tabler/icons-react';
import HeartBadge from '../ui/HeartBadge';

export default function Menu() {

  const { userData } = useContext(AppContext)
  console.log(userData);

  const { viewType, setViewType } = useContext(AppContext)

  const theme = useMantineTheme()

  return (
    <Flex
      mih={50}
      gap="sm"
      justify="space-between"
      align="center"
      direction="row"
      wrap="nowrap"
      mx="sm"
      my="xs"
    >
      <Flex
        gap="sm"
        justify="flex-start"
        align="center"
        direction="row"
        wrap="nowrap">
        <div style={{ position: "relative" }}>
          <Avatar src={null} alt={`${userData?.firstName} ${userData?.lastName}`} color="red">{`${userData?.firstName ? userData.firstName[0].toUpperCase() : ''}${userData?.lastName ? userData.lastName[0].toUpperCase() : ''}`}</Avatar>
          <HeartBadge count={4} style={{ position: "absolute", bottom: "-25%", right: "-25%", display: "flex" }}></HeartBadge>

        </div>
        <Flex
          align="start"
          direction="column">
          <Text>{`${userData?.firstName} ${userData?.lastName}`}</Text>
          <Flex
            gap={5}
            align="center"
            direction="row">
          </Flex>

        </Flex>
      </Flex>

      <SegmentedControl
        radius={"xl"}
        value={viewType}
        onChange={setViewType}
        withItemsBorders={false}
        color={theme.primaryColor}
        data={[{ value: 'map', label: 'Map' }, { value: 'list', label: 'List' }, { value: 'admin', label: 'Admin' }]} />
    </Flex>
  )
}
