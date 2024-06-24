import { Modal } from '@mantine/core'
import OfferCard from '../components/OfferCard'
import React from 'react'
import { OfferRequest } from '../utils/data'

export default function OfferModal({ offer, opened, onClose }: { offer: OfferRequest, opened: boolean, onClose: () => void }) {
  return (
    <Modal opened={opened} onClose={onClose} padding={0} styles={{
      header: { position: "absolute", width: "100%", backgroundColor: "transparent" },
      overlay: {
        backdropFilter: "blur(4px)"
      }
      // content: { paddingInline: "20px" }
    }}>
      <OfferCard popUp={offer}></OfferCard>
    </Modal>
  )
}

// overlayProps={{ backgroundOpacity: 0.4, blur: 1, color: "#E7EAF7" }}