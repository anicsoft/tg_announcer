import React, { useContext } from 'react';
import { Avatar, Text, Paper, List, Anchor, Button } from '@mantine/core';
import { useQuery } from '@tanstack/react-query';
import { AppContext } from '../../context/AppContext';
import { getFavorites, removeFavorite } from '../../shared/api/favorites';

const UserProfile = () => {
    const { userData } = useContext(AppContext);
    const userId = userData?.id;

   const { data, isLoading, isError, refetch } = useQuery({
        queryKey: ["userFavorite", userId],
        queryFn: () => getFavorites(userId),
        enabled: !!userId,
    });

    const favorites = data?.companies || [];



    const onRemoveFavorite = async (companyId: number) => {
        try {
            await removeFavorite(userId, companyId);
            console.log('Favorite removed successfully!');
            refetch(); 
        
        } catch (error) {
            console.error('Error removing favorite:', error);
        }
    };

    return (
        <>
            <Paper radius="md" withBorder p="lg" bg="var(--mantine-color-body)">
                <Avatar
                    src="https://upload.wikimedia.org/wikipedia/commons/thumb/5/59/User-avatar.svg/2048px-User-avatar.svg.png"
                    size={120}
                    radius={120}
                    mx="auto"
                />
                <Text ta="center" fz="lg" fw={500} mt="md">
                    Hello {userData?.firstName}
                </Text>
                <Text ta="center" c="dimmed" fz="sm">
                    user
                </Text>
            </Paper>
            <Text style={{ margin: "20px" }} ta="center" fz="lg" fw={500} mt="lg" mb="lg">
                Your Favorites
            </Text>
            {isLoading && <p>Loading favorites...</p>}
            {isError && <p>Error loading favorites.</p>}
            {favorites.length > 0 ? (
                <Paper radius="md" withBorder p="lg" mt="md">
                    <List>
                        {favorites.map((company) => (
                            <Paper key={company.company_id} radius="md" withBorder p="lg" mt="md">
                                <div style={{ position: 'relative' }}>
                                    <Text fz="lg" fw={500}>{company.name}</Text>
                                    <Avatar src={company.logo_url} size={40} radius="md" style={{ position: 'absolute', top: '0', left: '0' }} />
                                    <Text c="dimmed" fz="sm">{company.description}</Text>
                                    <Text c="dimmed" fz="sm">{company.address}</Text>
                                    <Anchor href={company.website} target="_blank">{company.website}</Anchor>
                                    <Text c="dimmed" fz="sm">{company.tel_number}</Text>
                                    <Text c="dimmed" fz="sm">{company.email}</Text>
                                    <Anchor href={company.facebook} style={{ marginRight: '5px' }} target="_blank">Facebook</Anchor>
                                    <Anchor href={company.instagram} target="_blank">Instagram</Anchor>
                                    <Text c="dimmed" fz="sm">{company.telegram}</Text>

                                    <Button
                                        color="red"
                                        style={{ position: 'absolute', bottom: '0px', right: '0px', width: "60px", height: "20px" }}
                                        onClick={() => onRemoveFavorite(company.company_id)}
                                        variant="outline"
                                    >
                                        X
                                    </Button>
                                </div>
                            </Paper>
                        ))}
                    </List>
                </Paper>
            ) : (
                <Text ta="center" fz="lg" fw={500} mt="md">
                    No favorites available.
                </Text>
            )}
        </>
    );
};

export default UserProfile;
