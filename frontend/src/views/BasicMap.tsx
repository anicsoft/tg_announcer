import { ActionIcon, Flex, LoadingOverlay, Modal, Text, Title, UnstyledButton } from '@mantine/core'
import { useDisclosure } from '@mantine/hooks';
import { MapContainer, Marker, Popup, TileLayer } from 'react-leaflet'
import MarkerClusterGroup from "react-leaflet-cluster";
import OfferCard from '../components/OfferCard';
import MapFilterButton from '../ui/MapFilterButton';
import L, { Icon } from 'leaflet';
import { IconArrowRight } from '@tabler/icons-react';
import { useGeolocation } from './../hooks/useGeolocation';
import { CardProps, Company } from 'utils/data';


export default function BasicMap({ offers }: { offers: CardProps[] }) {
  // console.log("DATA ------------>");
  // console.log(data);

  const OpenStreetMap_HOT = L.tileLayer('https://{s}.tile.openstreetmap.fr/hot/{z}/{x}/{y}.png', {
    maxZoom: 19,
    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors, Tiles style by <a href="https://www.hotosm.org/" target="_blank">Humanitarian OpenStreetMap Team</a> hosted by <a href="https://openstreetmap.fr/" target="_blank">OpenStreetMap France</a>'
  });
  const { latitude, longitude, error } = useGeolocation();
  console.log(latitude, longitude);

  const [opened, { open, close }] = useDisclosure(false);

  const markers = [
    {
      geocode: [48.86, 2.3522],
      popUp: {
        businessName: "Nimeta baar",
        title: "Nihuemoe"
      }
    },
    {
      geocode: [48.92, 2.3522],
      popUp: {
        businessName: "Nimeta baar",
        title: "Huemoe"
      }
    },
    {
      geocode: [48.86, 2.3622],
      popUp: {
        businessName: "Nimeta baar",
        title: "With html",
        content: '<ul><li>General text formatting: <strong>bold</strong>, <em>italic</em>, <u>underline</u>, <s>strike-through</s> </li><li>Headings (h1-h6)</li><li>Sub and super scripts (<sup>&lt;sup /&gt;</sup> and <sub>&lt;sub /&gt;</sub> tags)</li><li>Ordered and bullet lists</li><li>Text align&nbsp;</li><li>And all <a href="https://tiptap.dev/extensions" target="_blank" rel="noopener noreferrer">other extensions</a></li></ul>'
      }
    }
  ]

  console.log(offers);


  // const createClusterCustomIcon = (cluster: any) => {
  //   const icon = divIcon({
  //     html: `<span class="cluster-icon">${cluster.getChildCount()}</span>`,
  //     className: "custom-marker-cluster",
  //     iconSize: point(33, 33, true)
  //   });
  //   return icon
  // };

  const greenIcon = new Icon({
    iconUrl: 'src/assets/anic_xs.svg',
    iconSize: [50, 50], // size of the icon

    // shadowUrl: 'leaf-shadow.png',

    shadowSize: [50, 64], // size of the shadow
    iconAnchor: [22, 94], // point of the icon which will correspond to marker's location
    shadowAnchor: [4, 62],  // the same for the shadow
    popupAnchor: [-3, -76] // point from which the popup should open relative to the iconAnchor
  });

  return (
    <>
      {latitude && longitude ?

        <MapContainer center={[latitude, longitude]} zoom={20}>
          <TileLayer
            attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
            url="https://{s}.tile.openstreetmap.fr/hot/{z}/{x}/{y}.png"
          />
          <MarkerClusterGroup
            chunkedLoading // Performance stuff

          >
            {offers ? offers.map((offer) => (

              <Marker key={offer?.announcement_id} position={[offer?.companyData?.latitude ?? 0, offer?.companyData?.longitude ?? 0]} icon={greenIcon}>
                <Popup>
                  <UnstyledButton onClick={open}>
                    <Title order={6}>{offer.title}</Title>
                    {/* <h2>{marker.popUp.title}</h2> */}
                    <Flex
                      gap="xs"
                      justify="space-between"
                      align="center"
                      direction="row"
                      wrap="nowrap"
                    >
                      <Text size="xs">{offer?.companyData?.name}</Text>
                      <IconArrowRight size={18}></IconArrowRight>
                    </Flex>
                  </UnstyledButton>
                </Popup>
                <Modal opened={opened} onClose={close} title={offer?.title} overlayProps={{ backgroundOpacity: 0.4, blur: 1, color: "#E7EAF7" }}>
                  <OfferCard popUp={offer}></OfferCard>
                </Modal>
              </Marker>
            )
            ) :
              <div> ERRORR!</div>
            }
          </MarkerClusterGroup>
          <Marker key="currentLocation" position={[latitude, longitude]}>
          </Marker>
        </MapContainer>
        :
        <LoadingOverlay visible={true} zIndex={1000} overlayProps={{ radius: "sm", blur: 2 }} />}
    </>
  )
}
