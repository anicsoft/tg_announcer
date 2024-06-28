import { Stack, Group, Avatar, Title, Text, Button, TypographyStylesProvider } from '@mantine/core'
import { IconMapPin, IconCalendar, IconMail, IconPhone } from '@tabler/icons-react'
import React from 'react'

export default function CompanyCard() {

  const data = {
    "company_id": "f8afcc36-45a1-4cae-bf60-308f3c405896",
    "name": "Reval Cafe",
    "description": "Electronics store in Tallinn",
    "address": "Müürivahe 14, Tallinn, Estonia",
    "logo_url": "https://tgannouncer.s3.amazonaws.com/uploads/company_5.png",
    "latitude": 59.436877,
    "longitude": 24.749122,
    "phone": "56355555",
    "email": "reval.cafe@example.com",
    "opening_start_time": "08:00",
    "opening_end_time": "22:00",
    "category": [
      "Shops",
      "Shops"
    ]
  }
  return (

    <Stack gap={0} h={"100vh"}>
      {/* <Image
        // src={popUp?.picture_url ?? "src/assets/zhraka.webp"}
        src="src/assets/zhraka.webp"
        style={{ aspectRatio: 2, borderBottomRightRadius: theme.radius.lg, borderBottomLeftRadius: theme.radius.lg, boxShadow: theme.shadows.md }}
      /> */}
      <Stack px="xs" pb="md" >
        <Group justify='space-between' align='flex-end' mt="xl" gap={4}>

          <Title ta="left" >{data.name}</Title>
          {/* <Button variant="subtle" size="xs" rightSection={<IconInfoSquare stroke={2} />}></Button> */}
          {/* <ActionIcon variant="transparent" aria-label="Company info">
            <IconInfoCircle stroke={2} />
          </ActionIcon> */}
          {/* <ActionIcon color="red" variant="transparent" aria-label="Favorites">
            <IconHeart stroke={2} />
          </ActionIcon> */}
          <Avatar radius="sm" size="lg" src={data?.logo_url}
            style={{
              boxShadow: "rgba(0, 0, 0, 0.1) 0px 4px 12px"
            }}
          />
        </Group>
        <Stack gap={5}>
          <Text ta="left" size='md'>{data?.name}</Text>
          {/* <Group justify='space-between' align='flex-end' gap={4}>
            <Text ta="left" size='md'>{popUp?.companyData?.name}</Text>
            <ActionIcon variant="transparent" aria-label="Company info">
              <IconInfoCircle stroke={2} />
            </ActionIcon>
          </Group> */}
          <Group gap={4} >
            {/* <IconMapPin stroke={1} />
            <ActionIcon variant="transparent" aria-label="Company info">
              <IconMapPin stroke={1} />
            </ActionIcon>
            <Text ta="left" size='md'>{popUp?.companyData?.address}</Text> */}
            <IconMapPin stroke={1} />
            <Text size='sm'>{data?.address}</Text>
          </Group>

          <Group gap={4} >
            <IconMapPin stroke={1} />
            <Text size='sm'>Opening hours: {data?.opening_start_time}-{data?.opening_end_time}</Text>
          </Group>
          <Group gap={4}>
            <IconPhone stroke={1} />
            <Text size='sm'>{data.phone}</Text>
          </Group>

          <Group gap={4}>
            <IconMail stroke={1} />
            <Text size='sm'>{data.email}</Text>
          </Group>
        </Stack>
        {/* <Group justify="space-between" mt="xs" mb="xs">
          <Text size="sm" c="dimmed">
            <TypographyStylesProvider component={'div'}>
              <p className='offerUl'
                dangerouslySetInnerHTML={{
                  __html: popUp?.content ?? `<p>This is an example of an HTML document with bullet points and bold text.</p>
              <ul>
                  <li><b>Lorem ipsum dolor sit amet</b>, consectetur adipiscing elit. Integer nec odio. Praesent libero. Sed cursus ante dapibus diam.</li>
                  <li><b>Sed nisi</b>. Nulla quis sem at nibh elementum imperdiet. Duis sagittis ipsum. Praesent mauris.</li>
                  <li><b>Fusce nec tellus</b> sed augue semper porta. Mauris massa. Vestibulum lacinia arcu eget nulla.</li>
                  <li><b>Class aptent taciti sociosqu</b> ad litora torquent per conubia nostra, per inceptos himenaeos.</li>
                  <li><b>Curabitur sodales ligula</b> in libero. Sed dignissim lacinia nunc. Curabitur tortor. Pellentesque nibh.</li>
              </ul>` }}
              />
            </TypographyStylesProvider>
          </Text>
        </Group> */}
        {/* <div>
          <Button variant="outline" color="blue" size="xs" fullWidth={false}>
            Show on map
          </Button>

        </div> */}
      </Stack>

    </Stack >
  )
}
