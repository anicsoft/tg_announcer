import { Box, Chip, Container, Drawer, Group } from '@mantine/core'
import React, { useContext, useState } from 'react'
import {AppContext} from '../context/AppContext';
import MapFilterButton from '../ui/MapFilterButton';

export default function FiltersDrawer() {
  const { filterDrawerOpened, filterDrawerHandlers } = useContext(AppContext);

  console.log(`In drawer it is ${filterDrawerOpened}`);
  const categories = ['category1', 'category2', 'category3', 'categ', 'categorycategory', 'cate']
  const offers = [
    "Discount Offers",
    "Buy One Get One (BOGO)",
    "Bundle Deals",
    "Free Shipping",
    "Loyalty Program Offers",
    "Flash Sales",
    "Clearance Sales",
    "Seasonal Promotions",
    "Referral Offers",
    "Cashback Offers",
    "Rebate Offers",
    "Gift with Purchase",
    "Early Bird Offers",
    "VIP/Exclusive Offers",
    "Membership Discounts",
    "Quantity Discounts",
    "Trade-In Offers",
    "First-Time Customer Offers",
    "Volume-Based Discounts",
    "Competitions or Sweepstakes"
]

  const [categoriesValue, setCategoriesValue] = useState(categories);
  const [offersValue, setOffersValue] = useState(offers);

  
  return (
    <Box pos="relative">
      
    <MapFilterButton></MapFilterButton>
      <Drawer position="bottom" radius="md" size="50vh"  withCloseButton={false} opened={filterDrawerOpened} onClose={filterDrawerHandlers.close} transitionProps={{timingFunction:"ease", duration: 200 }} >
      <Drawer.Title mb="md">Filters</Drawer.Title>
        <Container px={0}>
          Categories
          <Chip.Group multiple value={categoriesValue} onChange={setCategoriesValue}>
          <Group gap="xs" px={0} justify="start" mt="sm">
            {categories.map(cat => 
              <Chip key={cat}
              // icon={<IconX style={{ width: rem(16), height: rem(16) }} />}
              value={cat}
              color="red"
              variant="filled"
              >
              {cat}
            </Chip>
              )}
            <Chip key='catAll'
              // icon={<IconX style={{ width: rem(16), height: rem(16) }} />}
              value='All'
                color="red"
                
              checked={categoriesValue.length == categories.length}
                variant="filled"
                onChange={() => setCategoriesValue(categoriesValue.length < categories.length || categoriesValue.length == 0 ? categories : [])}
              >
              All
            </Chip>
            </Group>
          </Chip.Group>
        </Container>
        <Container px={0} my="sm">
          Offers
          <Chip.Group multiple value={offersValue} onChange={setOffersValue}>
          <Group gap="xs" px={0} justify="start" mt="xs">
            {offers.map(offer => 
              <Chip key={offer}
                value={offer}
              >
              {offer}
            </Chip>
              )}
              <Chip key='offAll'
              value='All'
              checked={offersValue.length == offers.length}
                onChange={() => setOffersValue(offersValue.length < offers.length || offersValue.length == 0 ? offers : [])}
              >
              All
            </Chip>
              </Group>
          </Chip.Group>
      </Container>
      </Drawer>
    </Box>
  )
}
