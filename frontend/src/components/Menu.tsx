import { useContext } from 'react'
import { Avatar, Flex, SegmentedControl, Text, useMantineTheme } from '@mantine/core';
import { AppContext } from '../context/AppContext';

export default function Menu() {

  const { userData } = useContext(AppContext)

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
        wrap="nowrap"
        onClick={()=>(setViewType("profile"))}
        style={{cursor: "pointer"}}
        >
          
        <div style={{ position: "relative" }}>
          <Avatar src={null} alt={`${userData?.firstName} ${userData?.lastName}`} color="red">{`${userData?.firstName ? userData.firstName[0].toUpperCase() : ''}${userData?.lastName ? userData.lastName[0].toUpperCase() : ''}`}</Avatar>

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
