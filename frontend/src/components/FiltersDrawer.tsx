import { BackgroundImage, Box, Chip, Container, Drawer, Group } from '@mantine/core'
import React, { useContext, useState } from 'react'
import { AppContext } from '../context/AppContext';
import MapFilterButton from '../ui/MapFilterButton';
import { useQuery } from '@tanstack/react-query';

export default function FiltersDrawer() {
  const { filterDrawerOpened, filterDrawerHandlers } = useContext(AppContext);

  console.log(`In drawer it is ${filterDrawerOpened}`);
  // const categories = ['category1', 'category2', 'category3', 'categ', 'categorycategory', 'cate']
  // const offers = [
  //   "Discount Offers",
  //   "Buy One Get One (BOGO)",
  //   "Bundle Deals",
  //   "Free Shipping",
  //   "Loyalty Program Offers",
  //   "Flash Sales",
  //   "Clearance Sales",
  //   "Seasonal Promotions",
  //   "Referral Offers",
  //   "Cashback Offers",
  //   "Rebate Offers",
  //   "Gift with Purchase",
  //   "Early Bird Offers",
  //   "VIP/Exclusive Offers",
  //   "Membership Discounts",
  //   "Quantity Discounts",
  //   "Trade-In Offers",
  //   "First-Time Customer Offers",
  //   "Volume-Based Discounts",
  //   "Competitions or Sweepstakes"
  // ]

  const { data: businessCategories } = useQuery({
    queryKey: ['categories'],
    queryFn: () =>
      fetch('http://localhost:8888/backend/categories/business').then((res) =>
        res.json()
      )
  })
  const { data: offerCategories } = useQuery({
    queryKey: ['offerCategories'],
    queryFn: () =>
      fetch('http://localhost:8888/backend/categories/offer').then((res) =>
        res.json(),
      )
  })

  console.log(businessCategories);
  console.log(offerCategories);


  const [categoriesValue, setCategoriesValue] = useState(businessCategories);
  const [offersValue, setOffersValue] = useState(offerCategories);


  return (
    <>
      {businessCategories && offerCategories && categoriesValue && offersValue &&
        <Box pos="relative">

          <MapFilterButton></MapFilterButton>
          <Drawer position="bottom" radius="md" size="50vh" withCloseButton={false} opened={filterDrawerOpened} onClose={filterDrawerHandlers.close} transitionProps={{ timingFunction: "ease", duration: 200 }} style={{ body: { backgroundColor: "red" } }}>
            <Drawer.Title mb="md">Filters</Drawer.Title>
            <Container px={0}>
              Categories
              <Chip.Group multiple value={categoriesValue} onChange={setCategoriesValue}>
                <Group gap="xs" px={0} justify="start" mt="sm">
                  {businessCategories && businessCategories?.map(cat =>
                    <Chip key={cat.name}
                      // icon={<IconX style={{ width: rem(16), height: rem(16) }} />}
                      value={cat.name}
                      color="red"
                      variant="filled"
                    >
                      {cat.name}
                    </Chip>
                  )}
                  <Chip key='catAll'
                    // icon={<IconX style={{ width: rem(16), height: rem(16) }} />}
                    value='All'
                    color="red"

                    checked={categoriesValue.length == businessCategories.length}
                    variant="filled"
                    onChange={() => setCategoriesValue(categoriesValue.length < businessCategories.length || categoriesValue.length == 0 ? businessCategories : [])}
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
                  {offerCategories && offerCategories?.map(offer =>
                    <Chip key={offer.name}
                      value={offer.name}
                    >
                      {offer.name}
                    </Chip>
                  )}
                  <Chip key='offAll'
                    value='All'
                    checked={offersValue.length == offerCategories.length}
                    onChange={() => setOffersValue(offersValue.length < offerCategories.length || offersValue.length == 0 ? offerCategories : [])}
                  >
                    All
                  </Chip>
                </Group>
              </Chip.Group>
            </Container>
          </Drawer>
        </Box>
      }
    </>

  )
}
