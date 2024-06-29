import { AppContext } from "../../context/AppContext"
import { useContext } from "react"
import { Avatar, Text, Button, Paper } from '@mantine/core';

const UserProfile = () => {
    const { userData } = useContext(AppContext)
    
  return (
    <>
       <Paper radius="md" withBorder p="lg" bg="var(--mantine-color-body)">
      <Avatar
        src="https://raw.githubusercontent.com/mantinedev/mantine/master/.demo/avatars/avatar-8.png"
        size={120}
        radius={120}
        mx="auto"
      />
          <div>Hello {userData?.firstName}</div>

      <Text ta="center" fz="lg" fw={500} mt="md">
        Jane Fingerlicker
      </Text>
      <Text ta="center" c="dimmed" fz="sm">
        jfingerlicker@me.io • Art director
      </Text>

      <Button variant="default" fullWidth mt="md">
        Send message
      </Button>
    </Paper>
    </>
  )
}

export default UserProfile