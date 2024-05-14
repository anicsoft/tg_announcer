import { ActionIcon } from '@mantine/core'
import { IconAdjustmentsHorizontal, IconX } from '@tabler/icons-react'
import React, { useContext } from 'react'
import { AppContext } from '../context/AppContext';

export default function MapFilterButton() {

  const { filterDrawerOpened, filterDrawerHandlers } = useContext(AppContext);

  return (
    <ActionIcon radius={"xl"}  variant="filled" size="xl" aria-label="Settings" pos="fixed" bottom={filterDrawerOpened ? 'calc(20px + 50vh)' : '1.5rem'} right="1.5rem" style={{zIndex:"201", transition: 'bottom 0.2s ease'}} onClick={filterDrawerHandlers.toggle}>
      
      {filterDrawerOpened ? (
        <IconX style={{ width: '70%', height: '70%' }} stroke={1.5} className='button-icon-transition'/>
        ) : (
          <IconAdjustmentsHorizontal style={{ width: '70%', height: '70%' }} stroke={1.5} className='button-icon-transition'/>
        )}
      
    </ActionIcon>
  )
}
