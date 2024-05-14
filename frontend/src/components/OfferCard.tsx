import {
  Card,
  Button,
  TypographyStylesProvider,
  Text,
  Group
} from '@mantine/core';
import { CardProps } from '../utils/data';

export default function OfferCard({popUp}: {popUp: CardProps}) {

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
        <Button variant="outline" color="blue" size="sm" onClick={popUp.onClick}>
          See More
        </Button>
      </Card.Section>
    </Card>
  );
}

