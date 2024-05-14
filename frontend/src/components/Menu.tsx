import React, { useContext, useEffect } from 'react'
import { Avatar, Flex, SegmentedControl } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { AppContext } from '../context/AppContext';

export default function Menu() {


  const {viewType, setViewType} = useContext(AppContext)
  useEffect(() => {
    console.log('View type is ' + viewType);
    
  }, [viewType])
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
        <Avatar src={null} alt="Vitaly Rtishchev" color="red">VR</Avatar>
        <Flex
        // gap="xs"
        // justify="flex-start"
        align="start"
        direction="column">
        <div>LALALALA</div>
        <div>bebeb</div>

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
