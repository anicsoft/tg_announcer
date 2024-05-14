
export interface CardProps {
  logo?: string; // URL for the logo/image
  title?: string;
  content?: string;
  address?: string;
  businessName?: string;
  workingHours?: string;
  distance?: number |string;
  categories?: string[];
  onClick?: () => void;
}

export const mock_cards:CardProps[] = [
  {
      logo: "logo1.png",
      title: "Store A",
      address: "Main Street 10, Tallinn, Estonia",
      workingHours: "09:00 - 21:00",
      distance: "400m",
      categories: ["Groceries", "Bakery", "Fresh Produce", "Beverages"],
      onClick: () => console.log("Store A clicked")
  },
  {
      logo: "logo2.png",
      title: "Cafe B",
      address: "Old Town Square, Tallinn, Estonia",
      workingHours: "08:00 - 23:00",
      distance: "600m",
      categories: ["Coffee", "Pastry", "Desserts"],
      onClick: () => console.log("Cafe B clicked")
  },
  {
      logo: "logo3.png",
      title: "Library C",
      address: "Library Street 5, Tallinn, Estonia",
      workingHours: "10:00 - 20:00",
      distance: "700m",
      categories: ["Books", "Magazines", "E-Books"],
      onClick: () => console.log("Library C clicked")
  },
  {
      logo: "logo4.png",
      title: "Fitness Center D",
      address: "Fitness Avenue, Tallinn, Estonia",
      workingHours: "06:00 - 22:00",
      distance: "1.2km",
      categories: ["Gym", "Yoga", "Pilates"],
      onClick: () => console.log("Fitness Center D clicked")
  },
  {
      logo: "logo5.png",
      title: "Restaurant E",
      address: "Gourmet Street, Tallinn, Estonia",
      workingHours: "11:30 - 23:30",
      distance: "1km",
      categories: ["Fine Dining", "Seafood", "Wine"],
      onClick: () => console.log("Restaurant E clicked")
  },
  {
      logo: "logo6.png",
      title: "Pharmacy F",
      address: "Health Street 15, Tallinn, Estonia",
      workingHours: "08:00 - 20:00",
      distance: "800m",
      categories: ["Medication", "Wellness", "Beauty"],
      onClick: () => console.log("Pharmacy F clicked")
  },
  {
      logo: "logo7.png",
      title: "Salon G",
      address: "Fashion Lane 3, Tallinn, Estonia",
      workingHours: "10:00 - 19:00",
      distance: "550m",
      categories: ["Hair", "Nail", "Beauty"],
      onClick: () => console.log("Salon G clicked")
  },
  {
      logo: "logo8.png",
      title: "Electronics H",
      address: "Tech Park 7, Tallinn, Estonia",
      workingHours: "10:00 - 20:00",
      distance: "900m",
      categories: ["Gadgets", "Appliances", "Computers"],
      onClick: () => console.log("Electronics H clicked")
  },
  {
      logo: "logo9.png",
      title: "Clothing Store I",
      address: "Fashion Mall, Tallinn, Estonia",
      workingHours: "10:00 - 21:00",
      distance: "1.1km",
      categories: ["Men's Fashion", "Women's Fashion", "Children's Fashion"],
      onClick: () => console.log("Clothing Store I clicked")
  },
  {
      logo: "logo10.png",
      title: "Spa J",
      address: "Resort Avenue, Tallinn, Estonia",
      workingHours: "09:00 - 21:00",
      distance: "2km",
      categories: ["Wellness", "Relaxation", "Therapy"],
      onClick: () => console.log("Spa J clicked")
  }
]