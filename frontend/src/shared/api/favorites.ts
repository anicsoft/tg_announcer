import { BASE_URL } from "./config";

export const getFavorites = async (userId: number) => {
    const response = await fetch(`${BASE_URL}/users/${userId}/favorites`);
    if (!response.ok) {
        throw new Error(`Failed to fetch favorites for user ${userId}`);
    }
    return response.json();
};

export const addFavorite = async (userId: number, companyId: number) => {
    const url = `${BASE_URL}/users/${userId}/favorites`;
    const requestOptions: RequestInit = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ company_id: companyId })
    };

    const response = await fetch(url, requestOptions);
    if (!response.ok) {
        throw new Error(`Failed to toggle favorite for user ${userId}`);
    }
    return response.json();
};

export const removeFavorite = async (userId: number, companyId: number) => {
    const url = `${BASE_URL}/users/${userId}/favorites`;
    const requestOptions: RequestInit = {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ company_id: companyId })
    };

    const response = await fetch(url, requestOptions);
    if (!response.ok) {
        throw new Error(`Failed to remove favorite for user ${userId}`);
    }
    return response.json();
};
