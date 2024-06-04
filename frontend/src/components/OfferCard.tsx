import {
  Card,
  Button,
  TypographyStylesProvider,
  Text,
  Group,
  Modal,
  Stack
} from '@mantine/core';
import { CardProps } from '../utils/data';
import { useDisclosure } from '@mantine/hooks';

export default function OfferCard({ popUp }: { popUp: CardProps }) {


  const [opened, { open, close }] = useDisclosure(false);
  console.log(popUp);

  return (
    <Stack>
      <Stack gap={5}>
        <Text size='xs'>24.05 10:30-12:30</Text>
        <Text ta="left" size='md'>{popUp.businessName}</Text>
        <Text ta="left" size='md'>{popUp.address}</Text>
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
      <div>
        <Button variant="outline" color="blue" size="xs" fullWidth={false}>
          Show on map
        </Button>

      </div>
    </Stack>
  );
}

