import {
  Card,
  Button,
  TypographyStylesProvider,
  Text,
  Group,
  Modal,
  Image,
  Stack,
  Title,
  useMantineTheme,
  ActionIcon,
  Avatar
} from '@mantine/core';
import { CardProps } from '../utils/data';
import { useDisclosure } from '@mantine/hooks';
import { IconCalendar, IconHeart, IconHeartFilled, IconInfoCircle, IconInfoSmall, IconInfoSquare, IconMapPin } from '@tabler/icons-react';

export default function OfferCard({ popUp }: { popUp: CardProps }) {

  const theme = useMantineTheme()
  const [opened, { open, close }] = useDisclosure(false);
  console.log(popUp);

  return (
    <Stack gap={0}>
      <Image
        // radius="md"
        // h={100}

        src="src/assets/zhraka.webp"
        style={{ aspectRatio: 2, borderBottomRightRadius: theme.radius.lg, borderBottomLeftRadius: theme.radius.lg, boxShadow: theme.shadows.md }}
      />
      <Stack px="xs" pb="md" pt="3.5rem" style={{ position: "relative", boxSizing: "border-box" }}>
        <Group
          style={{
            // right: 0,
            left: 'var(--mantine-spacing-xs)',
            top: 'calc(0px - var(--mantine-spacing-lg))',
            margin: "auto",
            zIndex: 300,
            position: "absolute",
          }}>
          <Avatar radius="sm" size="lg" src={`src/assets/cards_thumbnails/dummy_logo.webp`}
          // style={{
          // // right: 0,
          // left: 0,
          // top: 'calc(0px - var(--avatar-size-lg) / 2)',
          // margin: "auto",
          // zIndex: 300,
          // position: "absolute",
          // }}
          />
          <Group justify='space-between' align='flex-end' mt="xl" gap={4}>
            <Text ta="left" size='md'>{popUp?.companyData?.name}</Text>
            {/* <Button variant="subtle" size="xs" rightSection={<IconInfoSquare stroke={2} />}></Button> */}
            <ActionIcon variant="transparent" aria-label="Company info">
              <IconInfoCircle stroke={2} />
            </ActionIcon>
            {/* <ActionIcon color="red" variant="transparent" aria-label="Favorites">
              <IconHeart stroke={2} />
            </ActionIcon> */}
          </Group>

        </Group>
        <Stack gap={5}>
          <Title ta="left" >{popUp.title}</Title>
          <Group justify='space-between' align='flex-end' gap={4}>
            <Text ta="left" size='md'>{popUp?.companyData?.name}</Text>
            {/* <Button variant="subtle" size="xs" rightSection={<IconInfoSquare stroke={2} />}></Button> */}
            <ActionIcon variant="transparent" aria-label="Company info">
              <IconInfoCircle stroke={2} />
            </ActionIcon>
            {/* <ActionIcon color="red" variant="transparent" aria-label="Favorites">
              <IconHeart stroke={2} />
            </ActionIcon> */}
          </Group>
          <Group gap={4} >
            {/* <IconMapPin stroke={1} />
            <ActionIcon variant="transparent" aria-label="Company info">
              <IconMapPin stroke={1} />
            </ActionIcon>
            <Text ta="left" size='md'>{popUp?.companyData?.address}</Text> */}
            <Button variant="subtle" size="xs" p={0} leftSection={<IconMapPin stroke={1} />}>{popUp?.companyData?.address}</Button>
          </Group>
          <Group gap={4}>
            <IconCalendar stroke={1} />
            <Text size='sm'>{popUp.start_date_time}</Text>
          </Group>
          {popUp?.promoCode &&
            <Text ta="left" size='sm'>Promo code: {popUp.promoCode}</Text>
          }
        </Stack>
        <Group justify="space-between" mt="xs" mb="xs">
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
        </Group>
        {/* <div>
          <Button variant="outline" color="blue" size="xs" fullWidth={false}>
            Show on map
          </Button>

        </div> */}
      </Stack>

    </Stack >
  );
}

