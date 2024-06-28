import { ActionIcon, Avatar, Flex, LoadingOverlay, Modal, Text, Title, UnstyledButton } from '@mantine/core'
import { useDisclosure } from '@mantine/hooks';
import { MapContainer, Marker, Popup, TileLayer } from 'react-leaflet'
import MarkerClusterGroup from "react-leaflet-cluster";
import OfferCard from '../components/OfferCard';
import L, { Icon, divIcon, point } from 'leaflet';
import { IconArrowRight } from '@tabler/icons-react';
import { useGeolocation } from './../hooks/useGeolocation';
import { CardProps } from 'utils/data';
import OfferModal from '../ui/OfferModal';
import FoxMultipleIcon, { getGroupFox } from '../ui/FoxMultipleIcon';


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

  const markerIcon = new Icon({
    iconUrl: 'src/assets/map-pin.svg',
    iconSize: [36, 36], // size of the icon

    // shadowUrl: 'leaf-shadow.png',

    // shadowSize: [50, 36], // size of the shadow
    iconAnchor: [18, 18], // point of the icon which will correspond to marker's location
    // shadowAnchor: [4, 62],  // the same for the shadow
    popupAnchor: [12, 24] // point from which the popup should open relative to the iconAnchor
  });

  const createCustomClusterIcon = (cluster) => {
    // cluster
    const data = {
      html: getGroupFox(cluster.getChildCount()),
      // html: `<div class="cluster-icon"><h4>${cluster.getChildCount()}x</h4><img/ src="src/assets/pin-fox.svg"></div>`,
      iconSize: point(56, 56, true),
      className: "custom-marker-culster"
    }
    const icon = new divIcon(data)
    return icon
  }

  return (
    <>
      {latitude !== 0 && longitude !== 0 ?

        <MapContainer center={[59.4370, 24.7454]} zoom={20}>
          <TileLayer
            attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
            url="https://{s}.tile.openstreetmap.fr/hot/{z}/{x}/{y}.png"
          />
          <Marker key="currentLocation" position={[59.4370, 24.7454]} icon={markerIcon}>
          </Marker>
          <MarkerClusterGroup
            chunkedLoading // Performance stuff
            iconCreateFunction={createCustomClusterIcon}
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
                <OfferModal opened={opened} onClose={close} offer={offer}></OfferModal>
              </Marker>
            )
            ) :
              undefined
            }
          </MarkerClusterGroup>

        </MapContainer>
        :
        <LoadingOverlay visible={true} zIndex={1000} overlayProps={{ radius: "sm", blur: 2 }} />
      }
    </>
  )
}
