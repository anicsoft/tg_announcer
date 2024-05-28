import {
  Card,
  Button,
  TypographyStylesProvider,
  Text,
  Group,
  Modal
} from '@mantine/core';
import { CardProps } from '../utils/data';
import { useDisclosure } from '@mantine/hooks';

export default function OfferCard({popUp}: {popUp: CardProps}) {

  const [opened, { open, close }] = useDisclosure(false);
  console.log(popUp);
  
  return (
    <Card shadow="md" withBorder radius="md">
      <Card.Section>{popUp.title}</Card.Section>
      <Group justify="space-between" mt="md" mb="xs">
      <Text size="sm" c="dimmed">
      <TypographyStylesProvider>
        <div
          dangerouslySetInnerHTML={{ __html: popUp?.content ?? "LALALA" }}
          />
      </TypographyStylesProvider>
      </Text>
      </Group>
      <Card.Section>
        <Button variant="outline" color="blue" size="sm" onClick={open}>
          See More
          {/* <Modal opened={opened} onClose={close} title="Authentication">

          </Modal> */}
          {/* <UnstyledButton key={card?.title + card?.businessName} onClick={open}>
        </UnstyledButton> */}
        </Button>
        <Modal opened={opened} onClose={close} title="Authentication">
              <OfferCard popUp={popUp}></OfferCard>
        </Modal>
      </Card.Section>
    </Card>
  );
}

