import { Modal } from '@mantine/core'
import CompanyCard from '../components/CompanyCard'
import React from 'react'

export default function CompanyModal({ opened, onClose }: { opened: boolean, onClose: () => void }) {
  return (
    <Modal opened={opened} onClose={onClose} padding={0} styles={{
      header: { position: "absolute", width: "100%", backgroundColor: "transparent" },
      overlay: {
        backdropFilter: "blur(4px)"
      }
      // content: { paddingInline: "20px" }
    }}>
      <CompanyCard></CompanyCard>
    </Modal>
  )
}
