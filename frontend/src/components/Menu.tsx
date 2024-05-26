import React, { useContext, useEffect } from 'react'
import { Avatar, Flex, SegmentedControl, Text } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { AppContext } from '../context/AppContext';
import { useGeolocation } from "./../hooks/useGeolocation";
import { IconCurrentLocation } from '@tabler/icons-react';

export default function Menu() {

  const { userData } = useContext(AppContext)
  console.log(userData);
  
  const {viewType, setViewType} = useContext(AppContext)
  useEffect(() => {
    console.log('View type is ' + viewType);
    
  }, [viewType])  

  const { latitude, longitude, error } = useGeolocation();
  console.log(latitude);
  
  return (
    <Flex
      mih={50}
      // bg="rgba(0, 0, 0, .3)"
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
        <Avatar src={null} alt={`${userData?.firstName} ${userData?.lastName}`} color="red">{`${userData?.firstName ? userData.firstName[0].toUpperCase():''} ${userData?.lastName ? userData.lastName[0].toUpperCase() :''}`}</Avatar>
        <Flex
        // gap="xs"
        // justify="flex-start"
        align="start"
        direction="column">
          <Text>{`${userData?.firstName} ${userData?.lastName}`}</Text>
          <Flex
            gap={5}
            // justify="flex-start"
            align="center"
            direction="row">
        <Text size='sm'>lat:{latitude}, lon:{longitude}</Text>
            <IconCurrentLocation size={14}></IconCurrentLocation>
        </Flex>

        </Flex>
      </Flex>

      <SegmentedControl
        radius={"xl"}
        value={viewType}
        onChange={setViewType}
        data={[{ value: 'map', label: 'Map' }, { value: 'list', label: 'List' }]} />
    </Flex>
  )
}
