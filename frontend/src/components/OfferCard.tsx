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
  Avatar,
} from "@mantine/core";
import { CardProps } from "../utils/data";
import { useDisclosure } from "@mantine/hooks";
import {
  IconCalendar,
  IconHeart,
  IconHeartFilled,
  IconInfoCircle,
  IconInfoSmall,
  IconInfoSquare,
  IconMapPin,
} from "@tabler/icons-react";
import CompanyModal from "../ui/CompanyModal";
import HeartBadge from "../ui/HeartBadge";
import { addFavorite, removeFavorite } from "../shared/api/favorites";
import { useContext, useState, useEffect } from "react";
import { AppContext } from "../context/AppContext";
import { useFetchOffers } from "../shared/api/annoucments";

export default function OfferCard({ popUp }: { popUp: CardProps }) {
  const theme = useMantineTheme();
  const [opened, { open, close }] = useDisclosure(false);
  console.log("pop up", popUp);

  const { userData } = useContext(AppContext);


  const [isFavorite, setIsFavorite] = useState(popUp.company.is_favorite);

  const { data, refetch } = useFetchOffers();


  const userId = userData.id;
  const companyId = popUp.company_id;

  useEffect(() => {
    setIsFavorite(popUp.company.is_favorite); 
  }, [popUp.company.is_favorite]);

  const toggleFavorites = async () => {
    try {
      if (!isFavorite) {
        await addFavorite(userId, companyId);
      } else {
        await removeFavorite(userId, companyId);
      }
      await refetch(); 
    } catch (error) {
      console.error("Error toggling favorites:", error);
    }
  };

  function formatDateRange(start, end) {
    const options = { year: "numeric", month: "short", day: "2-digit" };

    // Parse the dates
    const startDate = new Date(start);
    const endDate = new Date(end);

    // Extract the time components
    const startTime = startDate.toISOString().substr(11, 5); // "HH:MM"
    const endTime = endDate.toISOString().substr(11, 5); // "HH:MM"

    // Compare the dates (ignoring the time part)
    const sameDate = startDate.toDateString() === endDate.toDateString();
    const sameMonth = startDate.getMonth() === endDate.getMonth();
    const sameYear = startDate.getFullYear() === endDate.getFullYear();

    if (sameDate) {
      // Format as "05 Jun 2024 10:00-18:00"
      const formattedDate = startDate.toLocaleDateString("en-GB", options);
      return `${formattedDate} ${startTime}-${endTime}`;
    } else if (sameMonth && sameYear) {
      // Format as "05-10 Jun 2024 10:00-18:00"
      const startFormattedDate = startDate
        .toLocaleDateString("en-GB", options)
        .split(" ")[0];
      const endFormattedDate = endDate.toLocaleDateString("en-GB", options);
      const [_, month, year] = endFormattedDate.split(" ");
      return `${startFormattedDate}-${
        endFormattedDate.split(" ")[0]
      } ${month} ${year} ${startTime}-${endTime}`;
    } else if (sameYear) {
      // Format as "05 Jun - 10 Jul 2024 10:00-18:00"
      const startFormattedDate = startDate
        .toLocaleDateString("en-GB", options)
        .split(" ")
        .slice(0, 2)
        .join(" ");
      const endFormattedDate = endDate
        .toLocaleDateString("en-GB", options)
        .split(" ")
        .slice(0, 2)
        .join(" ");
      const year = startDate.getFullYear();
      return `${startFormattedDate} - ${endFormattedDate} ${year} ${startTime}-${endTime}`;
    } else {
      // Format as "05 Jun 2024 - 10 Jul 2025 10:00-18:00"
      const startFormattedDate = startDate.toLocaleDateString("en-GB", options);
      const endFormattedDate = endDate.toLocaleDateString("en-GB", options);
      return `${startFormattedDate} - ${endFormattedDate} ${startTime}-${endTime}`;
    }
  }

  return (
    <Stack gap={0}>
      <Image
        // src={popUp?.picture_url ?? "src/assets/zhraka.webp"}
        src="src/assets/zhraka.webp"
        style={{
          aspectRatio: 2,
          borderBottomRightRadius: theme.radius.lg,
          borderBottomLeftRadius: theme.radius.lg,
          boxShadow: theme.shadows.md,
        }}
      />
      <Stack
        px="xs"
        pb="md"
        pt="3.5rem"
        style={{ position: "relative", boxSizing: "border-box" }}
      >
        <Group
          style={{
            // right: 0,
            left: "var(--mantine-spacing-xs)",
            top: "calc(0px - var(--mantine-spacing-lg))",
            margin: "auto",
            zIndex: 300,
            position: "absolute",
          }}
        >
          <Avatar
            radius="sm"
            size="lg"
            src={`src/assets/cards_thumbnails/dummy_logo.webp`}
            style={{
              boxShadow: "rgba(0, 0, 0, 0.1) 0px 4px 12px",
            }}
          />

          <Group justify="space-between" align="flex-end" mt="xl" gap={4}>
            <Text ta="left" size="md">
              {popUp?.companyData?.name}
            </Text>
            {/* <Button variant="subtle" size="xs" rightSection={<IconInfoSquare stroke={2} />}></Button> */}
            <ActionIcon
              variant="transparent"
              aria-label="Company info"
              onClick={open}
            >
              <IconInfoCircle stroke={2} />
            </ActionIcon>
            <ActionIcon variant="transparent" onClick={toggleFavorites}>
              <HeartBadge
                color={isFavorite ? "#9F003E" : "#fff"}
                style={{ cursor: "pointer" }}
              />
            </ActionIcon>
            {/* <ActionIcon color="red" variant="transparent" aria-label="Favorites">
              <IconHeart stroke={2} />
            </ActionIcon> */}
          </Group>
        </Group>
        <Stack gap={5}>
          <Title ta="left">{popUp.title}</Title>
          {/* <Group justify='space-between' align='flex-end' gap={4}>
            <Text ta="left" size='md'>{popUp?.companyData?.name}</Text>
            <ActionIcon variant="transparent" aria-label="Company info">
              <IconInfoCircle stroke={2} />
            </ActionIcon>
          </Group> */}
          <Group gap={4}>
            {/* <IconMapPin stroke={1} />
            <ActionIcon variant="transparent" aria-label="Company info">
              <IconMapPin stroke={1} />
            </ActionIcon>
            <Text ta="left" size='md'>{popUp?.companyData?.address}</Text> */}
            <Button
              variant="subtle"
              size="xs"
              p={0}
              leftSection={<IconMapPin stroke={1} />}
            >
              {popUp?.companyData?.address}
            </Button>
          </Group>
          <Group gap={4}>
            <IconCalendar stroke={1} />
            <Text size="sm">
              {formatDateRange(popUp.start_date_time, popUp.end_date_time)}
            </Text>
          </Group>
          {popUp?.promoCode && (
            <Text ta="left" size="sm">
              Promo code: {popUp.promoCode}
            </Text>
          )}
        </Stack>
        <Group justify="space-between" mt="xs" mb="xs">
          <Text size="sm" c="dimmed">
            <TypographyStylesProvider component={"div"}>
              <p
                className="offerUl"
                dangerouslySetInnerHTML={{
                  __html:
                    popUp?.content ??
                    `<p>This is an example of an HTML document with bullet points and bold text.</p>
              <ul>
                  <li><b>Lorem ipsum dolor sit amet</b>, consectetur adipiscing elit. Integer nec odio. Praesent libero. Sed cursus ante dapibus diam.</li>
                  <li><b>Sed nisi</b>. Nulla quis sem at nibh elementum imperdiet. Duis sagittis ipsum. Praesent mauris.</li>
                  <li><b>Fusce nec tellus</b> sed augue semper porta. Mauris massa. Vestibulum lacinia arcu eget nulla.</li>
                  <li><b>Class aptent taciti sociosqu</b> ad litora torquent per conubia nostra, per inceptos himenaeos.</li>
                  <li><b>Curabitur sodales ligula</b> in libero. Sed dignissim lacinia nunc. Curabitur tortor. Pellentesque nibh.</li>
              </ul>`,
                }}
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
      <CompanyModal opened={opened} onClose={close}></CompanyModal>
    </Stack>
  );
}
