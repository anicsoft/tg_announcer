
export interface Offer {
    announcement_id: number;
    company_id: number;
    title: string;
    categories: string[];
    start_date_time: string; // Assuming ISO 8601 time string
    end_date_time: string; // Assuming ISO 8601 time string
    promo_code: string;
    created_at?: string; // Assuming ISO 8601 date-time string
    image?: File | undefined | null;
}

export interface OfferRequest {
    company_id: number;
    title: string;
    categories: string[];
    start_date_time: string; // Assuming ISO 8601 time string
    end_date_time: string; // Assuming ISO 8601 time string
    promo_code: string;
    image?: File | undefined | null;
}

// {
//     "categories": [
//       "Special Offer"
//     ],
//     "company_id": 1,
//     "end_date": "string",
//     "end_time": "string",
//     "promo_code": "null",
//     "start_date": "string",
//     "start_time": "string",
//     "title": "We have free food!"
//   }

export interface CardProps extends Offer {
    logo?: string; // URL for the logo/image
    title: string;
    content?: string;
    address?: string;
    businessName: string;
    workingHours?: string;
    distance?: number | string;
    promoCode?: string;
    companyData?: Company;
    picture_url?: string;
    onClick?: () => void;
}

export interface Company {
    company_id: number;
    name: string;
    description: string;
    address: string;
    latitude: number;
    longitude: number;
    category: string[] | null;
    logo_url: string;
}

export const mock_cards: CardProps[] = [
    {
        logo: "dummy_logo.webp",
        title: "Store A",
        businessName: "Business Name",
        address: "Main Street 10, Tallinn, Estonia",
        workingHours: "09:00 - 21:00",
        distance: "400m",
        categories: ["Groceries", "Bakery", "Fresh Produce", "Beverages"],
        onClick: () => console.log("Store A clicked")
    },
    {
        logo: "dummy_logo.webp",
        title: "Cafe B",
        businessName: "Business Name",
        address: "Old Town Square, Tallinn, Estonia",
        workingHours: "08:00 - 23:00",
        distance: "600m",
        categories: ["Coffee", "Pastry", "Desserts"],
        onClick: () => console.log("Cafe B clicked")
    },
    {
        logo: "dummy_logo.webp",
        title: "Library C",
        businessName: "Business Name",
        address: "Library Street 5, Tallinn, Estonia",
        workingHours: "10:00 - 20:00",
        distance: "700m",
        categories: ["Books", "Magazines", "E-Books"],
        onClick: () => console.log("Library C clicked")
    },
    {
        logo: "dummy_logo.webp",
        title: "Fitness Center D",
        businessName: "Business Name",
        address: "Fitness Avenue, Tallinn, Estonia",
        workingHours: "06:00 - 22:00",
        distance: "1.2km",
        categories: ["Gym", "Yoga", "Pilates"],
        onClick: () => console.log("Fitness Center D clicked")
    },
    {
        logo: "dummy_logo.webp",
        title: "Restaurant E",
        businessName: "Business Name",
        address: "Gourmet Street, Tallinn, Estonia",
        workingHours: "11:30 - 23:30",
        distance: "1km",
        categories: ["Fine Dining", "Seafood", "Wine"],
        onClick: () => console.log("Restaurant E clicked")
    },
    {
        logo: "dummy_logo.webp",
        title: "Pharmacy F",
        businessName: "Business Name",
        address: "Health Street 15, Tallinn, Estonia",
        workingHours: "08:00 - 20:00",
        distance: "800m",
        categories: ["Medication", "Wellness", "Beauty"],
        onClick: () => console.log("Pharmacy F clicked")
    },
    {
        logo: "dummy_logo.webp",
        title: "Salon G",
        businessName: "Business Name",
        address: "Fashion Lane 3, Tallinn, Estonia",
        workingHours: "10:00 - 19:00",
        distance: "550m",
        categories: ["Hair", "Nail", "Beauty"],
        onClick: () => console.log("Salon G clicked")
    },
    {
        logo: "dummy_logo.webp",
        title: "Electronics H",
        businessName: "Business Name",
        address: "Tech Park 7, Tallinn, Estonia",
        workingHours: "10:00 - 20:00",
        distance: "900m",
        categories: ["Gadgets", "Appliances", "Computers"],
        onClick: () => console.log("Electronics H clicked")
    },
    {
        logo: "dummy_logo.webp",
        title: "Clothing Store I",
        businessName: "Business Name",
        address: "Fashion Mall, Tallinn, Estonia",
        workingHours: "10:00 - 21:00",
        distance: "1.1km",
        categories: ["Men's Fashion", "Women's Fashion", "Children's Fashion"],
        onClick: () => console.log("Clothing Store I clicked")
    },
    {
        logo: "dummy_logo.webp",
        title: "Spa J",
        businessName: "Business Name",
        address: "Resort Avenue, Tallinn, Estonia",
        workingHours: "09:00 - 21:00",
        distance: "2km",
        categories: ["Wellness", "Relaxation", "Therapy"],
        onClick: () => console.log("Spa J clicked")
    }
]